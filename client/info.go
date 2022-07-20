// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
//
// This file is a derived work, based on ava-labs code whose
// original notices appear below.
//
// It is distributed under the same license conditions as the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********************************************************

// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package client

import (
	api_info "github.com/chain4travel/caminogo/api/info"
)

type Info interface {
	Client() api_info.Client
}

type info struct {
	cli api_info.Client
	cfg Config
}

func newInfo(cfg Config) *info {
	// "NewClient" already appends "/ext/info"
	// e.g., https://api.avax-test.network
	// ref. https://docs.avax.network/build/avalanchego-apis/info
	uri := cfg.u.Scheme + "://" + cfg.u.Host
	cli := api_info.NewClient(uri)
	return &info{
		cli: cli,
		cfg: cfg,
	}
}

func (i *info) Client() api_info.Client { return i.cli }
