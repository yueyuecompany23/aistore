// Package cmn provides common constants, types, and utilities for AIS clients
// and AIStore.
/*
 * Copyright (c) 2023, NVIDIA CORPORATION. All rights reserved.
 */
package cmn

import (
	"time"

	"github.com/NVIDIA/aistore/cmn/feat"
)

// read-mostly and most often used timeouts: assign at startup to reduce the number of GCO.Get() calls
// updating: a) upon startup, b) periodically, via stats runner, and c) upon receiving new global config

type readMostly struct {
	timeout struct {
		cplane    time.Duration // Config.Timeout.CplaneOperation
		keepalive time.Duration // ditto MaxKeepalive
	}
	features   feat.Flags
	testingEnv bool
}

var Rom readMostly

func (rom *readMostly) init() {
	rom.timeout.cplane = time.Second + time.Millisecond
	rom.timeout.keepalive = 2*time.Second + time.Millisecond
}

func (rom *readMostly) Set(cfg *ClusterConfig) {
	rom.timeout.cplane = cfg.Timeout.CplaneOperation.D()
	rom.timeout.keepalive = cfg.Timeout.MaxKeepalive.D()
	rom.features = cfg.Features
}

func (rom *readMostly) CplaneOperation() time.Duration { return rom.timeout.cplane }
func (rom *readMostly) MaxKeepalive() time.Duration    { return rom.timeout.keepalive }
func (rom *readMostly) Features() feat.Flags           { return rom.features }
func (rom *readMostly) TestingEnv() bool               { return rom.testingEnv }
