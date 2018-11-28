package format

import (
  "regexp"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/node"
)

var errorRegex = regexp.MustCompile(`.*\.orig$`)
var ignoreRegex = regexp.MustCompile(`^.DS_Store$`)

func nodeColor(node node.Node) *color.Color {
  if node.Symlink != "" { return color.New(color.FgMagenta) }
  if node.IsDir { return color.New(color.FgCyan, color.Bold) }
  if ignoreRegex.MatchString(node.Name) { return color.New(color.FgBlack, color.Bold) }
  if errorRegex.MatchString(node.Name) { return color.New(color.FgRed) }
  return color.New(color.FgWhite)
}
