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

package platformvm

import (
	"context"
	"errors"
	"testing"
)

func TestChecker(t *testing.T) {
	t.Parallel()

	ck := NewChecker(nil, nil)
	_, err := ck.PollBlockchain(context.Background())
	if !errors.Is(err, ErrEmptyID) {
		t.Fatalf("unexpected error %v, expected %v", err, ErrEmptyID)
	}
}
