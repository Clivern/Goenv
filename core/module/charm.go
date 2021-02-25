// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

// Charm type
type Charm struct {
	Items    []string
	Selected string
	Title    string
	Callback func(string) error
}
