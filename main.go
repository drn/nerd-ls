package main

import (
  "os"
  "fmt"
  "flag"
  "strings"
  "github.com/drn/nerd-ls/node"
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

  nodes := node.Fetch(
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

func displayLong(nodes []node.Node) {
  for _, node := range nodes {
    fmt.Printf("%s  %d %s\n", node.Mode, node.Size, node.Name)
  }
}

func displayCompact(nodes []node.Node) {
  width := width()
  count := 0
  maxLength := maxLength(nodes)

  padding := 0
  for _, node := range nodes {
    if padding > 0 { fmt.Print(strings.Repeat(" ", padding)) }

    count += maxLength
    if count >= width {
      fmt.Println()
      count = maxLength
    }

    padding = maxLength - node.Length

    fmt.Print(node.Name)
  }
  fmt.Println()
}

func maxLength(nodes []node.Node) int {
  maxLength := 0
  for _, node := range nodes {
    size := node.Length
    if maxLength < size { maxLength = size }
  }
  return maxLength
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
