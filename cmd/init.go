// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/norwik/goenv/core/module"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init the import path for goenv shims.",
	Run: func(cmd *cobra.Command, args []string) {

		if HOME == "" {
			return
		}

		golang := module.NewGolangEnvironment(HOME)

		fmt.Printf("export PATH=\"%s/%s:${PATH}\"\n", golang.RootPath, golang.ShimsDir)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
