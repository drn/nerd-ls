package format

import (
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/node"
)

func nodeColor(node node.Node) *color.Color {
  if !node.IsDir { return color.New(color.FgWhite) }
  return color.New(color.FgCyan, color.Bold)
}
