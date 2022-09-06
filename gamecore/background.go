package snake

import "math/rand"

//背景画布
type Coord struct {
	x, y int
}

//战场背景
type Background struct {
	Beans      *Beans
	Snake      *Snake
	hasBeans   func(*Background, Coord) bool
	Height     int
	Width      int
	pointsChan chan (int)
}

func newBackground(snake *Snake, h, w int) *Background {
	b := &Background{
		Snake:    snake,
		Height:   h,
		Width:    w,
		hasBeans: hasBeans,
	}

	return b
}

//随机生产豆子
func (b *Background) makeBeans() {
	var x, y int
	for {
		x = rand.Intn(b.Width)
		y = rand.Intn(b.Height)

		if !b.Snake.isHits(Coord{x: x, y: y}) {
			break
		}
	}
	b.Beans = newBeans(x, y)
}

func (b *Background) moveSnake() error {
	if err := b.Snake.move(); err != nil {
		return err
	}
	if b.isLeftBackground() {
		return b.Snake.over()
	}

	return nil
}

func (b *Background) isLeftBackground() bool {
	h := b.Snake.head()
	return h.x > b.Width-1 || h.y > b.Height-1 || h.x < 0 || h.y < 0
}

func hasBeans(background *Background, coord Coord) bool {
	return coord.x == background.Beans.x && coord.y == background.Beans.y
}
