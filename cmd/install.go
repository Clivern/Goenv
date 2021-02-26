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

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a go version.",
	Run: func(cmd *cobra.Command, args []string) {

		golang := module.NewGolangEnvironment(HOME)

		if len(args) == 0 {
			// Show the go releases select box
			list := module.NewCharmSelect(
				"Which go version to install?",
				module.GolangReleases,
			)

			if err := list.Start(); err != nil {
				fmt.Printf("Error showing releases list: %s\n", err.Error())
				os.Exit(1)
			}
		} else {
			args[0] = strings.TrimPrefix(args[0], "go")

			module.SelectedValue = args[0]
		}

		if !golang.ValidateVersion(module.SelectedValue) {
			fmt.Printf("Error! Invalid version provided %s\n", module.SelectedValue)
			os.Exit(1)
		}

		isInstalled, err := golang.ValidateInstalledVersion(module.SelectedValue)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if isInstalled {
			fmt.Printf("Error! Version %s is already installed\n", module.SelectedValue)
			os.Exit(1)
		}

		// Show a loading spinner
		spinner := module.NewCharmSpinner(fmt.Sprintf(
			"Getting Go v%s Environment Ready",
			module.SelectedValue,
		))

		go func() {
			// Download and install go selected version
			err = golang.Install(module.SelectedValue)

			if err != nil {
				fmt.Println(err.Error())
			}

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
