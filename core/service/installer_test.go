// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"testing"

	"github.com/norwik/goenv/pkg"

	"github.com/franela/goblin"
)

// TestUnitInstaller
func TestUnitInstaller(t *testing.T) {
	g := goblin.Goblin(t)

	installer := NewInstaller()
	fs := NewFileSystem()

	g.Describe("#Installer", func() {
		g.It("It should satisfy test cases", func() {
			out, err := installer.DownloadFromURL(
				fmt.Sprintf("%s/cache/", pkg.GetBaseDir("cache")),
				"https://github.com/Clivern/Rhino/releases/download/1.6.1/Rhino_1.6.1_Darwin_x86_64.tar.gz",
			)

			g.Assert(err).Equal(nil)
			g.Assert(out != "").Equal(true)

			// Create a dir for output
			fs.EnsureDir(fmt.Sprintf("%s/cache/rhino", pkg.GetBaseDir("cache")), 0755)

			err = installer.Untar(
				fmt.Sprintf("%s/cache/rhino", pkg.GetBaseDir("cache")),
				fmt.Sprintf("%s/cache/Rhino_1.6.1_Darwin_x86_64.tar.gz", pkg.GetBaseDir("cache")),
			)

			g.Assert(err).Equal(nil)

			g.Assert(fs.FileExists(fmt.Sprintf("%s/cache/rhino/README.md", pkg.GetBaseDir("cache")))).Equal(true)

			// Cleanup
			fs.ClearDir(fmt.Sprintf("%s/cache/rhino", pkg.GetBaseDir("cache")))
			fs.DeleteDir(fmt.Sprintf("%s/cache/rhino", pkg.GetBaseDir("cache")))
			g.Assert(fs.DeleteFile(fmt.Sprintf("%s/cache/Rhino_1.6.1_Darwin_x86_64.tar.gz", pkg.GetBaseDir("cache")))).Equal(nil)
		})
	})
}
