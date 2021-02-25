// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"testing"

	"github.com/franela/goblin"
)

// TestUnitCorrelation
func TestUnitCorrelation(t *testing.T) {
	g := goblin.Goblin(t)

	corr := NewCorrelation()

	g.Describe("#Correlation", func() {
		g.It("It should satisfy test cases", func() {
			g.Assert(corr.UUIDv4() == "").Equal(false)
			g.Assert(corr.UUIDv4() != "").Equal(true)
		})
	})
}
