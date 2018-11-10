package format

import (
  "fmt"
  "math"
  "time"
  "regexp"
  "strings"
  "strconv"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/node"
  "github.com/drn/nerd-ls/humanize"
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
    values[i] = extractValues(nodes[i])
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

func extractValues(node node.Node) []string {
  return []string{
    formatMode(node.Mode),
    strconv.Itoa(node.LinkCount),
    fmt.Sprintf("%s ", node.User),
    fmt.Sprintf("%s  ", node.Group),
    formatSize(node.Size),
    formatTime(node.Time),
    fmt.Sprintf(" %s", node.Name),
  }
}

func formatSize(sizeInt int) string {
  str := humanize.Bytes(sizeInt)
  size := float64(sizeInt)
  base := float64(1024)
  // less than 1K
  if size < base { return str }
  // less than 10M
  if size < math.Pow(base, 2) * 10 { return color.GreenString(str) }
  // less than 100M
  if size < math.Pow(base, 2) * 100 { return color.YellowString(str) }
  // less than 1G
  if size < math.Pow(base, 3) { return color.RedString(str) }
  // above 1G
  return color.New(color.FgRed, color.Bold).Sprint(str)
}

func formatTime(time time.Time) string {
  return fmt.Sprintf(
    "%s %s %02d:%02d",
    time.Month().String()[:3],
    fmt.Sprintf("%2d", time.Day()),
    time.Hour(),
    time.Minute(),
  )
}

// strips ANSI color codes from string
func strip(str string) string {
  return ansiRegex.ReplaceAllString(str, "")
}

func formatMode(mode string) string {
  runes := []rune(mode)

  return fmt.Sprintf(
    "%s%s%s%s%s%s%s%s%s%s ",
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
