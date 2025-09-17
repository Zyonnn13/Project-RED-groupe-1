package histoire

import (
	"Project-RED-groupe-1/combat"
	"Project-RED-groupe-1/player"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Intro de l'histoire Nomade
func NomadeHistoire() {
	fmt.Println("\n=== Histoire Nomade ===")
	fmt.Println("Le désert s’étend à perte de vue. Les moteurs de ton clan vrombissent derrière toi.")
	fmt.Println("Tu approches de Night City, mais les Wraiths rodent, prêts à vous attaquer.")
}

// Suite de l'histoire avec les choix
func StartNomade(p *player.Designplayer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nQue fais-tu ?")
	fmt.Println("1 - Protéger ton clan et attaquer les Wraiths en premier")
	fmt.Println("2 - [Nomade] Utiliser ton véhicule pour une manœuvre d’intimidation")
	fmt.Println("3 - Tenter de négocier")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("\n💥 Tu charges les Wraiths, arme en main !")
		ennemi := combat.Wraiths
		fmt.Printf("⚔️ %s apparaît ! (HP: %d | ATK: %d)\n", ennemi.Name, ennemi.Hp, ennemi.Attaque)
		combat.LancerCombat(p, ennemi)

	case "2":
		fmt.Println("\n🚗 Grâce à ta conduite, tu fais fuir une partie du gang… mais certains restent.")
		ennemi := combat.Tygerclaws
		fmt.Printf("⚔️ %s surgit ! (HP: %d | ATK: %d)\n", ennemi.Name, ennemi.Hp, ennemi.Attaque)
		combat.LancerCombat(p, ennemi)

	case "3":
		fmt.Println("\n🤝 Tu lèves les mains et proposes un deal…")
		// Tu peux imaginer un mini test de persuasion plus tard
		fmt.Println("Mais les Wraiths refusent et ouvrent le feu !")
		ennemi := combat.Wraiths
		fmt.Printf("⚔️ %s attaque ! (HP: %d | ATK: %d)\n", ennemi.Name, ennemi.Hp, ennemi.Attaque)
		combat.LancerCombat(p, ennemi)

	default:
		fmt.Println("Choix invalide, les Wraiths attaquent par surprise !")
		ennemi := combat.Wraiths
		combat.LancerCombat(p, ennemi)
	}
}
