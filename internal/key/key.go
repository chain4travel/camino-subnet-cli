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

// Package key implements key manager and helper functions.
package key

import (
	"bytes"
	"errors"
	"sort"

	"github.com/chain4travel/caminogo/ids"
	"github.com/chain4travel/caminogo/vms/components/avax"
	"github.com/chain4travel/caminogo/vms/platformvm"
	"github.com/chain4travel/caminogo/vms/secp256k1fx"
)

var (
	ErrInvalidType = errors.New("invalid type")
	ErrCantSpend   = errors.New("can't spend")
)

// Key defines methods for key manager interface.
type Key interface {
	// P returns all formatted P-Chain addresses.
	P() []string
	// Addresses returns the all raw ids.ShortID address.
	Addresses() []ids.ShortID
	// Match attempts to match a list of addresses up to the provided threshold.
	Match(owners *secp256k1fx.OutputOwners, time uint64) ([]uint32, []ids.ShortID, bool)
	// Spend attempts to spend all specified UTXOs (outputs)
	// and returns the new UTXO inputs.
	//
	// If target amount is specified, it only uses the
	// outputs until the total spending is below the target
	// amount.
	Spends(outputs []*avax.UTXO, opts ...OpOption) (
		totalBalanceToSpend uint64,
		inputs []*avax.TransferableInput,
		signers [][]ids.ShortID,
	)
	// Sign generates [numSigs] signatures and attaches them to [pTx].
	Sign(pTx *platformvm.Tx, signers [][]ids.ShortID) error
}

type Op struct {
	time         uint64
	targetAmount uint64
	feeDeduct    uint64
}

type OpOption func(*Op)

func (op *Op) applyOpts(opts []OpOption) {
	for _, opt := range opts {
		opt(op)
	}
}

func WithTime(t uint64) OpOption {
	return func(op *Op) {
		op.time = t
	}
}

func WithTargetAmount(ta uint64) OpOption {
	return func(op *Op) {
		op.targetAmount = ta
	}
}

// To deduct transfer fee from total spend (output).
// e.g., "units.MilliAvax" for X/P-Chain transfer.
func WithFeeDeduct(fee uint64) OpOption {
	return func(op *Op) {
		op.feeDeduct = fee
	}
}

type innerSortTransferableInputsWithSigners struct {
	ins     []*avax.TransferableInput
	signers [][]ids.ShortID
}

func (ins *innerSortTransferableInputsWithSigners) Less(i, j int) bool {
	iID, iIndex := ins.ins[i].InputSource()
	jID, jIndex := ins.ins[j].InputSource()

	switch bytes.Compare(iID[:], jID[:]) {
	case -1:
		return true
	case 0:
		return iIndex < jIndex
	default:
		return false
	}
}
func (ins *innerSortTransferableInputsWithSigners) Len() int { return len(ins.ins) }
func (ins *innerSortTransferableInputsWithSigners) Swap(i, j int) {
	ins.ins[j], ins.ins[i] = ins.ins[i], ins.ins[j]
	ins.signers[j], ins.signers[i] = ins.signers[i], ins.signers[j]
}

// SortTransferableInputsWithSigners sorts the inputs and signers based on the
// input's utxo ID.
//
// This is based off of (generics?): https://github.com/ava-labs/avalancheg_o/blob/224c9fd23d41839201dd0275ac864a845de6e93e/vms/components/avax/transferables.go#L202
func SortTransferableInputsWithSigners(ins []*avax.TransferableInput, signers [][]ids.ShortID) {
	sort.Sort(&innerSortTransferableInputsWithSigners{ins: ins, signers: signers})
}
