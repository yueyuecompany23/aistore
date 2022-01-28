// Package xreg provides registry and (renew, find) functions for AIS eXtended Actions (xactions).
/*
 * Copyright (c) 2018-2021, NVIDIA CORPORATION. All rights reserved.
 */
package xreg

import (
	"context"

	"github.com/NVIDIA/aistore/cluster"
	"github.com/NVIDIA/aistore/cmn"
	"github.com/NVIDIA/aistore/cmn/debug"
	"github.com/NVIDIA/aistore/stats"
	"github.com/NVIDIA/aistore/xact"
)

type BckSummaryArgs struct {
	Ctx context.Context
	Msg *cmn.BucketSummaryMsg
}

func RegNonBckXact(entry Renewable) {
	debug.Assert(xact.Table[entry.Kind()].Scope != xact.ScopeBck)

	// It is expected that registrations happen at the init time. Therefore, it
	// is safe to assume that no `RenewXYZ` will happen before all xactions
	// are registered. Thus, no locking is needed.
	dreg.nonbckXacts[entry.Kind()] = entry
}

func RenewRebalance(id int64) RenewRes {
	e := dreg.nonbckXacts[cmn.ActRebalance].New(Args{UUID: xact.RebID2S(id)}, nil)
	return dreg.renew(e, nil)
}

func RenewResilver(id string) cluster.Xact {
	e := dreg.nonbckXacts[cmn.ActResilver].New(Args{UUID: id}, nil)
	rns := dreg.renew(e, nil)
	debug.Assert(!rns.IsRunning()) // NOTE: resilver is always preempted
	return rns.Entry.Get()
}

func RenewElection() RenewRes {
	e := dreg.nonbckXacts[cmn.ActElection].New(Args{}, nil)
	return dreg.renew(e, nil)
}

func RenewLRU(id string) RenewRes {
	e := dreg.nonbckXacts[cmn.ActLRU].New(Args{UUID: id}, nil)
	return dreg.renew(e, nil)
}

func RenewStoreCleanup(id string) RenewRes {
	e := dreg.nonbckXacts[cmn.ActStoreCleanup].New(Args{UUID: id}, nil)
	return dreg.renew(e, nil)
}

func RenewDownloader(t cluster.Target, statsT stats.Tracker) RenewRes {
	e := dreg.nonbckXacts[cmn.ActDownload].New(Args{T: t, Custom: statsT}, nil)
	return dreg.renew(e, nil)
}

func RenewETL(t cluster.Target, msg interface{}) RenewRes {
	e := dreg.nonbckXacts[cmn.ActETLInline].New(Args{T: t, Custom: msg}, nil)
	return dreg.renew(e, nil)
}

func RenewBckSummary(ctx context.Context, t cluster.Target, bck *cluster.Bck, msg *cmn.BucketSummaryMsg) RenewRes {
	custom := &BckSummaryArgs{Ctx: ctx, Msg: msg}
	e := dreg.nonbckXacts[cmn.ActSummaryBck].New(Args{T: t, UUID: msg.UUID, Custom: custom}, bck)
	return dreg.renew(e, bck)
}