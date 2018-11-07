package format

import (
  "fmt"
  "strconv"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/node"
)

// Long - Format listing in long format.
func Long(nodes []node.Node) {
  for _, node := range nodes {
    fmt.Printf(
      "%s %d %s %s %d %s\n",
      formatMode(node.Mode),
      node.LinkCount,
      node.User,
      node.Group,
      node.Size,
      node.Name,
    )
  }
}

func formatMode(mode string) string {
  runes := []rune(mode)

  return fmt.Sprintf(
    "%s%s%s%s%s%s%s%s%s%s",
    colorize(runes[0], color.New(color.FgWhite, color.Bold)),
    colorize(runes[1], color.New(color.FgGreen)),
    colorize(runes[2], color.New(color.FgGreen)),
    colorize(runes[3], color.New(color.FgGreen)),
    colorize(runes[4], color.New(color.FgYellow)),
    colorize(runes[5], color.New(color.FgYellow)),
    colorize(runes[6], color.New(color.FgYellow)),
    colorize(runes[7], color.New(color.FgRed)),
    colorize(runes[8], color.New(color.FgRed)),
    colorize(runes[9], color.New(color.FgRed)),
  )
}

func colorize(mode rune, color *color.Color) string {
  if mode == '-' { return "-" }
  return color.Sprintf("%c", mode)
}

func intLength(size int) int {
  return len([]rune(strconv.Itoa(size)))
}
