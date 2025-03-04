// Copyright 2021 Trim21 <trim21.me@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.

package vars

import (
	"os"
	"path/filepath"

	"sci_hub_p2p/pkg/consts"
)

var appBaseDir string

func GetAppBaseDir() string {
	if appBaseDir != "" {
		return appBaseDir
	}

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	appBaseDir = filepath.Join(home, ".sci-hub-p2p")

	return appBaseDir
}

func GetAppTmpDir() string {
	return filepath.Join(GetAppBaseDir(), "tmp")
}

func IndexesBoltPath() string {
	return filepath.Join(GetAppBaseDir(), "indexes.bolt")
}

func TorrentDBPath() string {
	return filepath.Join(GetAppBaseDir(), "torrent.bolt")
}

func IpfsDBPath() string {
	return filepath.Join(GetAppBaseDir(), consts.IPFSBlockDB)
}
