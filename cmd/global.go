// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/clevenio/goenv/core/module"

	"github.com/spf13/cobra"
)

var globalCmd = &cobra.Command{
	Use:   "global",
	Short: "Set or show the global go version.",
	Run: func(cmd *cobra.Command, args []string) {

		golang := module.NewGolangEnvironment(HOME)

		if len(args) == 1 {
			if !golang.ValidateVersion(args[0]) {
				fmt.Printf("Error! Invalid version provided %s\n", args[0])
				os.Exit(1)
			}

			// Set the global version
			err := golang.SetGlobalVersion(args[0])

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			fmt.Printf("Global version updated into %s\n", args[0])
			return
		}

		// Get the global version
		version, err := golang.GetGlobalVersion()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(globalCmd)
}
