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

package poll

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestPoll(t *testing.T) {
	t.Parallel()

	pl := New(time.Minute)

	rootCtx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := pl.Poll(rootCtx, nil); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected Poll error %v", err)
	}
}
