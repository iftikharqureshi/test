package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	x := 123
	m := map[string]int{"a": 1,
		"b": 2}
	spew.Dump(x)
	spew.Dump(m)
}
