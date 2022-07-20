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

// Package implements "status" sub-commands.
package cmd

import (
	"context"

	internal_platformvm "github.com/chain4travel/camino-subnet-cli/internal/platformvm"
	"github.com/chain4travel/camino-subnet-cli/pkg/color"
	"github.com/chain4travel/caminogo/ids"
	pstatus "github.com/chain4travel/caminogo/vms/platformvm/status"
	"github.com/spf13/cobra"
)

func newStatusBlockchainCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blockchain [BLOCKCHAIN ID]",
		Short: "blockchain commands",
		Long: `
Checks the status of the blockchain.

$ subnet-cli status blockchain \
--blockchain-id=[BLOCKCHAIN ID] \
--private-uri=http://localhost:49738 \
--check-bootstrapped

`,
		RunE: createStatusFunc,
	}

	cmd.PersistentFlags().StringVar(&blockchainID, "blockchain-id", "", "blockchain to check the status of")
	cmd.PersistentFlags().BoolVar(&checkBootstrapped, "check-bootstrapped", false, "'true' to wait until the blockchain is bootstrapped")
	return cmd
}

func createStatusFunc(cmd *cobra.Command, args []string) error {
	cli, _, err := InitClient(privateURI, false)
	if err != nil {
		return err
	}

	blkChainID, err := ids.FromString(blockchainID)
	if err != nil {
		return err
	}

	opts := []internal_platformvm.OpOption{
		internal_platformvm.WithBlockchainID(blkChainID),
		internal_platformvm.WithBlockchainStatus(pstatus.Validating),
	}
	if checkBootstrapped {
		opts = append(opts, internal_platformvm.WithCheckBlockchainBootstrapped(cli.Info().Client()))
	}

	color.Outf("\n{{blue}}Checking blockchain...{{/}}\n")
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.P().Checker().PollBlockchain(ctx, opts...)
	cancel()
	return err
}
