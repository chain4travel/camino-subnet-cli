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
	"time"

	"github.com/spf13/cobra"

	"github.com/chain4travel/camino-subnet-cli/pkg/logutil"
)

var rootCmd = &cobra.Command{
	Use:        "subnet-cli",
	Short:      "subnet-cli CLI",
	Version:    Version,
	SuggestFor: []string{"subnet-cli", "subnetcli", "subnetctl"},
}

var (
	enablePrompt bool
	logLevel     string

	privKeyPath string

	privateURI string
	publicURI  string

	pollInterval   time.Duration
	requestTimeout time.Duration

	subnetIDs   string
	nodeIDs     []string
	stakeAmount uint64

	validateEnds             string
	validateWeight           uint64
	validateRewardFeePercent uint32

	rewardAddrs string
	changeAddrs string

	chainName     string
	vmIDs         string
	vmGenesisPath string

	blockchainID      string
	checkBootstrapped bool
)

func init() {
	cobra.EnablePrefixMatching = true

	rootCmd.AddCommand(
		CreateCommand(),
		AddCommand(),
		StatusCommand(),
		WizardCommand(),
	)

	rootCmd.PersistentFlags().BoolVar(&enablePrompt, "enable-prompt", true, "'true' to enable prompt mode")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", logutil.DefaultLogLevel, "log level")
	rootCmd.PersistentFlags().DurationVar(&pollInterval, "poll-interval", time.Second, "interval to poll tx/blockchain status")
	rootCmd.PersistentFlags().DurationVar(&requestTimeout, "request-timeout", 2*time.Minute, "request timeout")
}

func Execute() error {
	if err := CreateLogger(); err != nil {
		return err
	}
	return rootCmd.Execute()
}
