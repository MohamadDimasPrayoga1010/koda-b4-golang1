package circle

func (c Circle) Area() float64 {
	const phi = 3.14
	return phi * c.Radius * c.Radius
}
