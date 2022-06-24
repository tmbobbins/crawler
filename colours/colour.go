package colours

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func WrapRed(body any) string {
	return Wrap(Red, body)
}

func WrapYellow(body any) string {
	return Wrap(Yellow, body)
}

func WrapPurple(body any) string {
	return Wrap(Purple, body)
}

func WrapCyan(body any) string {
	return Wrap(Cyan, body)
}

func Wrap(colour string, body any) string {
	return fmt.Sprintf("%s%s%s", colour, body, Reset)
}

func PrintRed(body any) {
	fmt.Println(WrapRed(body))
}

func PrintYellow(body any) {
	fmt.Println(WrapYellow(body))
}

func PrintCyan(body any) {
	fmt.Println(WrapCyan(body))
}

func PrintPurple(body any) {
	fmt.Println(WrapPurple(body))
}
