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

// Package client implements client.
// TODO: TO BE MIGRATED TO UPSTREAM AVALANCHEGO.
package client

import (
	"context"
	"errors"
	"net/url"
	"time"

	internal_platformvm "github.com/chain4travel/camino-subnet-cli/internal/platformvm"
	"github.com/chain4travel/camino-subnet-cli/internal/poll"
	"github.com/chain4travel/caminogo/ids"
	avago_constants "github.com/chain4travel/caminogo/utils/constants"
	"github.com/chain4travel/caminogo/vms/avm"
	"github.com/chain4travel/caminogo/vms/platformvm"
	"go.uber.org/zap"
)

var (
	ErrEmptyID         = errors.New("empty ID")
	ErrEmptyURI        = errors.New("empty URI")
	ErrInvalidInterval = errors.New("invalid interval")
)

type Config struct {
	URI          string
	u            *url.URL
	PollInterval time.Duration
}

var _ Client = &client{}

type Client interface {
	NetworkID() uint32
	Config() Config
	Info() Info
	KeyStore() KeyStore
	P() P
}

type client struct {
	cfg Config

	// fetched automatic
	networkName string
	networkID   uint32
	assetID     ids.ID
	xChainID    ids.ID
	pChainID    ids.ID

	i *info
	k *keyStore
	p *p
}

func New(cfg Config) (Client, error) {
	if cfg.URI == "" {
		return nil, ErrEmptyURI
	}
	if cfg.PollInterval == time.Duration(0) {
		return nil, ErrInvalidInterval
	}

	u, err := url.Parse(cfg.URI)
	if err != nil {
		return nil, err
	}
	cfg.u = u

	cli := &client{
		cfg:      cfg,
		pChainID: avago_constants.PlatformChainID,
		i:        newInfo(cfg),
		k:        newKeyStore(cfg),
	}

	zap.L().Info("fetching X-Chain id")
	xChainID, err := cli.i.Client().GetBlockchainID(context.TODO(), "X")
	if err != nil {
		return nil, err
	}
	cli.xChainID = xChainID
	zap.L().Info("fetched X-Chain id", zap.String("id", cli.xChainID.String()))

	uriX := u.Scheme + "://" + u.Host
	xChainName := cli.xChainID.String()
	if u.Port() == "" {
		// ref. https://docs.avax.network/build/avalanchego-apis/x-chain
		// e.g., https://api.avax-test.network
		xChainName = "X"
	}
	zap.L().Info("fetching CAM asset id",
		zap.String("uri", uriX),
	)
	xc := avm.NewClient(uriX, xChainName)
	avaxDesc, err := xc.GetAssetDescription(context.TODO(), "CAM")
	if err != nil {
		return nil, err
	}
	cli.assetID = avaxDesc.AssetID
	zap.L().Info("fetched CAM asset id", zap.String("id", cli.assetID.String()))

	zap.L().Info("fetching network information")
	cli.networkName, err = cli.i.Client().GetNetworkName(context.TODO())
	if err != nil {
		return nil, err
	}
	cli.networkID, err = avago_constants.NetworkID(cli.networkName)
	if err != nil {
		return nil, err
	}
	zap.L().Info("fetched network information",
		zap.Uint32("networkId", cli.networkID),
		zap.String("networkName", cli.networkName),
	)

	// "NewClient" already appends "/ext/P"
	// e.g., https://api.avax-test.network
	// ref. https://docs.avax.network/build/avalanchego-apis/p-chain
	uriP := u.Scheme + "://" + u.Host
	pc := platformvm.NewClient(uriP)
	cli.p = &p{
		cfg: cfg,

		networkName: cli.networkName,
		networkID:   cli.networkID,
		assetID:     cli.assetID,
		pChainID:    cli.pChainID,

		cli:  pc,
		info: cli.i.Client(),
		checker: internal_platformvm.NewChecker(
			poll.New(cfg.PollInterval),
			pc,
		),
	}
	return cli, nil
}

func (cc *client) NetworkID() uint32 { return cc.networkID }
func (cc *client) Config() Config    { return cc.cfg }

func (cc *client) Info() Info         { return cc.i }
func (cc *client) KeyStore() KeyStore { return cc.k }

func (cc *client) P() P { return cc.p }
