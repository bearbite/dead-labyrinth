package tunnel

import "math/rand"

type Tunnel interface {
	GetFront() string
	GetRight() string
	GetLeft() string
}

type tunnel struct {
	front string
	right string
	left  string
	y     int
	x     int
}

func New(x int, y int) Tunnel {

	t := &tunnel{}

	t.front, t.right, t.left = t.randTunnel()

	return t

}

func (t *tunnel) GetFront() string {
	return t.front
}

func (t *tunnel) GetRight() string {
	return t.right
}

func (t *tunnel) GetLeft() string {
	return t.left
}

func (t *tunnel) randTunnel() (front, right, left string) {
	randPG := []string{"room", "tunnel"}
	front = randPG[rand.Intn(2)]
	right = randPG[rand.Intn(2)]
	left = randPG[rand.Intn(2)]
	return front, right, left
}
