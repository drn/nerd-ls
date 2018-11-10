package util

import (
  "os"
  "regexp"
  "golang.org/x/crypto/ssh/terminal"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]" +
             "*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=>" +
             "<~]))"
var ansiRegex = regexp.MustCompile(ansi)

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
