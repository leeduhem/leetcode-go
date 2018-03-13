package main

import (
	"os"
	"fmt"
	"flag"
	"strconv"
)

// Global flags
var Verbose bool

func init() {
	flag.BoolVar(&Verbose, "verbose", false, "Be verbose")
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [global options] command [options]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%-16s%s\n", "initiate", "Initiate source code directory")
	fmt.Fprintf(os.Stderr, "\t%-16s%s\n", "shift",  "Shift an existed solution")

	fmt.Fprintf(os.Stderr, "\nGlobal options:\n")
	flag.CommandLine.PrintDefaults()
}

func cmdUsage(f *flag.FlagSet, usage, cmd string) {
	fmt.Fprintf(f.Output(), usage, cmd)
	// TODO: Print usage message of options ONLY when there are some
	fmt.Fprintf(f.Output(), "\nCommand options:\n")
	f.PrintDefaults()
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

	case "initiate":
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

		doInitiateCmd(id)

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

		doShiftCmd(id)
	}
}

func doInitiateCmd(id int) {
	fmt.Println("Doing Initiate Command")
}

func doShiftCmd(id int) {
	fmt.Println("Doing Shift Command")
}
