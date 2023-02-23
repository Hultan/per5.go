package per5

type Vector2D struct {
	X, Y float64
}

func (p *Per5) CreateVector2D(x, y float64) Vector2D {
	return Vector2D{x, y}
}

func (v *Vector2D) Mult(f float64) Vector2D {
	return Vector2D{v.X * f, v.Y * f}
}
