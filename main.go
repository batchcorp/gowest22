package main

import (
	"fmt"

	"github.com/alecthomas/kong"

	"github.com/batchcorp/gowest22/build/go/opts"
)

func main() {
	options := &opts.Base{}
	kongCtx := kong.Parse(options)

	if options.Debug {
		fmt.Println("Debug is enabled")
	}

	switch kongCtx.Command() {
	case "sub":
		fmt.Println("sub command")
	default:
		panic("unknown command")
	}
}
