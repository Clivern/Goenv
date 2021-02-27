// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Show the current go version.",
	Run: func(cmd *cobra.Command, args []string) {
		// If the command is go install or go get, it will create a new chim with the new binary in .goenv/shims
		// otherwise it return the current version path
		fmt.Printf("%s", "/Users/ahmetwal/.goenv/versions/1.16")
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
