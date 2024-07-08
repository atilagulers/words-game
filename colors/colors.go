package colors

import "fmt"

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

func PrintColored(text, color string) {
	fmt.Printf("%s%s%s", color, text, Reset)
}

func PrintRed(text string) {
	PrintColored(text, Red)
}

func PrintGreen(text string) {
	PrintColored(text, Green)
}

func PrintYellow(text string) {
	PrintColored(text, Yellow)
}

func PrintBlue(text string) {
	PrintColored(text, Blue)
}

func PrintMagenta(text string) {
	PrintColored(text, Magenta)
}

func PrintCyan(text string) {
	PrintColored(text, Cyan)
}

func PrintWhite(text string) {
	PrintColored(text, White)
}
