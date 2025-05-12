// Copyright 2025 Dandelion Good. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/clivern/goenv/core/module"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func inStringSlice(slice []string, test string) bool {
	for _, sliceVal := range slice {
		if sliceVal == test {
			return true
		}
	}
	return false
}

var showPreRelease bool
var noPatch bool
var displayColor bool

var noColor bool

var maxShow int

var installedColor = color.New(color.FgHiGreen)
var preReleaseColor = color.New(color.FgRed)
var installedPreReleaseColor = color.New(color.FgYellow)


type printerType func(format string, a ...interface{}) 

func defaultPrint(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

var lsCmd = &cobra.Command{
	Use:   "ls-remote",
	Short: "List installable go versions.",
	Run: func(cmd *cobra.Command, args []string) {
		if noColor == true {
			// Handle the --no_color flag
			fmt.Println("checked val of no color is true")
			displayColor = false
		}
		var printer printerType
		golang := module.NewGolangEnvironment(HOME)

		installed, err := golang.GetInstalledVersions()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		for _, remoteRelease := range module.GolangReleases {
			//remoteRelease = "v" + remoteRelease
			printer = defaultPrint
			isPreRelease := (
					strings.Contains(remoteRelease, "rc") ||
					strings.Contains(remoteRelease, "beta") ||
					strings.Contains(remoteRelease, "alpha"))

			if ! showPreRelease && isPreRelease {
				continue
			}
			if isPreRelease && displayColor {
				printer = preReleaseColor.PrintfFunc()
			}

			if noPatch && strings.Count(remoteRelease, ".") > 1 {
				continue
			}

			suffix := ""
			if inStringSlice(installed, remoteRelease) {
				if displayColor && isPreRelease {
					printer = installedPreReleaseColor.PrintfFunc()
				} else if displayColor {
					printer = installedColor.PrintfFunc()
				}
				suffix = " *"
			}
			printer("%s%s\n", remoteRelease, suffix)

		}

	},

}

func init() {



	colorHelp := fmt.Sprintf("display with colors: %s %s %s", installedColor.Sprint("installed version"), preReleaseColor.Sprint("Pre-release version"), installedPreReleaseColor.Sprint("Installed pre-release version"))

	flags := lsCmd.Flags()
	flags.BoolVarP(&showPreRelease, "preRelease", "p", false, "show pre-release releases")
	flags.Lookup("preRelease").NoOptDefVal = "true"
	flags.BoolVarP(&noPatch, "skip_patch", "M", false, "don't show patch releases (major/minor only)")
	flags.Lookup("skip_patch").NoOptDefVal = "true"
	flags.BoolVarP(&displayColor, "color", "c", true, colorHelp)
	flags.BoolVarP(&noColor, "no_color", "", false, "don't display color")
	lsCmd.MarkFlagsMutuallyExclusive("color", "no_color")
	rootCmd.AddCommand(lsCmd)
	

}
