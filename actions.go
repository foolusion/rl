package rl

type Action interface {
	do(*Game)
}

type move V

func (m move) do(g *Game) {
	pos := g.Player.Position.Add(V(m))
	if !g.Map.HasAttribute("walkable", pos) {
		return
	}
	g.Player.Position.Set(pos)
}

var (
	// Move actions for controllers
	MoveLeft  = move{-1, 0, 0, 0}
	MoveRight = move{1, 0, 0, 0}
	MoveUp    = move{0, -1, 0, 0}
	MoveDown  = move{0, 1, 0, 0}
)

type quit struct{}

func (q quit) do(g *Game) {
	g.done = true
}

var Quit = quit{}
