package fatal

import (
	"fmt"
	"os"
	"runtime/debug"

	"goutil/logx"
)

func LogError(message string) {
	logx.Log("FATAL", logx.ColorRed, fmt.Sprintf("'%s'\n%s", message, debug.Stack()))
	os.Exit(1)
}

func LogErrorf(format string, args ...any) {
	LogError(fmt.Sprintf(format, args...))
}
