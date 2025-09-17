package armes

import "math/rand"

type Arme struct {
	Nom        string
	Degats     int
	Precision  float64
	CritChance float64
	Niveau     int
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

var Pistolet1 = Arme{
	Nom:        "Pistolet",
	Degats:     45,
	Precision:  0.70,
	CritChance: 0.20,
	Niveau:     1,
}

var Pistolet2 = Arme{
	Nom:        "Pistolet",
	Degats:     10,
	Precision:  0.40,
	CritChance: 0.0,
	Niveau:     2,
}
var Pistolet3 = Arme{
	Nom:        "Pistolet",
	Degats:     10,
	Precision:  0.40,
	CritChance: 0.0,
	Niveau:     3,
}

var Pistolet4 = Arme{
	Nom:        "Pistolet",
	Degats:     10,
	Precision:  0.40,
	CritChance: 0.0,
	Niveau:     4,
}

var Pistolet5 = Arme{
	Nom:        "Pistolet",
	Degats:     10,
	Precision:  0.40,
	CritChance: 0.0,
	Niveau:     5,
}

var Fusilspompe1 = Arme{
	Nom:        "Fusils a pompe",
	Degats:     50,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     1,
}
var Fusilspompe2 = Arme{
	Nom:        "Fusils a pompe",
	Degats:     50,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     2,
}
var Fusilspompe3 = Arme{
	Nom:        "Fusils a pompe",
	Degats:     50,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     3,
}
var Fusilspompe4 = Arme{
	Nom:        "Fusils a pompe",
	Degats:     50,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     4,
}
var Fusilspompe5 = Arme{
	Nom:        "Fusils a pompe",
	Degats:     50,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     5,
}

var Mitraillette = Arme{
	Nom:        "Mitraillette",
	Degats:     35,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     1,
}
var Mitraillette2 = Arme{
	Nom:        "Mitraillette",
	Degats:     35,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     2,
}
var Mitraillette3 = Arme{
	Nom:        "Mitraillette",
	Degats:     35,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     3,
}
var Mitraillette4 = Arme{
	Nom:        "Mitraillette",
	Degats:     35,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     4,
}
var Mitraillette5 = Arme{
	Nom:        "Mitraillette",
	Degats:     35,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     5,
}

var Fusil1 = Arme{
	Nom:        "Fusils d'Assault",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     1,
}

var Fusil2 = Arme{
	Nom:        "Fusils d'Assault",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     2,
}
var Fusil3 = Arme{
	Nom:        "Fusils d'Assault",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     3,
}
var Fusil4 = Arme{
	Nom:        "Fusil d'Assault",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     4,
}
var Fusil5 = Arme{
	Nom:        "Fusil d'Assault",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     5,
}
var Katana1 = Arme{
	Nom:        "Katana",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     1,
}
var Katana2 = Arme{
	Nom:        "Katana",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     2,
}
var Katana3 = Arme{
	Nom:        "Katana",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     3,
}
var Katana4 = Arme{
	Nom:        "Katana",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     4,
}
var Katana5 = Arme{
	Nom:        "Katana",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     5,
}
var Couteau1 = Arme{
	Nom:        "Couteau",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     1,
}
var Couteau2 = Arme{
	Nom:        "Couteau",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     2,
}
var Couteau3 = Arme{
	Nom:        "Couteau",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     3,
}
var Couteau4 = Arme{
	Nom:        "Couteau",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     4,
}
var Couteau5 = Arme{
	Nom:        "Couteau",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     5,
}
