package format

import (
  "fmt"
  "bytes"
  "github.com/drn/nerd-ls/node"
)

// Tree - Format listing in tree format.
func Tree(nodes []node.Node, options map[string]interface{}) {
  depth := 0
  if _, ok := options["depth"]; ok { depth = options["depth"].(int) }
  if _, ok := options["opened"]; !ok { options["opened"] = []int{} }

  length := len(nodes)
  for i, node := range nodes {
    var opened []int
    if i < length - 1 {
      opened = append(options["opened"].([]int), depth)
    } else {
      opened = options["opened"].([]int)
    }

    fmt.Printf(
      "%s%s\n",
      prefix(opened, depth),
      nodeColor(node).Sprint(node.Name),
    )

    if node.IsDir && node.Symlink == "" {
      var address string
      if prefix, ok := options["prefix"].(string); ok {
        address = fmt.Sprintf("%s/%s", prefix, node.Name)
      } else {
        address = node.Name
      }

      Display(
        address,
        map[string]interface{}{
          "all": options["all"],
          "long": options["long"],
          "icon": options["icon"],
          "tree": options["tree"],
          "prefix": address,
          "depth": depth + 1,
          "opened": opened,
        },
      )
    }
  }
}

func prefix(opened []int, depth int) string {
  var prefix bytes.Buffer

  for i := 0; i < depth; i++ {
    if contains(opened, i) {
      prefix.WriteString("│  ")
    } else {
      prefix.WriteString("   ")
    }
  }

  if contains(opened, depth) {
    prefix.WriteString("├── ")
  } else {
    prefix.WriteString("└── ")
  }

  return prefix.String()
}

func contains(array []int, value int) bool {
  for _, element := range array {
    if element == value {
      return true
    }
  }
  return false
}
