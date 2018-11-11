package main

import (
  "os"
  "fmt"
  "github.com/drn/nerd-ls/list"
  "github.com/drn/nerd-ls/format"
  "github.com/jessevdk/go-flags"
)

var opts struct {
  All bool `short:"a" long:"all" description:"Include directory entries whose names begin with a dot (.)."`
  Long bool `short:"l" long:"long" description:"List in long format."`
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

  dir := "."
  if len(args) > 1 { dir = args[1] }

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
