package player

import (
	"fmt"
	"strings"
)

type EtatJeu int

const (
	EtatExploration EtatJeu = iota
	EtatShop
	EtatCombat
)

type Joueur struct {
	Nom      string
	Sante    int
	SanteMax int
}

func (j Joueur) AfficherBarreDeSante() {
	longueur := 30
	remplissage := int(float64(j.Sante) / float64(j.SanteMax) * float64(longueur))

	barre := "[" + strings.Repeat("█", remplissage) + strings.Repeat("░", longueur-remplissage) + "]"
	fmt.Printf("\n%s - Santé: %d/%d %s\n", j.Nom, j.Sante, j.SanteMax, barre)
}
