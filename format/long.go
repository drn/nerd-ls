package format

import (
	"fmt"
	"math"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/drn/nerd-ls/humanize"
	"github.com/drn/nerd-ls/node"
	"github.com/drn/nerd-ls/util"
	"github.com/fatih/color"
)

// Long - Format listing in long format.
func Long(nodes []node.Node, options map[string]interface{}) {
	displaySummary(nodes)

	// populate values
	values := make([][]string, len(nodes))
	for i := range values {
		values[i] = extractValues(nodes[i], options)
	}

	// calculate lengths and max lengths
	lengths := make([][]int, len(values))
	maxLengths := make([]int, len(values[0]))
	for i := range values {
		lengths[i] = make([]int, len(values[i]))
		for j := range values[i] {
			length := len([]rune(util.StripColor(values[i][j])))
			lengths[i][j] = length
			if length > maxLengths[j] {
				maxLengths[j] = length
			}
		}
	}

	// output padded values
	for i := range values {
		for j := range values[i] {
			// Not (Link or Size) is left-aligned
			if !(j == 1 || j == 4) {
				fmt.Printf("%s ", values[i][j])
			}
			// padding
			fmt.Print(strings.Repeat(" ", maxLengths[j]-lengths[i][j]))
			// Link and Size right-aligned
			if j == 1 || j == 4 {
				fmt.Printf("%s ", values[i][j])
			}
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

func extractValues(node node.Node, options map[string]interface{}) []string {
	icon := ""
	if options["icon"].(bool) {
		icon = fmt.Sprintf("%c", node.Icon)
	}
	return []string{
		formatMode(node.Mode.String()),
		strconv.Itoa(node.LinkCount),
		formatUser(node.User),
		formatGroup(node.Group),
		formatSize(node.Size),
		formatTime(node),
		icon + formatName(node),
	}
}

func formatSize(sizeInt int) string {
	str := humanize.Bytes(sizeInt)
	size := float64(sizeInt)
	base := float64(1024)
	// less than 1K
	if size < base {
		return str
	}
	// less than 10M
	if size < math.Pow(base, 2)*10 {
		return color.GreenString(str)
	}
	// less than 100M
	if size < math.Pow(base, 2)*100 {
		return color.YellowString(str)
	}
	// less than 1G
	if size < math.Pow(base, 3) {
		return color.RedString(str)
	}
	// above 1G
	return color.New(color.FgRed, color.Bold).Sprint(str)
}

func formatUser(u string) string {
	cu, uerr := user.Current()
	if uerr != nil {
		return u
	}
	if u == cu.Username {
		return color.New(color.FgYellow, color.Bold).Sprint(u)
	}
	return u
}

func formatGroup(g string) string {
	cu, uerr := user.Current()
	if uerr != nil {
		return g
	}
	grp, gerr := user.LookupGroupId(cu.Gid)
	if gerr != nil {
		return g
	}
	if g == grp.Name {
		return color.New(color.FgYellow, color.Bold).Sprint(g)
	}
	return g
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
		if mode == '-' {
			return "-"
		}
		return color.Sprintf("%c", mode)
	}

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

func formatName(node node.Node) string {
	baseColor := nodeColor(node)

	if node.Symlink == "" {
		return baseColor.Sprintf(" %s", node.Name)
	}

	return fmt.Sprintf(
		" %s %s%s%s %s",
		baseColor.Sprint(node.Name),
		color.New(color.FgMagenta, color.Bold).Sprint("➤"),
		color.New(color.FgBlue, color.Bold).Sprint("➤"),
		color.New(color.FgMagenta, color.Bold).Sprint("➤"),
		util.ShortenPath(node.Symlink),
	)
}
