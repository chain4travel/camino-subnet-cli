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

package color

import (
	"fmt"

	formatter "github.com/onsi/ginkgo/v2/formatter"
)

// Outputs to stdout.
//
// e.g.,
//   Out("{{green}}{{bold}}hi there %q{{/}}", "aa")
//   Out("{{magenta}}{{bold}}hi therea{{/}} {{cyan}}{{underline}}b{{/}}")
//
// ref.
// https://github.com/onsi/ginkgo/blob/v2.0.0/formatter/formatter.go#L52-L73
//
func Outf(format string, args ...interface{}) {
	s := formatter.F(format, args...)
	fmt.Fprint(formatter.ColorableStdOut, s)
}

// Outputs to stderr.
func Errf(format string, args ...interface{}) {
	s := formatter.F(format, args...)
	fmt.Fprint(formatter.ColorableStdErr, s)
}

func Greenf(format string, args ...interface{}) {
	f := fmt.Sprintf("{{green}}%s{{/}}", format)
	Outf(f, args...)
}

func Redf(format string, args ...interface{}) {
	f := fmt.Sprintf("{{red}}%s{{/}}", format)
	Outf(f, args...)
}

func Bluef(format string, args ...interface{}) {
	f := fmt.Sprintf("{{blue}}%s{{/}}", format)
	Outf(f, args...)
}
