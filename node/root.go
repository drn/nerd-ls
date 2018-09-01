package node

import (
  "os"
  "log"
  "fmt"
  "io/ioutil"
  "github.com/fatih/color"
)

// Node - File or directory helper methods
type Node interface {
  Name() string
  Size() int
}

type node struct {
  file os.FileInfo
  name string
  size int
}

// Fetch - Fetch nodes in currently directory
func Fetch() []Node {
  files, err := ioutil.ReadDir(".")
  if err != nil { log.Fatal(err) }

  nodes := make([]Node, len(files))

  for i:=0; i<len(files); i++ {
    nodes[i] = node{files[i], "", 0}
  }

  return nodes
}

func (n node) Name() string {
  if n.name != "" { return n.name }
  n.name = fmt.Sprintf("%c %s/ ", n.icon(), n.color()(n.file.Name()))
  return n.name
}

func (n node) icon() rune {
  if n.file.IsDir() { return '' }
  return ' '
}

func (n node) color() func(a ...interface{}) string {
  return color.New(color.FgCyan, color.Bold).SprintFunc()
}

func (n node) Size() int {
  if n.size != 0 { return n.size }
  n.size = len([]rune(n.Name()))
  return n.size
}
