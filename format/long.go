package format

import (
  "fmt"
  "math"
  "time"
  "strings"
  "strconv"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/util"
  "github.com/drn/nerd-ls/node"
  "github.com/drn/nerd-ls/humanize"
)

// Long - Format listing in long format.
func Long(nodes []node.Node) {
  displaySummary(nodes)

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
      length := len(util.StripColor(values[i][j]))
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

func displaySummary(nodes []node.Node) {
  dirCount := 0
  dirSize := 0
  fileCount := 0
  fileSize := 0
  for _, node := range nodes {
    if node.IsDir {
      dirCount++
      dirSize += node.Size
    } else {
      fileCount++
      fileSize += node.Size
    }
  }
  fmt.Printf(
    "%s (%s @ %s) & %s (%s @ %s)\n",
    color.New(color.FgWhite, color.Bold).Sprint("files"),
    color.New(color.FgMagenta, color.Bold).Sprint(fileCount),
    formatSize(fileSize),
    color.New(color.FgWhite, color.Bold).Sprint("directories"),
    color.New(color.FgMagenta, color.Bold).Sprint(dirCount),
    formatSize(dirSize),
  )
}

func extractValues(node node.Node) []string {
  return []string{
    formatMode(node.Mode),
    strconv.Itoa(node.LinkCount),
    fmt.Sprintf("%s ", node.User),
    fmt.Sprintf("%s  ", node.Group),
    formatSize(node.Size),
    formatTime(node),
    fmt.Sprintf(" %c", node.Icon),
    formatName(node),
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

func formatTime(node node.Node) string {
  var baseColor *color.Color
  if util.IsToday(node.Time) {
    baseColor = color.New(color.FgMagenta)
  } else {
    baseColor = color.New(color.FgCyan)
  }

  var timeOrYear string
  if time.Now().Year() == node.Time.Year() {
    timeOrYear = fmt.Sprintf(
      "%s:%s",
      baseColor.Sprintf("%02d", node.Time.Hour()),
      baseColor.Sprintf("%02d", node.Time.Minute()),
    )
  } else {
    timeOrYear = color.New(
      color.FgMagenta,
      color.Bold,
    ).Sprintf(
      " %d",
      node.Time.Year(),
    )
  }

  return fmt.Sprintf(
    "%s %s %s",
    baseColor.Sprint(node.Time.Month().String()[:3]),
    baseColor.Sprintf("%2d", node.Time.Day()),
    timeOrYear,
  )
}

func formatMode(mode string) string {
  runes := []rune(mode)

  colorize := func(mode rune, color *color.Color) string {
    if mode == '-' { return "-" }
    return color.Sprintf("%c", mode)
  }

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

func formatName(node node.Node) string {
  baseColor := nodeColor(node)

  if node.Symlink == "" { return baseColor.Sprintf(" %s", node.Name) }

  return fmt.Sprintf(
    " %s %s%s%s %s",
    baseColor.Sprint(node.Name),
    color.New(color.FgMagenta, color.Bold).Sprint("➤"),
    color.New(color.FgBlue, color.Bold).Sprint("➤"),
    color.New(color.FgMagenta, color.Bold).Sprint("➤"),
    util.ShortenPath(node.Symlink),
  )
}
