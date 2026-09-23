package logx

import "fmt"

func Notice(message string) {
	Log("NOTICE", ColorCyan, message)
}

func Noticef(format string, args ...any) {
	Notice(fmt.Sprintf(format, args...))
}
