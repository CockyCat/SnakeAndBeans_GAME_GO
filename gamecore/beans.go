package snake

type Beans struct {
	points, x, y int
}

func newBeans(x, y int) *Beans {
	return &Beans{
		points: 10,
		x:      x,
		y:      y,
	}
}
