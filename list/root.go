package list

import (
	"github.com/drn/nerd-ls/node"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// Fetch - Fetch List representing current directory
func Fetch(dir string, options map[string]interface{}) []node.Node {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	if options["size-sort"].(bool) {
		sort.Slice(files, func(i, j int) bool {
			return files[i].Size() > files[j].Size()
		})
	}

	nodes := make([]node.Node, len(files)+2)

	index := 0
	if options["all"].(bool) {
		file, _ := os.Stat(".")
		nodes[0] = node.New(file)
		file, _ = os.Stat("..")
		nodes[1] = node.New(file)
		index += 2
	}

	for i := 0; i < len(files); i++ {
		if !options["all"].(bool) && []rune(files[i].Name())[0] == '.' {
			continue
		}
		nodes[index] = node.New(files[i])
		index++
	}

	return nodes[:index]
}
