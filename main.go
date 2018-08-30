package main

import (
  "os"
  "fmt"
  "log"
  "flag"
  "bytes"
  "strings"
  "io/ioutil"
  "github.com/fatih/color"
  "golang.org/x/crypto/ssh/terminal"
)

var all = flag.Bool(
  "a",
  false,
  "Include directory entries whose names begin with a dot (.).",
)

func main() {
  flag.Parse()

  if *all { fmt.Println("all") } else { fmt.Println("not all") }

  width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
  if err != nil {
    fmt.Printf("error getting terminal dimensions\n")
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }

  // fmt.Println(width)

  var buffer bytes.Buffer

  files, err := ioutil.ReadDir("./")
  if err != nil { log.Fatal(err) }

  count := 0
  maxSize := maxSize(files)

  dirColor := color.New(color.FgCyan, color.Bold).SprintFunc()

  for _, f := range files {
    // if this is a .dotfile and '-a' is not specified, skip it
    if !*all && []rune(f.Name())[0] == rune('.') {
      continue
    }

    name := f.Name()
    size := len(name)

    difference := maxSize - size
    if f.IsDir() { difference-- }

    if count + maxSize + 3 > width {
      buffer.WriteString("\n")
      count = 0
    }

    count += maxSize + 1

    if f.IsDir() {
      buffer.WriteString(" ")
      buffer.WriteString(dirColor(name))
      buffer.WriteRune('/')
    } else {
      buffer.WriteString("  ")
      buffer.WriteString(name)
    }
    buffer.WriteString(strings.Repeat(" ", difference))
    buffer.WriteRune(' ')
  }

  fmt.Println(buffer.String())
}

func maxSize(files []os.FileInfo) int {
  maxSize := 0

  for _, f := range files {
    // if this is a .dotfile and '-a' is not specified, skip it
    if !*all && []rune(f.Name())[0] == rune('.') {
      continue
    }
    name := f.Name()
    size := len(name) + 2
    if f.IsDir() { size++ }
    if maxSize < size { maxSize = size }
  }

  return maxSize
}
