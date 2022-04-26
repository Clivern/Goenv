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

var satisfyCmd = &cobra.Command{
	Use:   "satisfy",
	Short: "Satisfy the current directry go version.",
	Run: func(cmd *cobra.Command, args []string) {

		golang := module.NewGolangEnvironment(HOME)

		cdir, err := os.Getwd()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		version, err := golang.GetLocalVersion(cdir)

		if err == nil {
			// Validate version
			if !golang.ValidateVersion(version) {
				fmt.Printf("Error! Invalid version detected %s\n", version)
				os.Exit(1)
			}

			isInstalled, err := golang.ValidateInstalledVersion(version)

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			if isInstalled {
				fmt.Printf("Go Version %s is already installed\n", version)
				return
			}

			// Show a loading spinner
			spinner := module.NewCharmSpinner(fmt.Sprintf(
				"Getting Go v%s Environment Ready",
				version,
			))

			go func() {
				// Download and install the go version
				err = golang.Install(version)

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

			return
		}

		version, err = golang.GetGlobalVersion()

		if err == nil {
			// Validate version
			if !golang.ValidateVersion(version) {
				fmt.Printf("Error! Invalid version detected %s\n", version)
				os.Exit(1)
			}

			isInstalled, err := golang.ValidateInstalledVersion(version)

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			if isInstalled {
				fmt.Printf("Go Version %s is already installed\n", version)
				return
			}

			// Show a loading spinner
			spinner := module.NewCharmSpinner(fmt.Sprintf(
				"Getting Go v%s Environment Ready",
				version,
			))

			go func() {
				// Download and install the go version
				err = golang.Install(version)

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

			return
		}

		fmt.Println(err.Error())
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(satisfyCmd)
}
