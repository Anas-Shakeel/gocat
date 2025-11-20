// gocat implements `Lolcat` in Golang
// https://github.com/busyloop/lolcat
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

// Named command-line arguments
// var spread = flag.Float64("spread", 0.1, "controls how wide the gradient is spread")
var spread = flag.Float64("spread", 0.1, "gradient spread (lower values make smoother gradients)")
var textFile = flag.String("file", "", "Read input from a file")

func main() {
	// Custom usage text
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), getCustomUsage())
		flag.PrintDefaults()
	}

	// Parse command line arguments
	flag.Parse()
	positionalArgs := flag.Args()

	switch {
	case isPiped():
		// Read from stdin
		text := readFromPipe()
		printRainbow(text, *spread)

	case *textFile != "":
		// Read input from file
		text, err := readFromFile(*textFile)
		if err != nil {
			log.Fatal(err)
		}
		printRainbow(string(*text), *spread)

	case len(positionalArgs) > 0:
		// Read text from cmd args
		text := strings.Join(positionalArgs, " ")
		printRainbow(text, *spread)

	default:
		// Default case, Quit!
		flag.Usage()
		os.Exit(1)
	}
}

// Returns the custom usage text to be used in help message.
func getCustomUsage() string {
	return `Usage: gocat [OPTIONS] [text...]

Options:`
}

// Returns true if the stdin is being piped, otherwise false
func isPiped() bool {
	info, _ := os.Stdin.Stat()
	return info.Mode()&os.ModeCharDevice == 0 // Bitwise AND
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

// Reads file and returns it's contents as []byte slice (*[]byte, a pointer actually!)
func readFromFile(filename string) (*[]byte, error) {
	// Open file to read
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &content, nil
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

// printRainbow prints `text` as rainbow-style text in terminal.
// This method calls `makeRainbow()` under the hood.
// text: the text to color
// spread: controls gradient spread(ness)
func printRainbow(text string, spread float64) {
	fmt.Println(makeRainbow(text, spread))
}

// makeRainbow colors `text` as rainbow-style text and returns it.
// text: the text to color
// spread: controls gradient spread(ness)
func makeRainbow(text string, spread float64) string {
	var output strings.Builder

	// Colorize the text
	for i := 0; i < len(text); i++ {
		r, g, b := rgb(i, spread) // Generate RGB based on i
		output.WriteString(fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, text[i]))
	}

	return output.String()
}
