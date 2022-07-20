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

package cmd

import (
	"os"

	"github.com/chain4travel/camino-subnet-cli/internal/key"
	"github.com/chain4travel/camino-subnet-cli/pkg/color"
	"github.com/spf13/cobra"
)

func newCreateKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "key [options]",
		Short: "Generates a private key",
		Long: `
Generates a private key.

$ subnet-cli create key --private-key-path=.insecure.test.key

`,
		RunE: createKeyFunc,
	}
	return cmd
}

func createKeyFunc(cmd *cobra.Command, args []string) error {
	if _, err := os.Stat(privKeyPath); err == nil {
		color.Outf("{{red}}key already found at %q{{/}}\n", privKeyPath)
		return os.ErrExist
	}
	k, err := key.NewSoft(0)
	if err != nil {
		return err
	}
	if err := k.Save(privKeyPath); err != nil {
		return err
	}
	color.Outf("{{green}}created a new key %q{{/}}\n", privKeyPath)
	return nil
}
