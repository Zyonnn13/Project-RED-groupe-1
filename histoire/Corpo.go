package histoire

import (
	"Project-RED-groupe-1/combat"
	"Project-RED-groupe-1/player"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CorpoHistoire() {
	fmt.Println("\n=== Histoire Corpo ===")
	fmt.Println("Les néons de Night City se reflètent sur les vitres teintées de la tour Arasaka. Dans ton bureau, un message crypté s’affiche : « V, on a un problème. »")
	fmt.Println("Ton supérieur, Jenkins, t’explique qu’il faut discréditer un rival politique. Tu sais que ce genre de mission ne se termine jamais proprement.")
}

func StartCorpo(p *player.Designplayer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Que fais-tu ?")
	fmt.Println("1- Accepter sans poser de questions ")
	fmt.Println("2- [Corpo] Utiliser ton réseau pour saboter ton rival avant même la mission ")
	fmt.Println("3- Refuser ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("Jenkins t’accorde sa confiance, mais tu te retrouves sur le terrain face à des opposants.")
		ennemi := combat.Ncpd
		fmt.Printf("⚔️ %s apparaît ! (HP: %d | ATK: %d)\n", ennemi.Name, ennemi.Hp, ennemi.Attaque)

		// 👉 ICI tu déclenches le combat
		combat.LancerCombat(p, ennemi)

	case "2":
		fmt.Println("Grâce à tes contacts, ton rival est neutralisé avant même d’avoir commencé. Tu évites un affrontement direct.")
		// Ici pas de combat → tu continues l’histoire

	case "3":
		fmt.Println("Refuser un ordre direct de Jenkins est dangereux... Tu sens que cela ne va pas bien se terminer.")
		ennemi := combat.Agentcorpo
		fmt.Printf("⚔️ %s est envoyé pour te punir ! (HP: %d | ATK: %d)\n", ennemi.Name, ennemi.Hp, ennemi.Attaque)

		// 👉 Combat déclenché ici aussi
		combat.LancerCombat(p, ennemi)
	}
}
