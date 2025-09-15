package player

import "fmt"

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

// Fonction pour subir des dégâts
func (p *Pvjoueur) SubirDegats(degats int) {
	p.Pv -= degats
	if p.Pv < 0 {
		p.Pv = 0
	}
}

// Fonction pour recevoir des soins
func (p *Pvjoueur) Soigner(soin int) {
	p.Pv += soin
	if p.Pv > 100 {
		p.Pv = 100
	}
}

// Vérifie si le joueur est en vie
func (p Pvjoueur) EstEnVie() bool {
	return p.Pv > 0
}
