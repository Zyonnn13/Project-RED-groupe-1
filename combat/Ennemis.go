package combat

import (
	"fmt"
	"math/rand"
	"strings"
)

type Ennemis struct {
	Name      string
	MaxHP     int
	HP        int
	Attack    int
	Degats    int
	Precision float64
}

var Opposant = Ennemis{"Opposant", 100, 100, 50, 50, 0.50}
var Valentinos = Ennemis{"Valentinos", 100, 100, 50, 50, 0.50}
var Ncpd = Ennemis{"Ncpd", 100, 100, 50, 50, 0.50}
var Agentcorpo = Ennemis{"Agent Corpo", 100, 100, 50, 50, 0.50}
var Agentrivaux = Ennemis{"Agent Rival", 100, 100, 50, 50, 0.50}
var Kirk = Ennemis{"Kirk", 120, 120, 60, 50, 0.50}
var Agentsecu = Ennemis{"Agent de Sécurité", 100, 100, 50, 50, 0.50}
var Agentarasaka = Ennemis{"Agent Arasaka", 100, 100, 50, 50, 0.50}
var Drone = Ennemis{"Drone", 75, 75, 40, 50, 0.50}
var Wraiths = Ennemis{"Wraiths", 100, 100, 50, 50, 0.50}
var Tygerclaws = Ennemis{"Tyger Claws", 100, 100, 50, 50, 0.50}
var Chefclaws = Ennemis{"Chef des Tyger Claws", 150, 150, 100, 50, 0.50}
var Dronelourd = Ennemis{"Drone Lours", 90, 90, 50, 50, 0.50}
var Tourelleauto = Ennemis{"Tourelle Automatique", 10, 10, 200, 50, 0.50}
var Sniper = Ennemis{"Sniper", 50, 50, 100, 50, 0.50}
var Netrunner = Ennemis{"Netrunner", 50, 50, 100, 50, 0.50}
var Adam = Ennemis{"Adam Smasher", 200, 200, 200, 50, 0.50}

func (a Ennemis) Attaque() int {
	if rand.Float64() <= a.Precision {
		return a.Degats
	}
	return 0
}

func (e *Ennemis) AfficherBarreDeVie(style string) {
	barLength := 20
	filled := (e.HP * barLength) / e.MaxHP

	var bar string
	switch style {
	case "compact":
		bar = strings.Repeat("█", filled) + strings.Repeat("░", barLength-filled)
		fmt.Printf("HP de %s: [%s] %d/%d\n", e.Name, bar, e.HP, e.MaxHP)
	case "classic":
		fmt.Printf("\nSanté de %s : %d/%d\n", e.Name, e.HP, e.MaxHP)
		fmt.Print("[")
		for i := 0; i < barLength; i++ {
			if i < filled {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("]")
	default:
		fmt.Println("Style inconnu.")
	}

}
