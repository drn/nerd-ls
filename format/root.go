package format

import (
  "regexp"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/node"
  "github.com/drn/nerd-ls/list"
)

var errorRegex = regexp.MustCompile(`.*\.orig$`)
var ignoreRegex = regexp.MustCompile(`^.DS_Store$`)

// Display - Runs display logic based on input address and options
func Display(address string, options map[string]interface{}) {
  nodes := list.Fetch(address, options)

  if options["long"].(bool) {
    Long(nodes, options)
  } else {
    Compact(nodes, options)
  }
}

func nodeColor(node node.Node) *color.Color {
  if node.Symlink != "" { return color.New(color.FgMagenta) }
  if node.IsDir { return color.New(color.FgCyan, color.Bold) }
  if ignoreRegex.MatchString(node.Name) { return color.New(color.FgBlack, color.Bold) }
  if errorRegex.MatchString(node.Name) { return color.New(color.FgRed) }
  return color.New(color.FgWhite)
}
