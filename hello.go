// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	"github.com/chai2010/ident"
)

func main() {
	fmt.Println(ident.New("main", "main"))
	fmt.Println(ident.New("path/to/pkg", "main", "count"))
	fmt.Println(ident.New("github.com/chai2010/pbgo", "HttpRule"))
	fmt.Println(ident.New("主包", "主函数"))

	// Output:
	// $main_$main
	// $path_to_pkg_$main_$count
	// $github_com_chai2010_pbgo_$HttpRule
	// $主包_$主函数
}
