package main

import (
	"github.com/alecthomas/kong"
)

func main() {
	opts := &flags.Blah{}
	kongCtx := kong.Parse(opts)
}
