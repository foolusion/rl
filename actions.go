package rl

type action interface {
	do(*game)
}

type move struct {
	id           int
	displacement V
}

func (m move) do(g *game) {
	g.positions[m.id] = walk(g.gMap, g.positions[m.id], m.displacement)
}

func walk(m Map, position, displacement V) V {
	if !m.HasAttribute("walkable", addVector(position, displacement)) {
		return position
	}
	return addVector(position, displacement)
}

var (
	// Move actions for controllers
	moveLeft  = move{displacement: V{-1, 0, 0, 0}}
	moveRight = move{displacement: V{1, 0, 0, 0}}
	moveUp    = move{displacement: V{0, -1, 0, 0}}
	moveDown  = move{displacement: V{0, 1, 0, 0}}
)

type quit struct{}

func (q quit) do(g *game) {
	g.done = true
}
