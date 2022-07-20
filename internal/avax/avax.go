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

package avax

import (
	"fmt"

	"github.com/chain4travel/caminogo/codec"
	"github.com/chain4travel/caminogo/vms/components/avax"
)

func ParseUTXO(ub []byte, cd codec.Manager) (*avax.UTXO, error) {
	utxo := new(avax.UTXO)
	if _, err := cd.Unmarshal(ub, utxo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal utxo bytes: %w", err)
	}
	return utxo, nil
}
