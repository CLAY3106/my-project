package main // this package builds an executable and will run the main() function.

import (
	"fmt" // fmt is formatted I/O
	"math"
	"strings" // imports strings package for string manipulation

	"syreclabs.com/go/faker" // imports the faker package for generating fake data
)

// rgb function generates RGB color values based on an integer input.
func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func main() {
	var phrases []string // phrases is a slice of strings to hold generated phrases

	// Loop to generate phrases using the faker library

	for i := 1; i < 3; i++ {
		phrases = append(phrases, faker.Hacker().Phrases()...) // append generated phrases to the slice
	}

	output := strings.Join(phrases[:], "; ") // join the phrases into a single string separated by "; "

	// Define RGB color values for gold

	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)                                              // get RGB values based on the character index
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j]) // print the output with color
	}
}
