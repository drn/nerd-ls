package main

import (
  "os"
  "fmt"
  "flag"
  "strings"
  "strconv"
  "github.com/drn/nerd-ls/list"
  "golang.org/x/crypto/ssh/terminal"
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
  if len(os.Args) > 1 { dir = os.Args[1] }

  nodes := list.Fetch(
    dir,
    map[string]bool{
      "all": *all,
      "long": *long,
    },
  )

  if *long {
    displayLong(nodes)
  } else {
    displayCompact(nodes)
  }
}

func displayLong(list list.List) {
  for _, node := range list.Nodes {
    padding := strings.Repeat(
      " ",
      intLength(list.MaxSize) - intLength(node.Size),
    )
    fmt.Printf(
      "%s  %s%d  %s\n",
      node.Mode,
      padding,
      node.Size,
      node.Name,
    )
  }
}

func intLength(size int) int {
  return len([]rune(strconv.Itoa(size)))
}

func displayCompact(list list.List) {
  width := width()
  count := 0

  padding := 0
  for _, node := range list.Nodes {
    if padding > 0 { fmt.Print(strings.Repeat(" ", padding)) }

    count += list.MaxLength
    if count >= width {
      fmt.Println()
      count = list.MaxLength
    }

    padding = list.MaxLength - node.Length

    fmt.Print(node.Name)
  }
  fmt.Println()
}

func width() int {
  width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
  if err != nil {
    fmt.Printf("error getting terminal dimensions\n")
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
  return width
}
