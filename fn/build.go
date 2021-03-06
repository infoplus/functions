package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func build() cli.Command {
	cmd := buildcmd{}
	flags := append([]cli.Flag{}, cmd.flags()...)
	return cli.Command{
		Name:   "build",
		Usage:  "build function version",
		Flags:  flags,
		Action: cmd.build,
	}
}

type buildcmd struct {
	verbose bool
}

func (b *buildcmd) flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:        "v",
			Usage:       "verbose mode",
			Destination: &b.verbose,
		},
	}
}

// build will take the found valid function and build it
func (b *buildcmd) build(c *cli.Context) error {
	verbwriter := verbwriter(b.verbose)

	path, err := os.Getwd()
	if err != nil {
		return err
	}
	fn, err := findFuncfile(path)
	if err != nil {
		return err
	}

	fmt.Fprintln(verbwriter, "building", fn)
	ff, err := buildfunc(verbwriter, fn)
	if err != nil {
		return err
	}

	fmt.Printf("Function %v built successfully.\n", ff.FullName())
	return nil
}
