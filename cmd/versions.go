// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/clivern/goenv/core/module"

	"github.com/spf13/cobra"
)

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List installed go versions.",
	Run: func(cmd *cobra.Command, args []string) {

		golang := module.NewGolangEnvironment(HOME)

		versions, err := golang.GetInstalledVersions()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		for i := 0; i < len(versions); i++ {
			fmt.Println(versions[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(versionsCmd)
}
