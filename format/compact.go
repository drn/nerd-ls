package format

import (
  "fmt"
  "strings"
  "github.com/drn/nerd-ls/node"
  "github.com/drn/nerd-ls/util"
)

// Compact - Format listing in compact format.
func Compact(nodes []node.Node, options map[string]int) {
  width := util.TerminalWidth()

  if width == 0 {
    pipedDisplay(nodes)
  } else {
    compactDisplay(nodes, width, options)
  }
}

func pipedDisplay(nodes []node.Node) {
  for _, node := range nodes {
    fmt.Println(nodeColor(node).Sprint(node.Name))
  }
}

func compactDisplay(nodes []node.Node, width int, options map[string]int) {
  // determine max node length
  maxLength := 0
  for _, node := range nodes {
    length := len(node.Name)
    if maxLength < length { maxLength = length }
  }

  lengthPerNode := maxLength + 2 // name + 2 spaces
  if options["icon"] == 1 { lengthPerNode += 3 } // icon + 2 spaces
  nodesPerRow := width / lengthPerNode
  nodesLength := len(nodes)

  for i := 0; i < nodesLength; i ++ {
    node := nodes[i]

    // print node
    if options["icon"] == 0 {
      fmt.Print(nodeColor(node).Sprint(node.Name))
    } else {
      fmt.Printf(
        "%c  %s",
        node.Icon,
        nodeColor(node).Sprint(node.Name),
      )
    }

    if (i + 1) % nodesPerRow == 0 {
      // start a new row
      fmt.Println()
    } else {
      // print right padding
      fmt.Printf(
        "  %s",
        strings.Repeat(" ", maxLength - len(node.Name)),
      )
    }
  }

  // skip last linebreak if already printed
  if nodesLength % nodesPerRow != 0 {
    fmt.Println()
  }
}
