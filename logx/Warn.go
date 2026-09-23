package logx

import "fmt"

func Warn(message string) {
	Log("WARN", ColorYellow, message)
}

func Warnf(format string, args ...any) {
	Warn(fmt.Sprintf(format, args...))
}
