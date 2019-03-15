// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	"github.com/chai2010/cident"
)

func main() {
	fmt.Println(cident.New("main", "main"))
	fmt.Println(cident.New("path/to/pkg", "main", "count"))
	fmt.Println(cident.New("github.com/chai2010/pbgo", "HttpRule"))
	fmt.Println(cident.New("主包", "主函数"))

	// Output:
	// $main_$main
	// $path_to_pkg_$main_$count
	// $github_com_chai2010_pbgo_$HttpRule
	// $主包_$主函数
}
