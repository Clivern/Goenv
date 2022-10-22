// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"testing"

	"github.com/clivern/goenv/pkg"

	"github.com/franela/goblin"
)

// TestUnitGolang
func TestUnitGolang(t *testing.T) {
	g := goblin.Goblin(t)

	baseDir := fmt.Sprintf("%s/cache", pkg.GetBaseDir("cache"))
	goenv := NewGolangEnvironment(baseDir)

	g.Describe("#Golang", func() {
		g.It("It should satisfy test cases", func() {
			// Test GetVersions
			g.Assert(len(goenv.GetVersions())).Equal(273)

			// Test SetVersion
			versionFile := fmt.Sprintf("%s/.goenv/%s", baseDir, ".go-version")
			g.Assert(goenv.SetVersion(versionFile, "1.18")).Equal(nil)

			// Test Configure
			g.Assert(goenv.Configure()).Equal(nil)

			// Test GetLocalVersion
			out, err := goenv.GetLocalVersion(fmt.Sprintf("%s/.goenv", baseDir))
			g.Assert(out).Equal("1.18")
			g.Assert(err).Equal(nil)

			// Test GetGlobalVersion
			out, err = goenv.GetGlobalVersion()
			g.Assert(out).Equal("1.18")
			g.Assert(err).Equal(nil)

			// Cleanup
			goenv.FileSystem.ClearDir(fmt.Sprintf("%s/.goenv", baseDir))
			goenv.FileSystem.DeleteDir(fmt.Sprintf("%s/.goenv", baseDir))
		})
	})
}
