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
	"fmt"

	"github.com/chain4travel/camino-subnet-cli/pkg/color"
	"github.com/chain4travel/caminogo/ids"
	"github.com/chain4travel/caminogo/utils/hashing"
	"github.com/spf13/cobra"
)

const (
	IDLen = 32
)

var h bool

func newCreateVMIDCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "VMID [options] <identifier>",
		Short: "Creates a new encoded VMID from a string",
		RunE:  createVMIDFunc,
	}

	cmd.PersistentFlags().BoolVar(&h, "hash", false, "whether or not to hash the identifier argument")

	return cmd
}

func createVMIDFunc(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument but got %d", len(args))
	}

	identifier := []byte(args[0]) //nolint:ifshort
	var b []byte
	if h {
		b = hashing.ComputeHash256(identifier)
	} else {
		if len(identifier) > IDLen {
			return fmt.Errorf("non-hashed name must be <= 32 bytes, found %d", len(identifier))
		}
		b = make([]byte, IDLen)
		copy(b, identifier)
	}

	id, err := ids.ToID(b)
	if err != nil {
		return err
	}

	color.Outf("{{green}}created a new VMID %s from %s{{/}}\n", id.String(), args[0])
	return nil
}
