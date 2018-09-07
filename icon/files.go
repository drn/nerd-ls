package icon

import (
  "path/filepath"
)

// ForFile - Return rune icon corresponding to input file name
func ForFile(name string) rune {
  ext := filepath.Ext(name)
  if len(ext) == 0 { return files["file"] }
  icon := files[ext[1:]]
  if icon == 0 { return files["file"] }
  return icon
}

var files = map[string]rune{
  "ai":           '\ue7b4',
  "android":      '\ue70e',
  "apple":        '\uf179',
  "audio":        '\uf001',
  "avro":         '\ue60b',
  "c":            '\ue61e',
  "clj":          '\ue768',
  "coffee":       '\uf0f4',
  "conf":         '\ue615',
  "cpp":          '\ue61d',
  "css":          '\ue749',
  "d":            '\ue7af',
  "dart":         '\ue798',
  "db":           '\uf1c0',
  "diff":         '\uf440',
  "doc":          '\uf1c2',
  "docker":       '\uf308',
  "ebook":        '\ue28b',
  "env":          '\uf462',
  "epub":         '\ue28a',
  "erl":          '\ue7b1',
  "file":         '\uf15b',
  "font":         '\uf031',
  "gform":        '\uf298',
  "git":          '\uf1d3',
  "go":           '\ue626',
  "gruntfile.js": '\ue74c',
  "hs":           '\ue777',
  "html":         '\uf13b',
  "image":        '\uf1c5',
  "iml":          '\ue7b5',
  "java":         '\ue204',
  "js":           '\ue74e',
  "json":         '\ue60b',
  "jsx":          '\ue7ba',
  "less":         '\ue758',
  "log":          '\uf18d',
  "lua":          '\ue620',
  "md":           '\uf48a',
  "mustache":     '\ue60f',
  "npmignore":    '\ue71e',
  "pdf":          '\uf1c1',
  "php":          '\ue73d',
  "pl":           '\ue769',
  "ppt":          '\uf1c4',
  "psd":          '\ue7b8',
  "py":           '\ue606',
  "r":            '\uf25d',
  "rb":           '\ue21e',
  "rdb":          '\ue76d',
  "rss":          '\uf09e',
  "rubydoc":      '\ue73b',
  "sass":         '\ue603',
  "scala":        '\ue737',
  "shell":        '\uf489',
  "sqlite3":      '\ue7c4',
  "styl":         '\ue600',
  "tex":          '\ue600',
  "ts":           '\ue628',
  "twig":         '\ue61c',
  "txt":          '\uf15c',
  "video":        '\uf03d',
  "vim":          '\ue62b',
  "windows":      '\uf17a',
  "xls":          '\uf1c3',
  "xml":          '\ue619',
  "yarn.lock":    '\ue718',
  "yml":          '\uf481',
  "zip":          '\uf410',
}
