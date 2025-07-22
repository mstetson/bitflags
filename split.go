package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/gonuts/commander"
)

var splitCmd = &commander.Command{
	Run:       runSplit,
	UsageLine: "split <flag-value>",
	Short:     "print in hex each of the single-bit flags set in flag-value",
	Flag:      *flag.NewFlagSet("split", flag.ExitOnError),
}

func runSplit(cmd *commander.Command, args []string) error {
	if len(args) != 1 {
		cmd.Usage()
		return fmt.Errorf("wrong number of args: %d", len(args))
	}

	i, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		return err
	}

	for f := uint64(1); f <= i; f <<= 1 {
		if f&i == f {
			fmt.Printf("0x%x\n", f)
		}
	}

	return nil
}
