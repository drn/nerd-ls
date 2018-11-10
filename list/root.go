package list

import (
  "os"
  "log"
  "sync"
  "io/ioutil"
  "github.com/drn/nerd-ls/node"
)

var waitGroup sync.WaitGroup
var mutex = sync.RWMutex{}

// Fetch - Fetch List representing current directory
func Fetch(dir string, options map[string]bool) []node.Node {
  files, err := ioutil.ReadDir(dir)
  if err != nil { log.Fatal(err) }

  length := len(files)
  if options["all"] { length += 2 }

  nodes := make([]node.Node, length)

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

    waitGroup.Add(1)
    go func(file os.FileInfo, index int) {
      defer waitGroup.Done()
      mutex.Lock()
      nodes[index] = node.New(file)
      mutex.Unlock()
    }(files[i], index)

    index++
  }
  waitGroup.Wait()

  return nodes[:index]
}
