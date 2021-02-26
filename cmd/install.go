// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/clivern/goenv/core/module"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a go version.",
	Run: func(cmd *cobra.Command, args []string) {
		// Show the go releases select box
		list := module.NewCharmSelect(
			"Which go version to install?",
			module.GolangReleases,
		)

		if err := list.Start(); err != nil {
			fmt.Printf("Error showing releases list: %s\n", err.Error())
			os.Exit(1)
		}

		if module.SelectedValue == "" {
			os.Exit(1)
		}

		// Show a loading spinner
		spinner := module.NewCharmSpinner(fmt.Sprintf(
			"Getting Go v%s Environment Ready",
			module.SelectedValue,
		))

		go func() {
			// Download and install go selected version
			time.Sleep(20 * time.Second)
			spinner.Quit()
		}()

		// Start the spinner
		if err := spinner.Start(); err != nil {
			fmt.Printf("Error showing loading bar: %s\n", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
