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

package key

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/chain4travel/caminogo/utils/crypto"
	"github.com/chain4travel/caminogo/utils/formatting"
)

const (
	ewoqPChainAddr    = "P-custom18jma8ppw3nhx5r4ap8clazz0dps7rv5u9xde7p"
	fallbackNetworkID = 999999 // unaffiliated networkID should trigger HRP Fallback
)

func TestPrivateKeyToPKfile(t *testing.T) {
	pkWithoutPrefix := "5XnB92EgtSz438AvyTjRpiuekKskrKyLG3e6brughhaqxBf6X"
	pkBytes, err := formatting.Decode(formatting.CB58, pkWithoutPrefix)
	if err != nil {
		t.Fatal(err)
	}
	m := &SoftKey{
		privKeyRaw: pkBytes,
	}
	fmt.Println(t.TempDir())
	keyPath := filepath.Join("/home/kkyriakis/dev/c4t", pkWithoutPrefix+".pk")
	if err := m.Save(keyPath); err != nil {
		t.Fatal(err)
	}
	fmt.Println(t.TempDir())
}
func TestNewKeyEwoq(t *testing.T) {
	t.Parallel()

	m, err := NewSoft(
		fallbackNetworkID,
		WithPrivateKeyEncoded(EwoqPrivateKey),
	)
	if err != nil {
		t.Fatal(err)
	}

	if m.P()[0] != ewoqPChainAddr {
		t.Fatalf("unexpected P-Chain address %q, expected %q", m.P(), ewoqPChainAddr)
	}

	keyPath := filepath.Join(t.TempDir(), "key.pk")
	if err := m.Save(keyPath); err != nil {
		t.Fatal(err)
	}

	m2, err := LoadSoft(fallbackNetworkID, keyPath)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(m.Raw(), m2.Raw()) {
		t.Fatalf("loaded key unexpected %v, expected %v", m2.Raw(), m.Raw())
	}
}

func TestNewKey(t *testing.T) {
	t.Parallel()

	skBytes, err := formatting.Decode(formatting.CB58, rawEwoqPk)
	if err != nil {
		t.Fatal(err)
	}
	factory := &crypto.FactorySECP256K1R{}
	rpk, err := factory.ToPrivateKey(skBytes)
	if err != nil {
		t.Fatal(err)
	}
	ewoqPk, _ := rpk.(*crypto.PrivateKeySECP256K1R)

	rpk2, err := factory.NewPrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	privKey2, _ := rpk2.(*crypto.PrivateKeySECP256K1R)

	tt := []struct {
		name   string
		opts   []SOpOption
		expErr error
	}{
		{
			name:   "test",
			opts:   nil,
			expErr: nil,
		},
		{
			name: "ewop with WithPrivateKey",
			opts: []SOpOption{
				WithPrivateKey(ewoqPk),
			},
			expErr: nil,
		},
		{
			name: "ewop with WithPrivateKeyEncoded",
			opts: []SOpOption{
				WithPrivateKeyEncoded(EwoqPrivateKey),
			},
			expErr: nil,
		},
		{
			name: "ewop with WithPrivateKey/WithPrivateKeyEncoded",
			opts: []SOpOption{
				WithPrivateKey(ewoqPk),
				WithPrivateKeyEncoded(EwoqPrivateKey),
			},
			expErr: nil,
		},
		{
			name: "ewop with invalid WithPrivateKey",
			opts: []SOpOption{
				WithPrivateKey(privKey2),
				WithPrivateKeyEncoded(EwoqPrivateKey),
			},
			expErr: ErrInvalidPrivateKey,
		},
	}
	for i, tv := range tt {
		_, err := NewSoft(fallbackNetworkID, tv.opts...)
		if !errors.Is(err, tv.expErr) {
			t.Fatalf("#%d(%s): unexpected error %v, expected %v", i, tv.name, err, tv.expErr)
		}
	}
}
