// Package ais provides core functionality for the AIStore object storage.
/*
 * Copyright (c) 2023, NVIDIA CORPORATION. All rights reserved.
 */
package ais

import (
	"github.com/NVIDIA/aistore/api/apc"
	"github.com/NVIDIA/aistore/cluster/meta"
	"github.com/NVIDIA/aistore/cmn"
	"github.com/NVIDIA/aistore/cmn/cos"
	"github.com/NVIDIA/aistore/cmn/debug"
	"github.com/NVIDIA/aistore/cmn/nlog"
)

type lstcx struct {
	p *proxy
	// arg
	bckFrom *meta.Bck
	bckTo   *meta.Bck
	amsg    *apc.ActMsg // orig
	tcbmsg  *apc.TCBMsg
	config  *cmn.Config
	// work
	tsi    *meta.Snode
	lsmsg  apc.LsoMsg
	altmsg apc.ActMsg
	tcomsg cmn.TCObjsMsg
}

func (c *lstcx) do() (string, error) {
	var (
		p    = c.p
		smap = p.owner.smap.get()
	)
	// 1. lsmsg
	c.lsmsg = apc.LsoMsg{
		UUID:     cos.GenUUID(),
		Prefix:   c.tcbmsg.Prefix,
		Props:    apc.GetPropsName,
		PageSize: 0, // i.e., backend.MaxPageSize()
	}
	c.lsmsg.SetFlag(apc.LsNameOnly)
	tsi, err := smap.HrwTargetTask(c.lsmsg.UUID)
	if err != nil {
		return "", err
	}
	c.tsi = tsi
	c.lsmsg.SID = tsi.ID()

	// 2. ls 1st page
	var lst *cmn.LsoResult
	lst, err = p.lsObjsR(c.bckFrom, &c.lsmsg, smap, tsi /*designated target*/, c.config, true)
	if err != nil {
		return "", err
	}
	if len(lst.Entries) == 0 {
		// TODO: return http status to indicate exactly that (#6393)
		nlog.Infoln(c.amsg.Action, c.bckFrom.Cname(""), " to ", c.bckTo.Cname("")+": lso counts zero - nothing to do")
		return c.lsmsg.UUID, nil
	}

	// 3. tcomsg
	c.tcomsg.ToBck = c.bckTo.Clone()
	c.tcomsg.TCBMsg = *c.tcbmsg
	names := make([]string, 0, len(lst.Entries))
	for _, e := range lst.Entries {
		names = append(names, e.Name)
	}
	c.tcomsg.ListRange.ObjNames = names

	// 4. multi-obj action: transform/copy
	c.altmsg.Value = &c.tcomsg
	c.altmsg.Action = apc.ActCopyObjects
	if c.amsg.Action == apc.ActETLBck {
		c.altmsg.Action = apc.ActETLObjects
	}
	cnt := min(len(names), 10)
	nlog.Infof("(%s => %s): %s => %s %v...", c.amsg.Action, c.altmsg.Action, c.bckFrom, c.bckTo, names[:cnt])

	c.tcomsg.TxnUUID, err = p.tcobjs(c.bckFrom, c.bckTo, &c.altmsg, c.tcbmsg.DryRun)
	if lst.ContinuationToken != "" {
		c.lsmsg.ContinuationToken = lst.ContinuationToken
		go c.pages(smap)
	}
	return c.tcomsg.TxnUUID, err
}

// pages 2..last
func (c *lstcx) pages(smap *smapX) {
	p := c.p
	for {
		// next page
		lst, err := p.lsObjsR(c.bckFrom, &c.lsmsg, smap, c.tsi, c.config, true)
		if err != nil {
			nlog.Errorln(err)
			return
		}
		if len(lst.Entries) == 0 {
			return
		}

		// next tcomsg
		names := make([]string, 0, len(lst.Entries))
		for _, e := range lst.Entries {
			names = append(names, e.Name)
		}
		c.tcomsg.ListRange.ObjNames = names

		// next tco action
		c.altmsg.Value = &c.tcomsg
		xid, err := p.tcobjs(c.bckFrom, c.bckTo, &c.altmsg, c.tcbmsg.DryRun)
		if err != nil {
			nlog.Errorln(err)
			return
		}
		debug.Assertf(c.tcomsg.TxnUUID == xid, "%q vs %q", c.tcomsg.TxnUUID, xid)

		// last page?
		if lst.ContinuationToken == "" {
			return
		}
		c.lsmsg.ContinuationToken = lst.ContinuationToken
	}
}
