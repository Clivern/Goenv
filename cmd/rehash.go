// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/spacewalkio/goenv/core/module"

	"github.com/spf13/cobra"
)

var rehashCmd = &cobra.Command{
	Use:   "rehash",
	Short: "Refresh binaries under goenv shim directory.",
	Run: func(cmd *cobra.Command, args []string) {

		if HOME == "" {
			fmt.Println("Error! `HOME` environment variable is not set")
			os.Exit(1)
		}

		golang := module.NewGolangEnvironment(HOME)

		err := golang.Configure()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		err = golang.Rehash()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(rehashCmd)
}
