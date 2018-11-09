package format

import (
  "fmt"
  "regexp"
  "strings"
  "strconv"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/node"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]" +
             "*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=>" +
             "<~]))"
var ansiRegex = regexp.MustCompile(ansi)

// Long - Format listing in long format.
func Long(nodes []node.Node) {
  // populate values
  values := make([][]string, len(nodes))
  for i := range values {
    node := nodes[i]
    values[i] = []string{
      formatMode(node.Mode),
      strconv.Itoa(node.LinkCount),
      fmt.Sprintf("%s ", node.User),
      fmt.Sprintf("%s ", node.Group),
      strconv.Itoa(node.Size),
      node.Time.Month().String()[:3],
      fmt.Sprintf("%2d", node.Time.Day()),
      fmt.Sprintf("%02d:%02d", node.Time.Hour(), node.Time.Minute()),
      fmt.Sprintf(" %s", node.Name),
    }
  }

  // calculate lengths and max lengths
  lengths := make([][]int, len(values))
  maxLengths := make([]int, len(values[0]))
  for i := range values {
    lengths[i] = make([]int, len(values[i]))
    for j := range values[i] {
      length := len(strip(values[i][j]))
      lengths[i][j] = length
      if length > maxLengths[j] {
        maxLengths[j] = length
      }
    }
  }

  // output padded values
  length := len(values[0])
  for i := range values {
    for j := range values[i] {
      // pad all attributes except the last
      if j < length - 1 {
        padding := maxLengths[j] - lengths[i][j]
        fmt.Print(strings.Repeat(" ", padding))
      }
      fmt.Printf("%s ", values[i][j])
    }
    fmt.Print("\n")
  }
}

// strips ANSI color codes from string
func strip(str string) string {
  return ansiRegex.ReplaceAllString(str, "")
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
