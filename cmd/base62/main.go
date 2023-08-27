/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package main

import (
	"fmt"
	"github.com/codeallergy/base62/app"
	"os"
)

var (
	Name    = "base62"
	Version = ""
	Build   = ""
)

func main() {
	if err := app.Run(Name, Version, Build); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
