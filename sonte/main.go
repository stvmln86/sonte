package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Command interface {
	Name() string
	Flags() *flag.FlagSet
	Run(io.Writer) error
}

type MockCommand struct {
	Flag string // -f
}

func (c *MockCommand) Name() string {
	return "mock"
}

func (c *MockCommand) Flags() *flag.FlagSet {
	fset := flag.NewFlagSet("mock", flag.ContinueOnError)
	fset.StringVar(&c.Flag, "f", "", "string flag")
	return fset
}

func (c *MockCommand) Run(w io.Writer) error {
	fmt.Fprintf(w, "mock command, -f flag is %q\n", c.Flag)
	return nil
}

func main() {
	commands := map[string]Command{
		"mock": new(MockCommand),
	}

	if len(os.Args) < 2 {
		log.Fatal("no arguments")
	}

	name := os.Args[1]
	args := os.Args[2:]
	comm, ok := commands[name]
	if !ok {
		log.Fatalf("command %q does not exist", name)
	}

	if err := comm.Flags().Parse(args); err != nil {
		log.Fatalf("error parsing flags for command %q: %v", name, err)
	}

	if err := comm.Run(os.Stdout); err != nil {
		log.Fatalf("error running command %q: %v", name, err)
	}
}
