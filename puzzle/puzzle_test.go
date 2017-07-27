package puzzle

import (
	"fmt"
	"testing"
)

func TestPuzzleDisplay(t *testing.T) {
	p := NewPuzzle()
	fmt.Println(p)
}

func TestFirstLocation(t *testing.T) {
	p := NewPuzzle()
	l := p.NewLocation()
	if l == nil {
		t.Logf("did not select correct first location: %v\n", l)
		t.FailNow()
	}
	if l.Row != 0 && l.Col != 0 {
		t.Fail()
	}
}

func TestSecondLocation(t *testing.T) {
	p := NewPuzzle()

	p.Grid[0][0] = 1
	l := p.NewLocation()

	if l == nil {
		t.Logf("failed to get second location: %v\n", l)
		t.FailNow()
	}

	if l.Row != 0 && l.Col != 1 {
		t.Fail()
	}
}

func TestTenthLocation(t *testing.T) {
	p := NewPuzzle()
	for i := 1; i < 10; i++ {
		l := p.NewLocation()
		p.Assign(l, i)
	}

	l := p.NewLocation()
	if l == nil {
		t.Logf("failed to get correct location: %v\n", l)
		t.FailNow()
	}

	if l.Row != 1 && l.Col != 0 {
		t.Fail()
	}
}

func TestAlreadyInRow(t *testing.T) {
	p := NewPuzzle()
	p.Assign(&Location{Row: 0, Col: 7}, 3)

	l := p.NewLocation()
	if !p.alreadyInRow(l, 3) {
		t.Fail()
	}
}

func TestAlreadyInColumn(t *testing.T) {
	p := NewPuzzle()
	p.Assign(&Location{Row: 8, Col: 0}, 3)

	l := p.NewLocation()
	if !p.alreadyInColumn(l, 3) {
		t.Fail()
	}
}

func TestAlreadyInBlock(t *testing.T) {
	p := NewPuzzle()
	p.Assign(&Location{Row: 2, Col: 2}, 3)

	l := p.NewLocation()
	if !p.alreadyInBlock(l, 3) {
		t.Fail()
	}
}

func TestSafeToAssign(t *testing.T) {
	p := NewPuzzle()
	p.Assign(&Location{Row: 2, Col: 2}, 3)

	l := p.NewLocation()
	if p.SafeToAssign(l, 3) {
		t.Fail()
	}
}

func TestSolveEmpty(t *testing.T) {
	p := NewPuzzle()
	if !Solve(p) {
		t.Fail()
	}
	fmt.Println(p)

}

func TestSolveOneAssigned(t *testing.T) {
	p := NewPuzzle()
	p.Assign(&Location{Row: 2, Col: 2}, 3)

	if !Solve(p) {
		t.Fail()
	}
	fmt.Println(p)
}

func BenchmarkAlreadyInColumn(b *testing.B) {
	p := NewPuzzle()
	p.Assign(&Location{Row: 6, Col: 0}, 3)

	l := p.NewLocation()

	for i := 0; i < b.N; i++ {
		if !p.alreadyInColumn(l, 3) {
		}
	}
}

func BenchmarkAlreadyInRow(b *testing.B) {
	p := NewPuzzle()
	p.Assign(&Location{Row: 0, Col: 6}, 3)

	l := p.NewLocation()

	for i := 0; i < b.N; i++ {
		if !p.alreadyInRow(l, 3) {
		}
	}
}

func BenchmarkAlreadyInBlock(b *testing.B) {
	p := NewPuzzle()
	p.Assign(&Location{Row: 2, Col: 2}, 3)

	l := p.NewLocation()

	for i := 0; i < b.N; i++ {
		if !p.alreadyInBlock(l, 3) {
		}
	}
}

func BenchmarkSolvingEmptyPuzzle(b *testing.B) {

	p := NewPuzzle()
	for i := 0; i < b.N; i++ {
		if Solve(p) {
		}
	}
}

func BenchmarkSolvingPuzzleWithOneInput(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p := NewPuzzle()
		p.Assign(&Location{Row: 2, Col: 2}, 3)
		if Solve(p) {
		}
	}
}

func BenchmarkSolvingPuzzleWithTwoInputs(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p := NewPuzzle()
		p.Assign(&Location{Row: 2, Col: 2}, 3)
		p.Assign(&Location{Row: 8, Col: 0}, 3)
		if Solve(p) {
		}
	}
}

func BenchmarkSolvingPuzzleWithThreeInputs(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p := NewPuzzle()
		p.Assign(&Location{Row: 2, Col: 2}, 3)
		p.Assign(&Location{Row: 8, Col: 0}, 3)
		p.Assign(&Location{Row: 7, Col: 8}, 3)
		if Solve(p) {
		}
	}
}
