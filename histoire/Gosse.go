package histoire

import (
	"Project-RED-groupe-1/combat"
	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/player"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GosseHistoire() {
	fmt.Println("\n=== Histoire Gosse des rues ===")
	fmt.Println("Heywood, ton quartier. Les ruelles sentent la sueur, le chrome et les deals foireux.")
	fmt.Println("Ton ami Pepe t’appelle à l’aide : il doit de l’argent à Kirk, un petit caïd local lié aux Valentinos.")
	fmt.Println("Tu décides d’aller voir Kirk dans une arrière-salle d’un bar miteux...")
}

func StartGosse(p *player.Player) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nComment veux-tu gérer Kirk et sa dette ?")
	fmt.Println("1 \t- Accepter de négocier")
	fmt.Println("2 \t- Proposer de rembourser toi-même")
	fmt.Println("3 \t- Tendre un piège à Kirk")
	fmt.Println("4 \t- [Gosse des rues] Appeler un gang allié")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("\nTu rencontres Kirk dans une arrière-salle. Les Valentinos te fixent froidement.")
		if persuasionReussie() {
			fmt.Println("Ta persuasion réussit ! Pas de combat, Kirk accepte de réduire la dette.")
		} else {
			fmt.Println(" Échec de persuasion... Les Valentinos passent à l’attaque !")
			ennemi := combat.Valentinos
			fmt.Printf("\n %s apparaît !\n", ennemi.Name)
			fmt.Printf(" HP : %d   |    ATK : %d\n", ennemi.HP, ennemi.Attack)
			fmt.Println("\n Un combat commence ! Prépare-toi à riposter.")

			combat.LancerCombat(p, ennemi, &inventaire.Inventory{})
		}

	case "2":
		fmt.Println("\nTu poses ton fric sur la table. Kirk ricane, prend l’argent et repart satisfait.")
		fmt.Println("Tu perds beaucoup d’argent, mais Pepe devient ton allié loyal.")
		if p.Eddies.Spend(100) {
			fmt.Println(" Tu perds 200 eddies, mais Pepe devient ton allié loyal.")
		} else {
			fmt.Println(" Tu n’as pas assez d’eddies pour convaincre Pepe.")
			fmt.Printf("💰 Solde actuel : %d eddies\n", p.Eddies.GetBalance())
		}

	case "3":
		fmt.Println("\nTu tends un piège à Kirk dans une ruelle sombre...")
		ennemi1 := combat.Kirk
		ennemi2 := combat.Valentinos
		ennemi3 := combat.Valentinos
		fmt.Println(" Kirk et ses deux sbires sortent leurs armes !")
		combat.LancerCombat(p, ennemi1, &inventaire.Inventory{})
		if p.HP > 0 {
			combat.LancerCombat(p, ennemi2, &inventaire.Inventory{})
		}
		if p.HP > 0 {
			combat.LancerCombat(p, ennemi3, &inventaire.Inventory{})
		}

	case "4":
		fmt.Println("\nTu appelles un gang rival (les 6th Street). Ils interviennent et forcent Kirk à annuler la dette.")
		fmt.Println("Pas de combat, mais tu perds ton joker : ils ne t’aideront plus jamais.")
		fmt.Println("Les Valentinos, eux, te gardent désormais à l’œil...")

	default:
		fmt.Println("Choix invalide. Kirk t’ignore et la situation empire pour Pepe.")
	}
}

func persuasionReussie() bool {
	return randInt(0, 1) == 1
}

func randInt(min, max int) int {
	return min + int(int64(max-min+1)*int64(os.Getpid())%int64(max-min+1))
}
