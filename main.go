package main

import (
	"log"
	"os"

	"github.com/make-42/gokanjivg/colorize"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	modes := []string{"pastel", "spectrum", "double-spectrum"}
	characters := []string{"漢", "字", "鬱"}
	for _, character := range characters {
		for _, mode := range modes {
			svgData, err := colorize.Colorize(character, mode, 0.95, 0.75)
			check(err)
			err = os.WriteFile("./output/"+character+"-"+mode+".svg", []byte(svgData), 0644)
			check(err)
		}
	}
}
