package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	endpoint := Point{
		s.start.x + int(s.a),
		s.start.y + int(s.a),
	}

	return endpoint
}

func (s *Square) Area() uint {
	return s.a * s.a
}

func (s *Square) Perimeter() uint {
	return 2*s.a + 2*s.a
}
