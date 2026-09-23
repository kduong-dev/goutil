package fatal

func Unless(b bool, message string) {
	if !b {
		LogError(message)
	}
}

func Unlessf(b bool, format string, args ...any) {
	if !b {
		LogErrorf(format, args...)
	}
}
