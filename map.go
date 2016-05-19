package rl

// Map holds the state of the map.
type Map struct {
	size V
}

// HasAttribute returns wether the tile at the specificed pos has attribute
// attr.
func (g *Map) HasAttribute(attr string, pos V) bool {
	return true
}
