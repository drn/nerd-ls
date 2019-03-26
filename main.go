package main

import (
	"fmt"
	"github.com/drn/nerd-ls/format"
	"github.com/drn/nerd-ls/options"
	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
	"os"
)

var opts options.Options

func main() {
	args, err := flags.ParseArgs(&opts, os.Args)
	if flags.WroteHelp(err) {
		return
	}
	if err != nil {
		fmt.Println()
		flags.ParseArgs(&opts, []string{"--help"})
		os.Exit(1)
	}

	var dirs []string
	if len(args) == 1 {
		dirs = []string{"."}
	} else {
		dirs = args[1:]
	}

	for i, dir := range dirs {
		if len(dirs) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf(
				"%s:\n",
				color.New(color.FgMagenta, color.Bold).Sprint(dir),
			)
		}

		format.Display(dir, options.Parse(opts))
	}
}
