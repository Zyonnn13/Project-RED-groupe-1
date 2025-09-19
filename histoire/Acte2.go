package histoire

import (
	"Project-RED-groupe-1/combat"
	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/player"
	"fmt"
)

func Acte2_Alliances(p *player.Player, inv *inventaire.Inventory) {
	fmt.Println("\n=== ACTE 2 : ALLIANCES ===")
	fmt.Println("Après la mort de Jackie et le braquage raté, tu erres sans repères...")
	fmt.Println("Deux choix s’offrent à toi :")
	fmt.Println("1. Chercher de l’aide auprès de Rogue, légende des Fixers.")
	fmt.Println("2. Te rapprocher de Panam et du clan des Aldecaldos.")

	var choix string
	fmt.Print("\nFais ton choix : ")
	fmt.Scanln(&choix)

	switch choix {
	case "1":
		fmt.Println("\nTu pars à la recherche de Rogue dans l’Afterlife.")
		fmt.Println("Elle t’écoute, sceptique, mais accepte de t’aider…")
		combat.LancerCombat(p, combat.Agentcorpo, inv)
		fmt.Println("Avec Rogue, tu ouvres la voie vers une alliance puissante.")
	case "2":
		fmt.Println("\nTu retrouves Panam à l’extérieur de Night City.")
		fmt.Println("Avec son aide, tu gagnes la confiance des Aldecaldos.")
		combat.LancerCombat(p, combat.Agentcorpo, inv)
		fmt.Println("Le clan t’accepte comme l’un des leurs.")
	default:
		fmt.Println("Choix invalide. Tu restes perdu, errant sans but...")
	}
}
