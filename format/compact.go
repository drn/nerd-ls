package format

import (
  "os"
  "fmt"
  "strings"
  "github.com/drn/nerd-ls/node"
  "github.com/drn/nerd-ls/list"
  "golang.org/x/crypto/ssh/terminal"
)

func Compact(list list.List) {
  width := width()
  count := 0

  maxLength := maxLength(list.Nodes)

  padding := 0
  for _, node := range list.Nodes {
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
    length := node.Length
    if maxLength < length { maxLength = length }
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
