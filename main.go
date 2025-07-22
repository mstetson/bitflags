package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gonuts/commander"
)

var cmd = &commander.Command{
	UsageLine: "bitflags command",
	Short:     "manipulate bit flags on the command line",
	Subcommands: []*commander.Command{
		setCmd,
		splitCmd,
	},
}

func main() {
	err := cmd.Dispatch(context.Background(), os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, os.Args[0]+":", err)
		os.Exit(1)
	}
}
