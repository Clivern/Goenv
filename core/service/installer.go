// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Installer struct
type Installer struct {
}

// NewInstaller creates a new instance
func NewInstaller() *Installer {
	return &Installer{}
}

// DownloadFromURL downloads from a URL
func (d *Installer) DownloadFromURL(dir string, url string) (string, error) {

	tokens := strings.Split(url, "/")

	fileName := tokens[len(tokens)-1]

	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", fmt.Errorf("Unable to download from %s", url)
	}

	tarFile := filepath.Join(dir, fileName)

	output, err := os.Create(tarFile)

	if err != nil {
		return "", err
	}
	defer output.Close()

	_, err = io.Copy(output, response.Body)

	if err != nil {
		return "", err
	}

	return tarFile, nil
}
