// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/norwik/goenv/core/module"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a specific go version.",
	Run: func(cmd *cobra.Command, args []string) {

		golang := module.NewGolangEnvironment(HOME)

		if len(args) == 0 {
			// Show the go releases select box
			versions, err := golang.GetInstalledVersions()

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			list := module.NewCharmSelect(
				"Which go version to uninstall?",
				versions,
			)

			if err := list.Start(); err != nil {
				fmt.Printf("Error showing releases list: %s\n", err.Error())
				os.Exit(1)
			}
		} else {
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

		if !isInstalled {
			fmt.Printf("Error! Version %s is not installed\n", module.SelectedValue)
			os.Exit(1)
		}

		// Show a loading spinner
		spinner := module.NewCharmSpinner(fmt.Sprintf(
			"Removing Go v%s Environment",
			module.SelectedValue,
		))

		go func() {
			// Download and install go selected version
			err = golang.Uninstall(module.SelectedValue)

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
	rootCmd.AddCommand(uninstallCmd)
}
