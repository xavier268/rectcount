package rectcount

// Point
type Point struct {
	X int
	Y int
}

// Pair of Points
// Pairs of points represent a diagonal of the rectangle.
// We have a rectangle if and only if :
// 1) both diagonals are different and non degenerated
// 2) diagonal MIDDLE POINT are the same
// 3) diagonal LENGTH are the same

type pair struct {
	a Point
	b Point
}

// signature of a pair, that represent a DIAGONAL of the potential rectangle.
// Two distinct pairs with the same signature form a rectangle.
type signature struct {
	mx, my int // twice the middle coordinate (to avoid floating points !)
	d2     int // square of the diagonal distance
}

func (p pair) sign() (s signature) {
	s.mx, s.my = p.a.X+p.b.X, p.a.Y+p.b.Y
	s.d2 = (p.a.X-p.b.X)*(p.a.X-p.b.X) + (p.a.Y-p.b.Y)*(p.a.Y-p.b.Y)
	return s
}

// Count the number of rectangles that can be formed
// Compute time is measured at O(n^2.1), worst theoretical case O(n^3)
func Count(input []Point) int {

	// dedup
	points := dedup(input)
	answer := 0

	// pairs maps the signature to the pair of points that share it
	pairs := make(map[signature]([]pair), len(points)*len(points))

	for i := 0; i < len(points); i++ { // O(n)
		p1 := points[i]
		for j := i + 1; j < len(points); j++ { // O(n^2)
			p2 := points[j]
			// if p1 != p2 {
			// Unnecessay test, we did the dedup of points earlier and i,j loop indexes do not overlap.
			p := pair{p1, p2}
			s := p.sign()
			answer += len(pairs[s])        // add a rectangle with each eligible prior pair with that signature
			pairs[s] = append(pairs[s], p) // O(n^2) * n = O(n^3), according to runtime.mapaccess1 worst-case behaviour, but in practise, much better !
			// }
		}
	}
	return answer
}

// depup returns a set of Points in nLog(n) time
func dedup(in []Point) (out []Point) {

	set := make(map[Point]bool, len(in))
	for _, p := range in {
		set[p] = true
	}
	out = make([]Point, 0, len(in))
	for k := range set {
		out = append(out, k)
	}
	return out
}
