package parser

type TriangleType int

// Triangle types.
const (
	Obtuse TriangleType = iota
	Isoceles
	Scalene
)

type Triangle struct {
	Type   TriangleType
	Angles []float64
	Sides  int
}
