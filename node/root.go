package node

import (
  "os"
  "fmt"
  "syscall"
  "os/user"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/icon"
)

// Node - Contains all info necessary to render file or directory
type Node struct {
  Name string
  Length int
  Mode string
  Size int
  User string
  Group string
}

// New - Initializes Node with os.FileInfo
func New(file os.FileInfo) Node {
  name := rawName(file)
  length := len([]rune(name))
  name = colorize(file, name)

  uid := fmt.Sprint(file.Sys().(*syscall.Stat_t).Uid)
  gid := fmt.Sprint(file.Sys().(*syscall.Stat_t).Gid)

  fileUser, _ := user.LookupId(uid)
  fileGroup, _ := user.LookupGroupId(gid)

  return Node{
    name,
    length,
    file.Mode().String(),
    int(file.Size()),
    fileUser.Username,
    fileGroup.Name,
  }
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
