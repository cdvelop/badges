package shell

import "fmt"

const (
	ColorRed    = "\033[0;31m"
	ColorGreen  = "\033[0;32m"
	ColorYellow = "\033[0;33m"
	ColorCyan   = "\033[0;36m"
	ColorNone   = "\033[0m"
)

func Success(msg string) {
	fmt.Printf("%s%s%s\n", ColorGreen, msg, ColorNone)
}

func Warning(msg string) {
	fmt.Printf("%s%s%s\n", ColorYellow, msg, ColorNone)
}

func Error(msg string) {
	fmt.Printf("%sError: %s%s\n", ColorRed, msg, ColorNone)
}

func Info(msg string) {
	fmt.Printf("%s%s%s\n", ColorCyan, msg, ColorNone)
}
