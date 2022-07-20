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
	api_keystore "github.com/chain4travel/caminogo/api/keystore"
)

type KeyStore interface {
	Client() api_keystore.Client
}

type keyStore struct {
	cli api_keystore.Client
	cfg Config
}

func newKeyStore(cfg Config) *keyStore {
	// "NewClient" already appends "/ext/keystore"
	// e.g., https://api.avax-test.network
	// ref. https://docs.avax.network/build/avalanchego-apis/keystore
	uri := cfg.u.Scheme + "://" + cfg.u.Host
	cli := api_keystore.NewClient(uri)
	return &keyStore{
		cli: cli,
		cfg: cfg,
	}
}

func (k *keyStore) Client() api_keystore.Client { return k.cli }
