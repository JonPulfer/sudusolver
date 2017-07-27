package puzzle

import (
	"encoding/json"
	"io"
	"log"
	"strconv"
	"strings"
)

// Block holds the top left location of the block in the matrix.
type Block struct {
	Row, Col int
}

// Location describes a position in the puzzle matrix.
type Location struct {
	Row, Col int
}

// Block the location resides in.
func (l *Location) Block() *Block {
	br := 3 * (l.Row / 3)
	bc := 3 * (l.Col / 3)
	return &Block{Row: br, Col: bc}
}

// Puzzle provides the interactions with a suduko puzzle grid.
type Puzzle struct {
	Grid [][]int
}

// NewPuzzle initialises a puzzle grid.
func NewPuzzle() *Puzzle {

	grid := make([][]int, 9)
	for r := 0; r < 9; r++ {
		row := make([]int, 9)
		for c := 0; c < 9; c++ {
			row[c] = 0
		}
		grid[r] = row
	}

	return &Puzzle{
		Grid: grid,
	}
}

// ParseJSONEncodedPuzzle and return the Puzzle. If the parsing fails, nil is
// returned.
func ParseJSONEncodedPuzzle(r io.Reader) *Puzzle {
	p := NewPuzzle()
	if err := json.NewDecoder(r).Decode(&p.Grid); err != nil {
		log.Println(err.Error())
		return nil
	}
	return p
}

// Serialize the Puzzle.
func (p *Puzzle) Serialize(w io.Writer) error {
	return json.NewEncoder(w).Encode(&p.Grid)
}

// String representation of the puzzle grid.
func (p *Puzzle) String() string {
	output := make([]string, 0, 1)
	for _, r := range p.Grid {
		line := func() string {
			chars := make([]string, 0, 1)
			for x := 0; x < 9; x++ {
				chars = append(chars, strconv.Itoa(r[x]))
			}
			return strings.Join(chars, " ")
		}
		output = append(output, line())
	}
	return strings.Join(output, "\n")
}

// NewLocation returns the next unallocated location. If there are no locations
// left unallocated, this returns nil.
func (p *Puzzle) NewLocation() *Location {
	for rowNum, row := range p.Grid {
		for colNum, col := range row {
			if col < 1 {
				return &Location{Row: rowNum, Col: colNum}
			}
		}
	}
	return nil
}

// Assign a value directly to a location without validation.
func (p *Puzzle) Assign(l *Location, v int) {
	if l.Row < 9 && l.Col < 9 {
		p.Grid[l.Row][l.Col] = v
	}
}

// Reset the value at the location.
func (p *Puzzle) Reset(l *Location) {
	if l.Row < 9 && l.Col < 9 {
		p.Grid[l.Row][l.Col] = 0
	}
}

// alreadyInRow checks to see if the value is already in the row
func (p *Puzzle) alreadyInRow(l *Location, v int) bool {
	for _, col := range p.Grid[l.Row] {
		if col == v {
			return true
		}
	}
	return false
}

// alreadyInColumn checks whether the value is already assigned in the selected
// column.
func (p *Puzzle) alreadyInColumn(l *Location, v int) bool {
	for _, row := range p.Grid {
		if row[l.Col] == v {
			return true
		}
	}
	return false
}

// alreadyInBlock checks whether the value already exists in the block that
// encompasses the selected location.
func (p *Puzzle) alreadyInBlock(l *Location, v int) bool {
	b := l.Block()

	for r := b.Row; r < b.Row+3; r++ {
		for c := b.Col; c < b.Col+3; c++ {
			if p.Grid[r][c] == v {
				return true
			}
		}
	}
	return false
}

// SafeToAssign indicates whether it is safe to assign the value to the
// provided location.
func (p *Puzzle) SafeToAssign(l *Location, v int) bool {
	return !p.alreadyInRow(l, v) &&
		!p.alreadyInColumn(l, v) &&
		!p.alreadyInBlock(l, v)
}

// Solve the puzzle recursively.
func Solve(p *Puzzle) bool {

	l := p.NewLocation()
	if l == nil {
		return true
	}

	for i := 1; i <= 9; i++ {
		if p.SafeToAssign(l, i) {
			p.Assign(l, i)
			if Solve(p) {
				return true
			}
			p.Reset(l)
		}
	}

	return false
}
