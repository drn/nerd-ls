package format

import (
  "os"
  "fmt"
  "strings"
  "github.com/drn/nerd-ls/list"
  "golang.org/x/crypto/ssh/terminal"
)

func Compact(list list.List) {
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
