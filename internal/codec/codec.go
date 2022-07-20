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

package codec

import (
	"github.com/chain4travel/caminogo/codec"
	"github.com/chain4travel/caminogo/codec/linearcodec"
	"github.com/chain4travel/caminogo/utils/wrappers"
	"github.com/chain4travel/caminogo/vms/platformvm"
	"github.com/chain4travel/caminogo/vms/secp256k1fx"
)

var PCodecManager codec.Manager

func init() {
	pc := linearcodec.NewDefault()
	PCodecManager = codec.NewDefaultManager()
	errs := wrappers.Errs{}
	errs.Add(
		pc.RegisterType(&platformvm.ProposalBlock{}),
		pc.RegisterType(&platformvm.AbortBlock{}),
		pc.RegisterType(&platformvm.CommitBlock{}),
		pc.RegisterType(&platformvm.StandardBlock{}),
		pc.RegisterType(&platformvm.AtomicBlock{}),
		pc.RegisterType(&secp256k1fx.TransferInput{}),
		pc.RegisterType(&secp256k1fx.MintOutput{}),
		pc.RegisterType(&secp256k1fx.TransferOutput{}),
		pc.RegisterType(&secp256k1fx.MintOperation{}),
		pc.RegisterType(&secp256k1fx.Credential{}),
		pc.RegisterType(&secp256k1fx.Input{}),
		pc.RegisterType(&secp256k1fx.OutputOwners{}),
		pc.RegisterType(&platformvm.UnsignedAddValidatorTx{}),
		pc.RegisterType(&platformvm.UnsignedAddSubnetValidatorTx{}),
		pc.RegisterType(&platformvm.UnsignedAddDelegatorTx{}),
		pc.RegisterType(&platformvm.UnsignedCreateChainTx{}),
		pc.RegisterType(&platformvm.UnsignedCreateSubnetTx{}),
		pc.RegisterType(&platformvm.UnsignedImportTx{}),
		pc.RegisterType(&platformvm.UnsignedExportTx{}),
		pc.RegisterType(&platformvm.UnsignedAdvanceTimeTx{}),
		pc.RegisterType(&platformvm.UnsignedRewardValidatorTx{}),
		pc.RegisterType(&platformvm.StakeableLockIn{}),
		pc.RegisterType(&platformvm.StakeableLockOut{}),
		PCodecManager.RegisterCodec(0, pc),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}
