package armes

import "math/rand"

type Arme struct {
	Nom        string
	Degats     int
	Precision  float64
	CritChance float64
}

func (a Arme) Attaque() int {
	if rand.Float64() <= a.Precision {
		d := a.Degats
		if rand.Float64() <= a.CritChance {
			d *= 2
		}
		return d
	}
	return 0
}
