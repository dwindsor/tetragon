// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

// openfile-test calls strconv.ParseInt so uprobes can capture a known go_int (bitSize arg).
package main

import "strconv"

func main() {
	_, _ = strconv.ParseInt("42", 10, 64)
}
