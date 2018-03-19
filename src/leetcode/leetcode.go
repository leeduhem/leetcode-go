package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	cmdInit "leetcode/cmd/init"
	cmdShift "leetcode/cmd/shift"
)

// Global flags
var Verbose bool

func init() {
	flag.BoolVar(&Verbose, "verbose", false, "Be verbose")
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [global options] command [options]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%-16s%s\n", "init", "Initiate source code directory")
	fmt.Fprintf(os.Stderr, "\t%-16s%s\n", "shift", "Shift an existed solution")

	fmt.Fprintf(os.Stderr, "\nGlobal options:\n")
	flag.CommandLine.PrintDefaults()
}

func cmdUsage(f *flag.FlagSet, usage, cmd string) {
	fmt.Fprintf(f.Output(), usage, cmd)

	var flagNum int
	f.VisitAll(func(flg *flag.Flag) { flagNum++ })
	if flagNum > 0 {
		fmt.Fprintf(f.Output(), "\nCommand options:\n")
		f.PrintDefaults()
	}
}

func main() {
	flag.CommandLine.Usage = Usage
	flag.Parse()

	if flag.NArg() == 0 {
		Usage()
		os.Exit(2)
	}

	switch cmd := flag.Arg(0); cmd {
	default:
		Usage()
		os.Exit(2)

	case "init":
		cmdFlagSet := flag.NewFlagSet(cmd, flag.ExitOnError)
		cmdFlagSet.Usage = func() {
			cmdUsage(cmdFlagSet, "Usage: %s id\n", cmd)
		}

		cmdFlagSet.Parse(flag.Args()[1:])
		if cmdFlagSet.NArg() < 1 {
			cmdFlagSet.Usage()
			os.Exit(2)
		}

		id, err := strconv.Atoi(cmdFlagSet.Arg(0))
		if err != nil && id >= 1 {
			fmt.Fprintf(cmdFlagSet.Output(), "%s: invalid problem id\n", cmd)
			os.Exit(2)
		}

		cmdInit.SetVerbose(Verbose)
		cmdInit.Do(id)

	case "shift":
		cmdFlagSet := flag.NewFlagSet(cmd, flag.ExitOnError)
		cmdFlagSet.Usage = func() {
			cmdUsage(cmdFlagSet, "Usage: %s id\n", cmd)
		}

		cmdFlagSet.Parse(flag.Args()[1:])
		if cmdFlagSet.NArg() < 1 {
			cmdFlagSet.Usage()
			os.Exit(2)
		}

		id, err := strconv.Atoi(cmdFlagSet.Arg(0))
		if err != nil && id >= 1 {
			fmt.Fprintf(cmdFlagSet.Output(), "%s: invalid problem id\n", cmd)
			os.Exit(2)
		}

		cmdShift.SetVerbose(Verbose)
		cmdShift.Do(id)
	}
}
