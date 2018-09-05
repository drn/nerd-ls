package node

import (
  "os"
  "log"
  "fmt"
  "path/filepath"
  "io/ioutil"
  "github.com/fatih/color"
)

var icons = map[string]rune{
  ".DS_Store":     '',
  ".bash_history": '',
  ".bash_profile": '',
  ".conf":         '',
  ".env":          '',
  ".git":          '',
  ".go":           '',
  ".js":           '',
  ".json":         '',
  ".md":           '',
  ".rb":           '',
  ".yml":          '',
  "dir":           '',
}

// Node - File or directory helper methods
type Node interface {
  Name() string
  Size() int
}

type node struct {
  file os.FileInfo
}

// Fetch - Fetch nodes in currently directory
func Fetch() []Node {
  files, err := ioutil.ReadDir(".")
  if err != nil { log.Fatal(err) }

  nodes := make([]Node, len(files))

  for i:=0; i<len(files); i++ {
    nodes[i] = node{files[i]}
  }

  return nodes
}

func (n node) Name() string {
  return n.color().SprintFunc()(n.name())
}

func (n node) name() string {
  suffix := ""
  if n.file.IsDir() { suffix = "/" }
  return fmt.Sprintf("%c  %s%s   ", n.icon(), n.file.Name(), suffix)
}

func (n node) icon() rune {
  if n.file.IsDir() { return icons["dir"] }
  icon := icons[filepath.Ext(n.file.Name())]
  if icon == 0 { return ' ' }
  return icon
}

func (n node) color() *color.Color {
  if n.file.IsDir() { return color.New(color.FgCyan, color.Bold) }
  return color.New(color.FgWhite)
}

func (n node) Size() int {
  return len([]rune(n.name()))
}
