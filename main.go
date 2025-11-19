// gocat implements `Lolcat` in Golang
// https://github.com/busyloop/lolcat
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	// Is Stdin a TTY? (terminal)
	if isPiped() {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println(`Usage: echo "text to color" | gocat`)
		os.Exit(1)
	}

	// Read from Pipe
	text := readFromPipe()

	// Print the text
	printRainbow(string(text), 0.1)

}

// Returns true if the stdin is being piped, otherwise false
func isPiped() bool {
	info, _ := os.Stdin.Stat()
	return info.Mode()&os.ModeCharDevice != 0 // Bitwise AND
}

// Reads text from pipe (stdin) and returns it as a string
func readFromPipe() string {
	// Read from Stdin
	var text []rune
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		text = append(text, input)
	}

	return string(text)
}

/*
rgb generates a smooth, continuously changing RGB color based on an integer input `i`.

f: frequency (keep between 0 and 1)
smaller `f` means more spreaded gradient.
higher `f` means opposite.
*/
func rgb(i int, f float64) (int, int, int) {
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)

}

// printRainbow prints `output` text with rainbow colors in terminal.
// text: the text to print
// frequency: controls gradient spread(ness)
func printRainbow(text string, frequency float64) {
	// Colorize the text
	for i := 0; i < len(text); i++ {
		r, g, b := rgb(i, frequency) // Generate RGB based on i
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, text[i])
	}
	// fmt.Println()
}
