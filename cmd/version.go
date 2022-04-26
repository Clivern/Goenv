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

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current go version.",
	Run: func(cmd *cobra.Command, args []string) {

		golang := module.NewGolangEnvironment(HOME)

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
	rootCmd.AddCommand(versionCmd)
}
