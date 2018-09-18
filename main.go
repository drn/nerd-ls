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
    fmt.Printf("%s  %s\n", node.Mode, node.Name)
  }
}

func displayCompact(nodes []node.Node) {
  width := width()
  count := 0
  maxSize := maxSize(nodes)

  padding := 0
  for _, node := range nodes {
    if padding > 0 { fmt.Print(strings.Repeat(" ", padding)) }

    count += maxSize
    if count >= width {
      fmt.Println()
      count = maxSize
    }

    padding = maxSize - node.Size

    fmt.Print(node.Name)
  }
  fmt.Println()
}

func maxSize(nodes []node.Node) int {
  maxSize := 0
  for _, node := range nodes {
    size := node.Size
    if maxSize < size { maxSize = size }
  }
  return maxSize
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
