package fatal

import (
	"fmt"
)

func OnError(err error, args ...any) {
	if err == nil {
		return
	}
	message := fmt.Sprint(append(args, err)...)
	LogError(message)
}

func OnErrorf(err error, format string, args ...any) {
	if err == nil {
		return
	}
	LogErrorf("%s: %v", fmt.Sprintf(format, args...), err)
}
