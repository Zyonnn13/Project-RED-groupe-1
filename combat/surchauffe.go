package combat

import "fmt"

// ApplySurchauffe active l'effet pour 3 tours à 10 dégâts/tour.
func ApplySurchauffe(e *Ennemis) {
	// Si une surchauffe est déjà active, on la rafraîchit.
	e.SurchauffeTours = 3
	e.SurchauffeDegats = 10
	fmt.Printf("\n🔥 %s est en surchauffe ! Il subira %d dégâts pendant %d tours.\n", e.Name, e.SurchauffeDegats, e.SurchauffeTours)
}

// TickSurchauffe applique les dégâts en début de tour si l'effet est actif.
func TickSurchauffe(e *Ennemis) {
	if e.SurchauffeTours <= 0 || e.SurchauffeDegats <= 0 {
		return
	}

	e.HP -= e.SurchauffeDegats
	if e.HP < 0 {
		e.HP = 0
	}
	fmt.Printf("🔥 Dégâts de surchauffe: %s perd %d PV. PV restants: %d/%d (Tours restants: %d)\n",
		e.Name, e.SurchauffeDegats, e.HP, e.MaxHP, e.SurchauffeTours-1)

	e.SurchauffeTours--
}
