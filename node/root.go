package node

import (
  "os"
  "fmt"
  "time"
  "syscall"
  "os/user"
  "github.com/fatih/color"
  "github.com/drn/nerd-ls/icon"
)

// Node - Contains all info necessary to render file or directory
type Node struct {
  Name string
  Length int
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
  name := rawName(file)
  length := len([]rune(name))
  name = colorize(file, name)

  stat := file.Sys().(*syscall.Stat_t)

  uid := fmt.Sprint(stat.Uid)
  gid := fmt.Sprint(stat.Gid)

  fileUser, _ := user.LookupId(uid)
  fileGroup, _ := user.LookupGroupId(gid)

  nlink := int(stat.Nlink)

  time := stat.Ctimespec

  symlink := ""
  if file.Mode() & os.ModeSymlink == os.ModeSymlink {
    symlink, _ = os.Readlink(file.Name())
  }

  return Node{
    name,
    length,
    nlink,
    file.Mode().String(),
    int(file.Size()),
    fileUser.Username,
    fileGroup.Name,
    timespecToTime(time),
    symlink,
  }
}

func timespecToTime(ts syscall.Timespec) time.Time {
  return time.Unix(int64(ts.Sec), int64(ts.Nsec))
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
