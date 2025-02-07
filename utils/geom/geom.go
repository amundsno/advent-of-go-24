package geom

type Vec2D struct {
	X, Y int
}

func NewVec2D(x, y int) Vec2D {
	return Vec2D{X: x, Y: y}
}

func (v Vec2D) Add(w Vec2D) Vec2D {
	return Vec2D{v.X + w.X, v.Y + w.Y}
}

func (v Vec2D) TurnRight() Vec2D {
	return Vec2D{-v.Y, v.X}
}

func (v Vec2D) TurnLeft() Vec2D {
	return Vec2D{v.Y, -v.X}
}
