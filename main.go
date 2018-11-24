package main

import (
  "os"
  "fmt"
  "github.com/fatih/color"
  "github.com/jessevdk/go-flags"
  "github.com/drn/nerd-ls/list"
  "github.com/drn/nerd-ls/format"
)

var opts struct {
  All bool `short:"a" long:"all" description:"Include directory entries whose names begin with a dot (.)"`
  Long bool `short:"l" long:"long" description:"List in long format"`
  Icon bool `short:"i" long:"icon" description:"Display nerd-font icons"`
}

func main() {
  args, err := flags.ParseArgs(&opts, os.Args)
  if flags.WroteHelp(err) { return }
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
      if i > 0 { fmt.Println() }
      fmt.Printf(
        "%s:\n",
        color.New(color.FgMagenta, color.Bold).Sprint(dir),
      )
    }

    nodes := list.Fetch(
      dir,
      map[string]bool{
        "all": opts.All,
        "long": opts.Long,
      },
    )

    formatOptions := map[string]bool{"icon": opts.Icon}
    if opts.Long {
      format.Long(nodes, formatOptions)
    } else {
      format.Compact(nodes, formatOptions)
    }
  }
}
