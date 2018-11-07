package list

import (
  "os"
  "log"
  "io/ioutil"
  "github.com/drn/nerd-ls/node"
)

// List - Contains all info necessary to render list of nodes
type List struct {
  Nodes []node.Node
}

// Fetch - Fetch List representing current directory
func Fetch(dir string, options map[string]bool) List {
  files, err := ioutil.ReadDir(dir)
  if err != nil { log.Fatal(err) }

  nodes := make([]node.Node, len(files)+2)

  index := 0
  if options["all"] {
    file, _ := os.Stat(".")
    nodes[0] = node.New(file)
    file, _ = os.Stat("..")
    nodes[1] = node.New(file)
    index += 2
  }

  for i:=0; i<len(files); i++ {
    if !options["all"] && []rune(files[i].Name())[0] == '.' { continue }
    nodes[index] = node.New(files[i])
    index++
  }

  nodes = nodes[:index]

  return List{ nodes }
}