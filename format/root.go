package format

import (
  "regexp"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/node"
)

var origRegex = regexp.MustCompile(`.*\.orig$`)

func nodeColor(node node.Node) *color.Color {
  if node.Symlink != "" { return color.New(color.FgMagenta) }
  if node.IsDir { return color.New(color.FgCyan, color.Bold) }
  if origRegex.MatchString(node.Name) { return color.New(color.FgRed) }
  return color.New(color.FgWhite)
}
