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
	Nom:        "Pistolet1",
	Degats:     10,
	Precision:  0.40,
	CritChance: 0.0,
	Niveau:     1,
}

var Pistolet2 = Arme{
	Nom:        "Pistolet2",
	Degats:     10,
	Precision:  0.40,
	CritChance: 0.0,
	Niveau:     2,
}
var Pistolet3 = Arme{
	Nom:        "Pistolet3",
	Degats:     10,
	Precision:  0.40,
	CritChance: 0.0,
	Niveau:     3,
}

var Pistolet4 = Arme{
	Nom:        "Pistolet4",
	Degats:     10,
	Precision:  0.40,
	CritChance: 0.0,
	Niveau:     4,
}

var Pistolet5 = Arme{
	Nom:        "Pistolet5",
	Degats:     10,
	Precision:  0.40,
	CritChance: 0.0,
	Niveau:     5,
}

var Fusilspompe1 = Arme{
	Nom:        "Fusils a pompe1",
	Degats:     50,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     1,
}
var Fusilspompe2 = Arme{
	Nom:        "Fusils a pompe2",
	Degats:     50,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     2,
}
var Fusilspompe3 = Arme{
	Nom:        "Fusils a pompe3",
	Degats:     50,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     3,
}
var Fusilspompe4 = Arme{
	Nom:        "Fusils a pompe4",
	Degats:     50,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     4,
}
var Fusilspompe5 = Arme{
	Nom:        "Fusils a pompe5",
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
	Nom:        "Mitraillette2",
	Degats:     35,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     2,
}
var Mitraillette3 = Arme{
	Nom:        "Mitraillette3",
	Degats:     35,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     3,
}
var Mitraillette4 = Arme{
	Nom:        "Mitraillette4",
	Degats:     35,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     4,
}
var Mitraillette5 = Arme{
	Nom:        "Mitraillette5",
	Degats:     35,
	Precision:  0.50,
	CritChance: 0.0,
	Niveau:     5,
}

var Fusil1 = Arme{
	Nom:        "Fusils d'Assault1",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     1,
}

var Fusil2 = Arme{
	Nom:        "Fusils d'Assault2",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     2,
}
var Fusil3 = Arme{
	Nom:        "Fusils d'Assault3",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     3,
}
var Fusil4 = Arme{
	Nom:        "Fusils d'Assault4",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     4,
}
var Fusil5 = Arme{
	Nom:        "Fusils d'Assault5",
	Degats:     55,
	Precision:  0.60,
	CritChance: 0.0,
	Niveau:     5,
}

var FusilPrecision1 = Arme{
	Nom:        "Fusils de Precision1",
	Degats:     60,
	Precision:  0.80,
	CritChance: 0.0,
	Niveau:     1,
}
var FusilPrecision2 = Arme{
	Nom:        "Fusils de Precision2",
	Degats:     60,
	Precision:  0.80,
	CritChance: 0.0,
	Niveau:     2,
}

var FusilPrecision3 = Arme{
	Nom:        "Fusils de Precision3",
	Degats:     60,
	Precision:  0.80,
	CritChance: 0.0,
	Niveau:     3,
}

var FusilPrecision4 = Arme{
	Nom:        "Fusils de Precision4",
	Degats:     60,
	Precision:  0.80,
	CritChance: 0.0,
	Niveau:     4,
}

var FusilPrecision5 = Arme{
	Nom:        "Fusils de Precision5",
	Degats:     60,
	Precision:  0.80,
	CritChance: 0.0,
	Niveau:     5,
}
