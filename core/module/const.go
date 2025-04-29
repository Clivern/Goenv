// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"runtime"
	"strings"
)

// golangRegistry the registry to download from
var golangRegistry = "https://dl.google.com/go/go${VERSION}.${OS}-${ARCH}.tar.gz"

// golangRegistryChecksum the registry checksum URL
var golangRegistryChecksum = "https://dl.google.com/go/go${VERSION}.${OS}-${ARCH}.tar.gz.sha256"

// golangArchive the archive name
var golangArchive = "go${VERSION}.${OS}-${ARCH}.tar.gz"

// GolangReleases a list of golang releases
var GolangReleases = []string{
	"1.5beta1",
	"1.5beta2",
	"1.5beta3",
	"1.5rc1",
	"1.5",
	"1.5.1",
	"1.5.2",
	"1.5.3",
	"1.5.4",
	"1.6beta1",
	"1.6beta2",
	"1.6rc1",
	"1.6rc2",
	"1.6",
	"1.6.1",
	"1.6.2",
	"1.6.3",
	"1.6.4",
	"1.7beta1",
	"1.7beta2",
	"1.7rc1",
	"1.7rc2",
	"1.7rc3",
	"1.7rc4",
	"1.7rc5",
	"1.7rc6",
	"1.7",
	"1.7.1",
	"1.7.2",
	"1.7.3",
	"1.7.4",
	"1.7.5",
	"1.7.6",
	"1.8beta1",
	"1.8beta2",
	"1.8rc1",
	"1.8rc2",
	"1.8rc3",
	"1.8",
	"1.8.1",
	"1.8.2",
	"1.8.3",
	"1.8.4",
	"1.8.5",
	"1.8.5rc4",
	"1.8.6",
	"1.8.7",
	"1.9beta1",
	"1.9beta2",
	"1.9rc1",
	"1.9rc2",
	"1.9",
	"1.9.1",
	"1.9.2",
	"1.9.3",
	"1.9.4",
	"1.9.5",
	"1.9.6",
	"1.9.7",
	"1.10beta1",
	"1.10beta2",
	"1.10rc1",
	"1.10rc2",
	"1.10",
	"1.10.1",
	"1.10.2",
	"1.10.3",
	"1.10.4",
	"1.10.5",
	"1.10.6",
	"1.10.7",
	"1.10.8",
	"1.11beta1",
	"1.11beta2",
	"1.11beta3",
	"1.11rc1",
	"1.11rc2",
	"1.11",
	"1.11.1",
	"1.11.2",
	"1.11.3",
	"1.11.4",
	"1.11.5",
	"1.11.6",
	"1.11.7",
	"1.11.8",
	"1.11.9",
	"1.11.10",
	"1.11.11",
	"1.11.12",
	"1.11.13",
	"1.12beta1",
	"1.12beta2",
	"1.12rc1",
	"1.12",
	"1.12.1",
	"1.12.2",
	"1.12.3",
	"1.12.4",
	"1.12.5",
	"1.12.6",
	"1.12.7",
	"1.12.8",
	"1.12.9",
	"1.12.10",
	"1.12.11",
	"1.12.12",
	"1.12.13",
	"1.12.14",
	"1.12.15",
	"1.12.16",
	"1.12.17",
	"1.13beta1",
	"1.13rc1",
	"1.13rc2",
	"1.13",
	"1.13.1",
	"1.13.2",
	"1.13.3",
	"1.13.4",
	"1.13.5",
	"1.13.6",
	"1.13.7",
	"1.13.8",
	"1.13.9",
	"1.13.10",
	"1.13.11",
	"1.13.12",
	"1.13.13",
	"1.13.14",
	"1.13.15",
	"1.14beta1",
	"1.14rc1",
	"1.14",
	"1.14.1",
	"1.14.2",
	"1.14.3",
	"1.14.4",
	"1.14.5",
	"1.14.6",
	"1.14.7",
	"1.14.8",
	"1.14.9",
	"1.14.10",
	"1.14.11",
	"1.14.12",
	"1.14.13",
	"1.14.14",
	"1.14.15",
	"1.15beta1",
	"1.15rc1",
	"1.15rc2",
	"1.15",
	"1.15.1",
	"1.15.2",
	"1.15.3",
	"1.15.4",
	"1.15.5",
	"1.15.6",
	"1.15.7",
	"1.15.8",
	"1.15.9",
	"1.15.10",
	"1.15.11",
	"1.15.12",
	"1.15.13",
	"1.15.14",
	"1.15.15",
	"1.16beta1",
	"1.16rc1",
	"1.16",
	"1.16.1",
	"1.16.2",
	"1.16.3",
	"1.16.4",
	"1.16.5",
	"1.16.6",
	"1.16.7",
	"1.16.8",
	"1.16.9",
	"1.16.10",
	"1.16.11",
	"1.16.12",
	"1.16.13",
	"1.16.14",
	"1.16.15",
	"1.17beta1",
	"1.17rc1",
	"1.17rc2",
	"1.17",
	"1.17.1",
	"1.17.2",
	"1.17.3",
	"1.17.4",
	"1.17.5",
	"1.17.6",
	"1.17.7",
	"1.17.8",
	"1.17.9",
	"1.17.10",
	"1.17.11",
	"1.17.12",
	"1.17.13",
	"1.18beta1",
	"1.18beta2",
	"1.18rc1",
	"1.18",
	"1.18.1",
	"1.18.2",
	"1.18.3",
	"1.18.4",
	"1.18.5",
	"1.18.6",
	"1.18.7",
	"1.18.8",
	"1.18.9",
	"1.18.10",
	"1.19beta1",
	"1.19rc1",
	"1.19rc2",
	"1.19",
	"1.19.1",
	"1.19.2",
	"1.19.3",
	"1.19.4",
	"1.19.5",
	"1.19.6",
	"1.19.7",
	"1.19.8",
	"1.19.9",
	"1.19.10",
	"1.19.11",
	"1.19.12",
	"1.19.13",
	"1.20rc1",
	"1.20rc2",
	"1.20rc3",
	"1.20",
	"1.20.1",
	"1.20.2",
	"1.20.3",
	"1.20.4",
	"1.20.5",
	"1.20.6",
	"1.20.7",
	"1.20.8",
	"1.20.9",
	"1.20.10",
	"1.20.11",
	"1.20.12",
	"1.20.13",
	"1.20.14",
	"1.21rc1",
	"1.21rc2",
	"1.21rc3",
	"1.21rc4",
	"1.21.0",
	"1.21.1",
	"1.21.2",
	"1.21.3",
	"1.21.4",
	"1.21.5",
	"1.21.6",
	"1.21.7",
	"1.21.8",
	"1.21.9",
	"1.21.10",
	"1.21.11",
	"1.21.12",
	"1.21.13",
	"1.22rc1",
	"1.22rc2",
	"1.22.0",
	"1.22.1",
	"1.22.2",
	"1.22.3",
	"1.22.4",
	"1.22.5",
	"1.22.6",
	"1.22.7",
	"1.22.8",
	"1.22.9",
	"1.22.10",
	"1.22.11",
	"1.23rc1",
	"1.23rc2",
	"1.23.0",
	"1.23.1",
	"1.23.2",
	"1.23.3",
	"1.23.4",
	"1.23.5",
	"1.24rc1",
	"1.24rc2",
	"1.24.0",
	"1.24.1",
	"1.24.2",
}

// goShimContent shim for go binary
var goShimContent = `#!/usr/bin/env bash

GO_BINARY_NAME=go

GO_VERSION_PATH=$(goenv exec)

GOENV_ROOT="$HOME/%s" \
	GOPATH=$GO_VERSION_PATH \
	GOBIN=$GO_VERSION_PATH/bin \
	"$GO_VERSION_PATH/bin/$GO_BINARY_NAME" "$@"

if [[ "$@" =~ 'get' ]]
then
	goenv rehash
fi

if [[ "$@" =~ 'install' ]]
then
	goenv rehash
fi

`

// binaryShimContent shim other go binaries
var binaryShimContent = `#!/usr/bin/env bash

GO_BINARY_NAME=%s

GO_VERSION_PATH=$(goenv exec)

GOENV_ROOT="$HOME/%s" \
	GOPATH=$GO_VERSION_PATH \
	GOBIN=$GO_VERSION_PATH/bin \
	"$GO_VERSION_PATH/bin/$GO_BINARY_NAME" "$@"

`

// getDownloadURL returns the download link
func getDownloadURL(version string) string {
	url := strings.Replace(golangRegistry, "${VERSION}", version, -1)
	url = strings.Replace(url, "${OS}", runtime.GOOS, -1)
	url = strings.Replace(url, "${ARCH}", runtime.GOARCH, -1)

	return url
}

// getChecksumURL returns the checksum link
func getChecksumURL(version string) string {
	url := strings.Replace(golangRegistryChecksum, "${VERSION}", version, -1)
	url = strings.Replace(url, "${OS}", runtime.GOOS, -1)
	url = strings.Replace(url, "${ARCH}", runtime.GOARCH, -1)

	return url
}

// getArchiveName returns the archive name
func getArchiveName(version string) string {
	name := strings.Replace(golangArchive, "${VERSION}", version, -1)
	name = strings.Replace(name, "${OS}", runtime.GOOS, -1)
	name = strings.Replace(name, "${ARCH}", runtime.GOARCH, -1)

	return name
}
