// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"archive/tar"
	"compress/gzip"
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
func (i *Installer) DownloadFromURL(dir string, url string) (string, error) {

	tokens := strings.Split(url, "/")

	fileName := tokens[len(tokens)-1]

	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", fmt.Errorf(
			"Unable to download from %s",
			url,
		)
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

// Untar uncompress a .tar.gz file
func (i *Installer) Untar(extractPath, sourcefilePath string) error {
	file, err := os.Open(sourcefilePath)

	if err != nil {
		return err
	}

	defer file.Close()

	var fileReader io.ReadCloser = file

	if strings.HasSuffix(sourcefilePath, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {
			return err
		}

		defer fileReader.Close()
	}

	tarBallReader := tar.NewReader(fileReader)

	for {
		header, err := tarBallReader.Next()

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		filename := filepath.Join(extractPath, filepath.FromSlash(header.Name))

		switch header.Typeflag {

		case tar.TypeDir:
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				return err
			}

		case tar.TypeReg:
			writer, err := os.Create(filename)

			if err != nil {
				return err
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(filename, os.FileMode(header.Mode))

			if err != nil {
				return err
			}

			writer.Close()
		default:
			return fmt.Errorf("Unable to untar type: %c in file %s", header.Typeflag, filename)
		}
	}

	return nil
}
