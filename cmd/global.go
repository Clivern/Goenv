// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var globalCmd = &cobra.Command{
	Use:   "global",
	Short: "Set or show the global go version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("...")
	},
}

func init() {
	rootCmd.AddCommand(globalCmd)
}
