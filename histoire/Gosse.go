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
	fmt.Println("1 - Accepter de négocier")
	fmt.Println("2 - Proposer de rembourser toi-même")
	fmt.Println("3 - Tendre un piège à Kirk")
	fmt.Println("4 - [Gosse des rues] Appeler un gang allié")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("\nTu rencontres Kirk dans une arrière-salle. Les Valentinos te fixent froidement.")
		if persuasionReussie() {
			fmt.Println("✅ Ta persuasion réussit ! Pas de combat, Kirk accepte de réduire la dette.")
		} else {
			fmt.Println("❌ Échec de persuasion... Les Valentinos passent à l’attaque !")
			ennemi := combat.Valentinos
			fmt.Printf("⚔️ %s surgissent ! (HP: %d | ATK: %d)\n", ennemi.Name, ennemi.HP, ennemi.Attack)
			combat.LancerCombat(p, ennemi, &inventaire.Inventory{})
		}

	case "2":
		fmt.Println("\nTu poses ton fric sur la table. Kirk ricane, prend l’argent et repart satisfait.")
		fmt.Println("Tu perds beaucoup d’argent, mais Pepe devient ton allié loyal.")
		// Ici pas de combat

	case "3":
		fmt.Println("\nTu tends un piège à Kirk dans une ruelle sombre...")
		ennemi1 := combat.Kirk
		ennemi2 := combat.Valentinos
		ennemi3 := combat.Valentinos
		fmt.Println("⚔️ Kirk et ses deux sbires sortent leurs armes !")
		// Trois combats à la suite simulant un affrontement de groupe
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

// petite fonction utilitaire pour simuler une persuasion
func persuasionReussie() bool {
	// exemple simple : 50% de chance de réussite
	return randInt(0, 1) == 1
}

// helper pour un int aléatoire
func randInt(min, max int) int {
	return min + int(int64(max-min+1)*int64(os.Getpid())%int64(max-min+1))
}
