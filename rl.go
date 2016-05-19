package rl

// Game holds all the state for a game.
type Game struct {
	Player Player
	turn   int
	Map    Map
	done   bool
}

// NewGame creates a new game using the controller supplied.
func NewGame(c Controller) *Game {
	return &Game{
		Player: Player{
			Controller: c,
		},
	}
}

func (g *Game) Step() bool {
	a := g.Player.Controller.Get()
	if a == nil {
		return g.done
	}
	a.do(g)
	return g.done
}

// Player holds all the state for the player.
type Player struct {
	Position   V
	Controller Controller
}

// Controller is an interface for getting actions.
type Controller interface {
	Get() Action
}
