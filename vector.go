package rl

type V [4]int

func (v V) X() int {
	return v[0]
}

func (v V) Y() int {
	return v[1]
}

func (v V) Z() int {
	return v[2]
}

func addVector(t, v V) V {
	var o = v
	for i := range o {
		o[i] = t[i] + v[i]
	}
	return o
}
