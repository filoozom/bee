// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package topology exposes abstractions needed in
// topology-aware components.
package topology

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/ethersphere/bee/pkg/swarm"
)

var (
	ErrNotFound      = errors.New("no peer found")
	ErrWantSelf      = errors.New("node wants self")
	ErrOversaturated = errors.New("oversaturated")
)

type Driver interface {
	PeerAdder
	ClosestPeerer
	EachPeerer
	NeighborhoodDepth() uint8
	SubscribePeersChange() (c <-chan struct{}, unsubscribe func())
	io.Closer
	Snapshot() *KadParams
}

type PeerAdder interface {
	// AddPeers is called when peers are added to the topology backlog
	AddPeers(ctx context.Context, addr ...swarm.Address) error
}

type ClosestPeerer interface {
	// ClosestPeer returns the closest connected peer we have in relation to a
	// given chunk address.
	// This function will ignore peers with addresses provided in skipPeers.
	// Returns topology.ErrWantSelf in case base is the closest to the address.
	ClosestPeer(addr swarm.Address, skipPeers ...swarm.Address) (peerAddr swarm.Address, err error)
}

type EachPeerer interface {
	// EachPeer iterates from closest bin to farthest
	EachPeer(EachPeerFunc) error
	// EachPeerRev iterates from farthest bin to closest
	EachPeerRev(EachPeerFunc) error
}

// EachPeerFunc is a callback that is called with a peer and its PO
type EachPeerFunc func(swarm.Address, uint8) (stop, jumpToNext bool, err error)

type BinInfo struct {
	BinPopulation     uint     `json:"population"`
	BinConnected      uint     `json:"connected"`
	DisconnectedPeers []string `json:"disconnectedPeers"`
	ConnectedPeers    []string `json:"connectedPeers"`
}

type KadBins struct {
	Bin0  BinInfo `json:"bin_0"`
	Bin1  BinInfo `json:"bin_1"`
	Bin2  BinInfo `json:"bin_2"`
	Bin3  BinInfo `json:"bin_3"`
	Bin4  BinInfo `json:"bin_4"`
	Bin5  BinInfo `json:"bin_5"`
	Bin6  BinInfo `json:"bin_6"`
	Bin7  BinInfo `json:"bin_7"`
	Bin8  BinInfo `json:"bin_8"`
	Bin9  BinInfo `json:"bin_9"`
	Bin10 BinInfo `json:"bin_10"`
	Bin11 BinInfo `json:"bin_11"`
	Bin12 BinInfo `json:"bin_12"`
	Bin13 BinInfo `json:"bin_13"`
	Bin14 BinInfo `json:"bin_14"`
	Bin15 BinInfo `json:"bin_15"`
}

type KadParams struct {
	Base           string    `json:"baseAddr"`       // base address string
	Population     int       `json:"population"`     // known
	Connected      int       `json:"connected"`      // connected count
	Timestamp      time.Time `json:"timestamp"`      // now
	NNLowWatermark int       `json:"nnLowWatermark"` // low watermark for depth calculation
	Depth          uint8     `json:"depth"`          // current depth
	Bins           KadBins   `json:"bins"`           // individual bin info
}
