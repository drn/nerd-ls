package main

import (
  "os"
  "fmt"
  "flag"
  "strings"
  "strconv"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/list"
  "golang.org/x/crypto/ssh/terminal"
)

var all = flag.Bool(
  "a",
  false,
  "Include directory entries whose names begin with a dot (.).",
)

var long = flag.Bool(
  "l",
  false,
  "(The lowercase letter ``ell''.)  List in long format.  (See " +
  "below.)  If the output is to a terminal, a total sum for all the " +
  "file sizes is output on a line before the long listing.",
)

func main() {
  flag.Parse()

  dir := "."
  if len(flag.Args()) >= 1 { dir = flag.Args()[0] }

  nodes := list.Fetch(
    dir,
    map[string]bool{
      "all": *all,
      "long": *long,
    },
  )

  if *long {
    displayLong(nodes)
  } else {
    displayCompact(nodes)
  }
}

func displayLong(list list.List) {
  for _, node := range list.Nodes {
    padding := strings.Repeat(
      " ",
      intLength(list.MaxSize) - intLength(node.Size),
    )
    fmt.Printf(
      "%s %d %s %s %s%d %s\n",
      formatMode(node.Mode),
      node.LinkCount,
      node.User,
      node.Group,
      padding,
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

func displayCompact(list list.List) {
  width := width()
  count := 0

  padding := 0
  for _, node := range list.Nodes {
    if padding > 0 { fmt.Print(strings.Repeat(" ", padding)) }

    count += list.MaxLength
    if count >= width {
      fmt.Println()
      count = list.MaxLength
    }

    padding = list.MaxLength - node.Length

    fmt.Print(node.Name)
  }
  fmt.Println()
}

func width() int {
  width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
  if err != nil {
    fmt.Printf("error getting terminal dimensions\n")
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
  return width
}
