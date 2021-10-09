package rectcount

// Point
type Point struct {
	x int
	y int
}

// Pair of Points
// Pairs of points represent a diagonal of the rectangle.
// We have a rectangle if and only if :
// 1) both diagonals are different and non degenerated
// 2) diagonal MIDDLE POINT are the same
// 3) diagonal LENGTH are the same

type Pair struct {
	a Point
	b Point
}

// Signature of a pair, that represent a DIAGONAL of the potential rectangle.
// Two distinct pairs with the same signature form a rectangle.
type Signature struct {
	mx, my int // twice the middle coordinate (to avoid floating points !)
	d2     int // square of the diagonal distance
}

func (p Pair) sign() (s Signature) {
	s.mx, s.my = p.a.x+p.b.x, p.a.y+p.b.y
	s.d2 = (p.a.x-p.b.x)*(p.a.x-p.b.x) + (p.a.y-p.b.y)*(p.a.y-p.b.y)
	return s
}

// Pair maps the signature to the pair of points that share it
type Pairs map[Signature]([]Pair)

// count the number of rectangles that can be formed
// Compute time is O(n^2)log(n)
func count(input []Point) int {
	// dedup
	points := dedup(input)
	answer := 0
	pairs := make(Pairs, len(points)*len(points))
	for i := 0; i < len(points); i++ { // O(n)
		p1 := points[i]
		for j := i; j < len(points); j++ { // O(n^2)
			p2 := points[j]
			if p1 != p2 {
				// non degenerated pairs only ...
				p := Pair{p1, p2}
				s := p.sign()
				answer += len(pairs[s])        // add a rectangle with each eligible prior pair with that signature
				pairs[s] = append(pairs[s], p) // worst case, O(n^2)log(n)
			}
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
