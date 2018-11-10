package format

import (
  "os"
  "fmt"
  "strings"
  "github.com/drn/nerd-ls/node"
  "golang.org/x/crypto/ssh/terminal"
)

// Compact - Format listing in compact format.
func Compact(nodes []node.Node) {
  width := width()

  if width == 0 {
    pipedDisplay(nodes)
  } else {
    compactDisplay(nodes, width)
  }
}

func pipedDisplay(nodes []node.Node) {
  for _, node := range nodes {
    fmt.Println(node.Name)
  }
}

func compactDisplay(nodes []node.Node, width int) {
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
    length := node.Length
    if maxLength < length { maxLength = length }
  }
  return maxLength
}

func width() int {
  width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
  if err == nil { return width }
  return 0
}
