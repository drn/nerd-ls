package main

import (
  "os"
  "github.com/drn/nerd-ls/list"
  "github.com/drn/nerd-ls/format"
  "github.com/jessevdk/go-flags"
)

var opts struct {
  All bool `short:"a" long:"all" description:"Include directory entries whose names begin with a dot (.)."`
  Long bool `short:"l" long:"long" description:"List in long format."`
}

func main() {
  args, err := flags.ParseArgs(&opts, os.Args)
  if flags.WroteHelp(err) { return }

  dir := "."
  if len(args) > 1 { dir = args[1] }

  nodes := list.Fetch(
    dir,
    map[string]bool{
      "all": opts.All,
      "long": opts.Long,
    },
  )

  if opts.Long {
    format.Long(nodes)
  } else {
    format.Compact(nodes)
  }
}
