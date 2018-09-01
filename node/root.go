package node

import (
  "os"
  "log"
  "fmt"
  "io/ioutil"
)

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
  name := n.file.Name()
  if n.file.IsDir() {
    return fmt.Sprintf(" %s/ ", name)
  }
  return fmt.Sprintf("  %s ", name)
}

func (n node) Size() int {
  return len([]rune(n.Name()))
}
