// Copyright (c) arkade author(s) 2020. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package config

import (
	"fmt"
	"os"
	"path"
)

func GetUserDir() string {
	home := os.Getenv("HOME")
	root := fmt.Sprintf("%s/.arkade/", home)
	return root
}

func InitUserDir() (string, error) {
	home := os.Getenv("HOME")
	root := fmt.Sprintf("%s/.arkade/", home)

	if len(home) == 0 {
		return home, fmt.Errorf("env-var HOME, not set")
	}

	binPath := path.Join(root, "/bin/")
	err := os.MkdirAll(binPath, 0700)
	if err != nil {
		return binPath, err
	}

	helmPath := path.Join(root, "/.helm/")
	helmErr := os.MkdirAll(helmPath, 0700)
	if helmErr != nil {
		return helmPath, helmErr
	}

	return root, nil
}
