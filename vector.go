package rl

type V [4]int

func (v *V) Set(vv V) {
	*v = vv
}

func (v *V) X() int {
	return v[0]
}

func (v *V) SetX(i int) {
	v[0] = i
}

func (v *V) Y() int {
	return v[1]
}

func (v *V) SetY(i int) {
	v[1] = i
}

func (v *V) Z() int {
	return v[2]
}

func (v *V) SetZ(i int) {
	v[2] = i
}

func (v *V) Add(o V) V {
	var sum V
	for i := range sum {
		sum[i] = v[i] + o[i]
	}
	return sum
}
