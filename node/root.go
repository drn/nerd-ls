package node

import (
	"fmt"
	"github.com/drn/nerd-ls/icon"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"syscall"
	"time"
)

var whitespaceRegex = regexp.MustCompile(`\r|\r?\n`)

// Node - Contains all info necessary to render file or directory
type Node struct {
	Icon      rune
	IsDir     bool
	Name      string
	LinkCount int
	Mode      os.FileMode
	Size      int
	User      string
	Group     string
	Time      time.Time
	Symlink   string
}

// New - Initializes Node with os.FileInfo
func New(dir string, file os.FileInfo) Node {
	stat := file.Sys().(*syscall.Stat_t)

	uid := fmt.Sprint(stat.Uid)
	gid := fmt.Sprint(stat.Gid)

	fileUser, _ := user.LookupId(uid)
	fileUserName := uid
	if fileUser != nil {
		fileUserName = fileUser.Username
	}
	fileGroup, _ := user.LookupGroupId(gid)
	fileGroupName := gid
	if fileGroup != nil {
		fileGroupName = fileGroup.Name
	}

	nlink := int(stat.Nlink)

	symlink := ""
	if file.Mode()&os.ModeSymlink == os.ModeSymlink {
		symlink, _ = os.Readlink(fmt.Sprintf("%s/%s", dir, file.Name()))
	}

	return Node{
		fetchIcon(file),
		file.IsDir(),
		name(dir, file),
		nlink,
		file.Mode(),
		int(file.Size()),
		fileUserName,
		fileGroupName,
		file.ModTime(),
		symlink,
	}
}

func name(dir string, file os.FileInfo) string {
	baseName := file.Name()
	baseName = whitespaceRegex.ReplaceAllString(baseName, "?")
	if !file.IsDir() {
		return baseName
	}
	name := fmt.Sprintf("%s/", baseName)
	// inject name for current and parent directories
	if baseName == "." || baseName == ".." {
		fullpath, _ := filepath.Abs(fmt.Sprintf("%s/%s", dir, baseName))
		name = fmt.Sprintf("%s [%s]", name, filepath.Base(fullpath))
	}
	return name
}

func fetchIcon(file os.FileInfo) rune {
	if file.IsDir() {
		return icon.ForFolder(file.Name())
	}
	return icon.ForFile(file.Name())
}
