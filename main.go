package main

import (
	"bytes"
	"flag"
	"log"
	"os"

	"github.com/JonPulfer/sudusolver/puzzle"
)

var jsonPuzzle string

func init() {
	flag.StringVar(&jsonPuzzle, "json", "",
		"JSON encoded matrix of the puzzle with unassigned locations set to 0")
	flag.Parse()
}

func main() {

	p := puzzle.NewPuzzle()
	if len(jsonPuzzle) > 0 {
		var buf bytes.Buffer
		buf.WriteString(jsonPuzzle)
		p = puzzle.ParseJSONEncodedPuzzle(&buf)
	}

	if !puzzle.Solve(p) {
		log.Println("failed to solve the puzzle")
		return
	}

	if err := p.Serialize(os.Stdout); err != nil {
		log.Printf("error serializing puzzle: %s\n", err.Error())
	}
}
