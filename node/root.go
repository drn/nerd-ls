package node

import (
  "os"
  "log"
  "fmt"
  "io/ioutil"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/icon"
)

// Node - Contains all info necessary to render file or directory
type Node struct {
  Name string
  Size int
}

// Fetch - Fetch nodes in currently directory
func Fetch(options map[string]bool) []Node {
  files, err := ioutil.ReadDir(".")
  if err != nil { log.Fatal(err) }

  count := 0
  for i:=0; i<len(files); i++ {
    if !options["all"] && []rune(files[i].Name())[0] == '.' { continue }
    count++
  }

  nodes := make([]Node, count)

  count = 0
  for i:=0; i<len(files); i++ {
    if !options["all"] && []rune(files[i].Name())[0] == '.' { continue }
    nodes[count] = new(files[i])
    count++
  }

  return nodes
}

func new(file os.FileInfo) Node {
  name := rawName(file)
  size := len([]rune(name))
  name = colorize(file, name)
  return Node{name, size}
}

func rawName(file os.FileInfo) string {
  suffix := ""
  if file.IsDir() { suffix = "/" }

  return fmt.Sprintf(
    "%c  %s%s   ",
    fetchIcon(file),
    file.Name(),
    suffix,
  )
}

func fetchIcon(file os.FileInfo) rune {
  if file.IsDir() {
    return icon.ForFolder(file.Name())
  }
  return icon.ForFile(file.Name())
}

func colorize(file os.FileInfo, name string) string {
  colorConfig := color.New(color.FgWhite)
  if file.IsDir() { colorConfig = color.New(color.FgCyan, color.Bold) }
  return colorConfig.SprintFunc()(name)
}
