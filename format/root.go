package format

import (
  "regexp"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/node"
  "github.com/drn/nerd-ls/list"
  "github.com/drn/nerd-ls/options"
)

var errorRegex = regexp.MustCompile(`.*\.orig$`)
var ignoreRegex = regexp.MustCompile(`^.DS_Store$`)

// Display - Runs display logic based on input address and options
func Display(address string, opts options.Options) {
  nodes := list.Fetch(
    address,
    map[string]bool{
      "all": opts.All,
      "long": opts.Long,
    },
  )

  formatOptions := map[string]bool{"icon": opts.Icon}
  if opts.Long {
    Long(nodes, formatOptions)
  } else {
    Compact(nodes, formatOptions)
  }
}

func nodeColor(node node.Node) *color.Color {
  if node.Symlink != "" { return color.New(color.FgMagenta) }
  if node.IsDir { return color.New(color.FgCyan, color.Bold) }
  if ignoreRegex.MatchString(node.Name) { return color.New(color.FgBlack, color.Bold) }
  if errorRegex.MatchString(node.Name) { return color.New(color.FgRed) }
  return color.New(color.FgWhite)
}
