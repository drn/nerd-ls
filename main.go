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

func main() {
  flag.Parse()

  if *all { fmt.Println("all") } else { fmt.Println("not all") }

  width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
  if err != nil {
    fmt.Printf("error getting terminal dimensions\n")
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }

  nodes := node.Fetch()

  count := 0
  maxSize := maxSize(nodes)

  fmt.Printf("max size %d\n", maxSize)

  // dirColor := color.New(color.FgCyan, color.Bold).SprintFunc()

  for _, node := range nodes {
    count += maxSize
    if count >= width {
      fmt.Println()
      count = 0
    }
    name := node.Name()
    padding := maxSize - node.Size()

    fmt.Printf("%s%s", name, strings.Repeat(" ", padding))
  }
  fmt.Println()
}

func maxSize(nodes []node.Node) int {
  maxSize := 0
  for _, node := range nodes {
    size := node.Size()
    if maxSize < size { maxSize = size }
  }
  return maxSize
}
