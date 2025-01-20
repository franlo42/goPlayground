package main

import "fmt"

func main() {
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	blue := "\033[34m"
	magenta := "\033[35m"
	cyan := "\033[36m"
	reset := "\033[0m"

	fmt.Printf("%sEste texto es rojo%s\n", red, reset)
	fmt.Printf("%sEste texto es verde%s\n", green, reset)
	fmt.Printf("%sEste texto es amarillo%s\n", yellow, reset)
	fmt.Printf("%sEste texto es azul%s\n", blue, reset)
	fmt.Printf("%sEste texto es magenta%s\n", magenta, reset)
	fmt.Printf("%sEste texto es cian%s\n", cyan, reset)
}
