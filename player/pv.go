package player

import (
	"fmt"
)

// États possibles du jeu
type EtatJeu int

const (
	EtatExploration EtatJeu = iota
	EtatShop
	EtatCombat
)

// Structure des PV du joueur
type Pvjoueur struct {
	Pv int
}

// Affiche les PV seulement dans le shop ou en combat
func (p Pvjoueur) AfficherPV(etat EtatJeu) {
	if etat == EtatShop || etat == EtatCombat {
		fmt.Println("Points de vie :", p.Pv)
	}
}
