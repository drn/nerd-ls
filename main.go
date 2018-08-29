package main

import (
  "os"
  "fmt"
  "log"
  "bytes"
  "strings"
  "io/ioutil"
  "github.com/fatih/color"
  "golang.org/x/crypto/ssh/terminal"
)

func main() {
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
    if []rune(f.Name())[0] == rune('.') {
      continue
    }

    name := f.Name()
    size := len(name)

    difference := maxSize - size

    if count + maxSize + 1 > width {
      buffer.WriteString("\n")
      count = 0
    }

    count += maxSize + 1

    if f.IsDir() {
      buffer.WriteString(dirColor(name))
    } else {
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
    if []rune(f.Name())[0] == rune('.') {
      continue
    }
    name := f.Name()
    size := len(name)
    if maxSize < size { maxSize = size }
  }

  return maxSize
}
