package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	x := 123
	m := map[string]int{"a": 10,
		"b": 20}
	spew.Dump(x)
	spew.Dump(m)
}
