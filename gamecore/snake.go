package snake

import "errors"

const (
	RIGHT Direction = 1 + iota

	LEFT

	UP

	DOWN
)

type Direction int

type Snake struct {
	Body      []Coord   //坐标
	Direction Direction //方向
	Length    int       //长度
}

//初始化Snake对象
func newSnake(d Direction, b []Coord) *Snake {
	return &Snake{
		Length:    len(b),
		Body:      b,
		Direction: d,
	}
}

func (this *Snake) changeDirection(d Direction) {
	//定义反方向
	opposites := map[Direction]Direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}
	//
	if o := opposites[d]; o != 0 && o != this.Direction {
		this.Direction = d
	}
}

func (this *Snake) head() Coord {
	return this.Body[len(this.Body)-1]
}

func (this *Snake) over() error {
	return errors.New("Snake Game Over")
}

func (this *Snake) move() error {

	//身体的头坐标
	h := this.head()

	c := Coord{x: h.x, y: h.y}

	switch this.Direction {
	case RIGHT:
		c.y++
	case LEFT:
		c.y--
	case UP:
		c.x--
	case DOWN:
		c.x++
	}
	//碰撞检测
	if this.isHits(c) {
		return this.over()
	}
	if this.Length > len(this.Body) {
		this.Body = append(this.Body, c)
	} else {
		this.Body = append(this.Body[1:], c)
	}

	return nil
}

//是否碰撞
func (this *Snake) isHits(c Coord) bool {
	for _, b := range this.Body {
		if b.x == c.x && b.y == c.y {
			return true
		}
	}
	return false
}
