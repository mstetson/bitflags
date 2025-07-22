package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/gonuts/commander"
)

var setCmd = &commander.Command{
	Run:       runSet,
	UsageLine: "set [flag-to-set ...]",
	Short:     "set all the given flag values by ORing them together",
	Flag:      *flag.NewFlagSet("split", flag.ExitOnError),
}

func runSet(cmd *commander.Command, args []string) error {
	if len(args) < 1 {
		cmd.Usage()
		return fmt.Errorf("wrong number of args: %d", len(args))
	}

	var f uint64
	for _, arg := range args {
		if strings.Contains(arg, "=") {
			// Ignore anything before =.
			s := strings.SplitN(arg, "=", 2)
			arg = s[1]
		}
		i, err := strconv.ParseUint(arg, 0, 64)
		if err != nil {
			return err
		}
		f |= i
	}

	fmt.Println(f)

	return nil
}
