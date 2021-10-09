package rectcount

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestCount(t *testing.T) {

	tc(t, 0, []Point{
		{0, 1},
		{1, 0},
		{1, 2},
	})

	tc(t, 0, []Point{
		{0, 1},
		{1, 0},
		{1, 2},
		{5, 20},
	})

	tc(t, 1, []Point{ // square, 0°
		{0, 1},
		{1, 0},
		{1, 1},
		{0, 0},
	})

	tc(t, 1, []Point{ // square, 0°, with a duplicated point
		{0, 1},
		{1, 0},
		{1, 1},
		{0, 0},
		{1, 0}, // duplicated point !
	})

	tc(t, 3, []Point{ // imbricated, squares and rectangles,  0°
		{0, 1},
		{1, 0},
		{1, 1},
		{0, 0},
		{0, 2},
		{1, 2},
	})

	tc(t, 1, []Point{ // 45°
		{0, 1},
		{1, 0},
		{1, 2},
		{2, 1},
	})

	tc(t, 1, []Point{ // 30°
		{0, 0},
		{2, 1},
		{1, 3},
		{-1, 2},
	})

	tc(t, 2, []Point{ // 45° + 30°
		{0, 1},
		{1, 0},
		{1, 2},
		{2, 1},

		{0, 0},
		{1, 3},
		{-1, 2},
	})

	tc(t, 5, []Point{ // all mixed
		{0, 1},
		{1, 1},
		{1, 0},
		{0, 0},
		{1, 2},
		{0, 3},
		{3, 0},
		{-1, 0},
		{3, 3},
		{3, 1},
		{2, 5},
	})

	tc(t, 5, []Point{ // all mixed and duplications
		{0, 1},
		{1, 1},
		{1, 0},
		{0, 0},
		{1, 2},
		{0, 3},
		{3, 0},
		{-1, 0},
		{3, 3},
		{3, 1},
		{2, 5},
		// dups...
		{0, 0},
		{1, 2},
		{0, 3},
		{3, 0},
		{-1, 0},
	})

}

// -------------- benchmarks -----------------------------
func BenchmarkCount50(b *testing.B) {
	// 5000 random points
	data := make([]Point, 50)
	for i := range data {
		data[i] = Point{rand.Int(), rand.Int()}
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Count(data)
	}
}
func BenchmarkCount500(b *testing.B) {
	// 5000 random points
	data := make([]Point, 500)
	for i := range data {
		data[i] = Point{rand.Int(), rand.Int()}
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Count(data)
	}
}
func BenchmarkCount5000(b *testing.B) {
	// 5000 random points
	data := make([]Point, 5000)
	for i := range data {
		data[i] = Point{rand.Int(), rand.Int()}
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Count(data)
	}
}

// -------------- utility functions ----------------------
func tc(t *testing.T, want int, input []Point) {

	c := Count(input)

	if c != want {
		draw(input)
		t.Fatalf("unexpected count number, got %d, want %d", c, want)
	}
}

// draw from -20 to +20 included,
// for debugging purposes
func draw(points []Point) {

	if len(points) == 0 {
		return
	}

	var plan [41][41]bool // x+20 , y + 20
	for _, p := range points {
		plan[p.X+20][p.Y+20] = true
	}
	for x := 0; x < 41; x++ {
		for y := 0; y < 41; y++ {
			if plan[x][y] {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
