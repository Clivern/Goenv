// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetBaseDir returns the project base dir
// Base dir identified if dirName found
// This function for testing purposes only
func GetBaseDir(dirName string) string {
	baseDir, _ := os.Getwd()
	cacheDir := fmt.Sprintf("%s/%s", baseDir, dirName)

	for {
		if fi, err := os.Stat(cacheDir); err == nil {
			if fi.Mode().IsDir() {
				return baseDir
			}
		}
		baseDir = filepath.Dir(baseDir)
		cacheDir = fmt.Sprintf("%s/%s", baseDir, dirName)
	}
}
