// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/clivern/goenv/core/module"

	"github.com/spf13/cobra"
)

var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Set or show the local application-specific go version.",
	Run: func(cmd *cobra.Command, args []string) {

		golang := module.NewGolangEnvironment(HOME)

		if len(args) == 1 {
			args[0] = strings.TrimPrefix(args[0], "go")

			if !golang.ValidateVersion(args[0]) {
				fmt.Printf("Error! Invalid version provided %s\n", args[0])
				os.Exit(1)
			}

			// Set the local version
			err := golang.SetLocalVersion(args[0])

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			fmt.Printf("Local version updated into %s\n", args[0])
			return
		}

		cdir, err := os.Getwd()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		version, err := golang.GetLocalVersion(cdir)

		if err == nil {
			fmt.Println(version)
			return
		}

		fmt.Println("Unable to find local version, fallback into global version")

		version, err = golang.GetGlobalVersion()

		if err == nil {
			fmt.Println(version)
			return
		}

		fmt.Println(err.Error())
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(localCmd)
}
