package node

import (
  "os"
  "fmt"
  "time"
  "syscall"
  "os/user"
  "github.com/drn/nerd-ls/icon"
)

// Node - Contains all info necessary to render file or directory
type Node struct {
  Icon rune
  IsDir bool
  Name string
  LinkCount int
  Mode string
  Size int
  User string
  Group string
  Time time.Time
  Symlink string
}

// New - Initializes Node with os.FileInfo
func New(file os.FileInfo) Node {
  stat := file.Sys().(*syscall.Stat_t)

  uid := fmt.Sprint(stat.Uid)
  gid := fmt.Sprint(stat.Gid)

  fileUser, _ := user.LookupId(uid)
  fileGroup, _ := user.LookupGroupId(gid)

  nlink := int(stat.Nlink)

  symlink := ""
  if file.Mode() & os.ModeSymlink == os.ModeSymlink {
    symlink, _ = os.Readlink(file.Name())
  }

  return Node{
    fetchIcon(file),
    file.IsDir(),
    name(file),
    nlink,
    file.Mode().String(),
    int(file.Size()),
    fileUser.Username,
    fileGroup.Name,
    file.ModTime(),
    symlink,
  }
}

func name(file os.FileInfo) string {
  if !file.IsDir() { return file.Name() }
  return fmt.Sprintf("%s/", file.Name())
}

func fetchIcon(file os.FileInfo) rune {
  if file.IsDir() {
    return icon.ForFolder(file.Name())
  }
  return icon.ForFile(file.Name())
}
