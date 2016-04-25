package rl

type Game struct {
	Player  Player
	turn    int
	gameMap gameMap
}

type Player struct {
	name     string
	Position Vector
}

type gameMap struct {
	size Vector
}

func (g *gameMap) hasAttribute(attr string, pos Vector) bool {
	return true
}

type Vector [4]int

func (v Vector) X() int {
	return v[0]
}

func (v *Vector) setX(i int) {
	v[0] = i
}

func (v Vector) Y() int {
	return v[1]
}

func (v *Vector) setY(i int) {
	v[1] = i
}

func (v Vector) Z() int {
	return v[2]
}

func (v *Vector) setZ(i int) {
	v[2] = i
}

func (v Vector) add(o Vector) Vector {
	var sum Vector
	for i := range sum {
		sum[i] = v[i] + o[i]
	}
	return sum
}

func Move(g *Game, direction Vector) {
	newPos := g.Player.Position.add(direction)
	if !g.gameMap.hasAttribute("walkable", newPos) {
		return
	}
	g.Player.Position = newPos
}
