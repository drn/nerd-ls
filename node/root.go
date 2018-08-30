package node

import (
  "os"
  "log"
  "fmt"
  "io/ioutil"
)

// Fetch - Fetch nodes in currently directory
func Fetch() []Node {
  files, err := ioutil.ReadDir(".")
  if err != nil { log.Fatal(err) }

  nodes := make([]Node, len(files))

  for i:=0; i<len(files); i++ {
    nodes[i] = Node{files[i]}
  }

  return nodes
}

type Node struct {
  file os.FileInfo
}

func (node Node) Name() string {
  name := node.file.Name()
  if node.file.IsDir() {
    return fmt.Sprintf(" %s/ ", name)
  }
  return fmt.Sprintf("  %s ", name)
}

func (node Node) Size() int {
  name := node.file.Name()
  if node.file.IsDir() {
    return len(name) + 4
  }
  return len(name) + 3
}
