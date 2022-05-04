package main

import _ "embed"

//go:embed shoe-ascii-art.txt
var s string

func main() {
	println(s)
}
