package combat

type Ennemis struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

var Opposant = Ennemis{"Opposant", 100, 100, 50}
var Valentinos = Ennemis{"Valentinos", 100, 100, 50}
var Ncpd = Ennemis{"Ncpd", 100, 100, 50}
