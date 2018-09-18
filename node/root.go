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
  Length int
  Mode string
  Size int
}

// Fetch - Fetch nodes in currently directory
func Fetch(options map[string]bool) []Node {
  files, err := ioutil.ReadDir(".")
  if err != nil { log.Fatal(err) }

  nodes := make([]Node, len(files)+2)

  index := 0
  if options["all"] {
    file, _ := os.Stat(".")
    nodes[0] = new(file)
    file, _ = os.Stat("..")
    nodes[1] = new(file)
    index += 2
  }

  for i:=0; i<len(files); i++ {
    if !options["all"] && []rune(files[i].Name())[0] == '.' { continue }
    nodes[index] = new(files[i])
    index++
  }

  return nodes[:index]
}

func new(file os.FileInfo) Node {
  name := rawName(file)
  length := len([]rune(name))
  name = colorize(file, name)
  return Node{name, length, file.Mode().String(), int(file.Size())}
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
