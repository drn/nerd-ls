package util

import (
  "os"
  "time"
  "regexp"
  "strings"
  "golang.org/x/crypto/ssh/terminal"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]" +
             "*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=>" +
             "<~]))"
var ansiRegex = regexp.MustCompile(ansi)
var homePath = os.Getenv("HOME")

// StripColor - strips ANSI color codes from string
func StripColor(str string) string {
  return ansiRegex.ReplaceAllString(str, "")
}

// TerminalWidth - returns the width of the current terminal window
func TerminalWidth() int {
  width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
  if err == nil { return width }
  return 0
}

// IsToday - returns true if the input time is on today's date
func IsToday(input time.Time) bool {
  now := time.Now()
  if now.Year() != input.Year() { return false }
  if now.Month() != input.Month() { return false }
  if now.Day() != input.Day() { return false }
  return true
}

// ShortenPath - returns a shortened path
func ShortenPath(path string) string {
  if strings.Contains(path, homePath) {
    return strings.Replace(path, homePath, "~", 1)
  }
  return path
}
