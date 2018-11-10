package main

import (
  "flag"
  "github.com/drn/nerd-ls/list"
  "github.com/drn/nerd-ls/format"
)

var all = flag.Bool(
  "a",
  false,
  "Include directory entries whose names begin with a dot (.).",
)

var long = flag.Bool(
  "l",
  false,
  "(The lowercase letter ``ell''.)  List in long format.  (See " +
  "below.)  If the output is to a terminal, a total sum for all the " +
  "file sizes is output on a line before the long listing.",
)

func main() {
  flag.Parse()

  dir := "."
  if len(flag.Args()) >= 1 { dir = flag.Args()[0] }

  nodes := list.Fetch(
    dir,
    map[string]bool{
      "all": *all,
      "long": *long,
    },
  )

  if *long {
    format.Long(nodes)
  } else {
    format.Compact(nodes)
  }
}
