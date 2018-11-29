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

  if opts.Long {
    Long(nodes, formatOptions(opts))
  } else {
    Compact(nodes, formatOptions(opts))
  }
}

func formatOptions(opts options.Options) map[string]int {
  icon := 0; if opts.Icon { icon = 1 }
  return map[string]int{
    "icon": icon,
  }
}

func nodeColor(node node.Node) *color.Color {
  if node.Symlink != "" { return color.New(color.FgMagenta) }
  if node.IsDir { return color.New(color.FgCyan, color.Bold) }
  if ignoreRegex.MatchString(node.Name) { return color.New(color.FgBlack, color.Bold) }
  if errorRegex.MatchString(node.Name) { return color.New(color.FgRed) }
  return color.New(color.FgWhite)
}
