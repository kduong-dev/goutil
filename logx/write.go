package logx

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

type Color string

const (
	ColorReset  Color = "\033[0m"
	ColorRed    Color = "\033[31m"
	ColorYellow Color = "\033[33m"
	ColorCyan   Color = "\033[36m"
)

var isTerminal = term.IsTerminal(int(os.Stderr.Fd()))

// Log writes level and message to stderr, prefixed with a timestamp and
// colored with color when stderr is a terminal.
func Log(level string, color Color, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	if isTerminal {
		fmt.Fprintf(os.Stderr, "%s %s%-6s%s %s\n", timestamp, color, level, ColorReset, message)
	} else {
		fmt.Fprintf(os.Stderr, "%s %-6s %s\n", timestamp, level, message)
	}
}
