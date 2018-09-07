// Initial config from:
// - https://github.com/athityakumar/colorls/blob/master/lib/yaml/folders.yaml
// - https://github.com/athityakumar/colorls/blob/master/lib/yaml/folder_aliases.yaml

package icon

// ForFolder - Return rune icon corresponding to input folder name
func ForFolder(name string) rune {
  alias := folderAliases[name]
  if alias != "" { name = alias }
  icon := folders[name]
  if icon == 0 { return folders["folder"] }
  return icon
}

var folderAliases = map[string]string{
  "bin":     "config",
  "include": "config",
}

var folders = map[string]rune {
  ".atom":        '\ue764',
  ".git":         '\ue5fb',
  ".github":      '\uf408',
  ".rvm":         '\ue21e',
  ".Trash":       '\uf1f8',
  ".vscode":      '\ue70c',
  "config":       '\ue5fc',
  "folder":       '\uf115',
  "hidden":       '\uf023',
  "lib":          '\uf121',
  "node_modules": '\ue718',
}
