package circle

func (c Circle) Circumference() float64 {
	const phi = 3.14
	return 2 * phi * c.Radius
}
