// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

// Golang type
type Golang struct {
}

// GetVersions returns a list of all available versions
func (g *Golang) GetVersions() []string {
	return golangReleases
}
