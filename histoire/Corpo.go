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

// Histoire Corpo avec nom du joueur
func CorpoHistoire(p *player.Player) {
	fmt.Println("\n=== Histoire Corpo ===")
	fmt.Printf(
		"Les néons de Night City se reflètent sur les vitres teintées de la tour Arasaka. "+
			"Dans ton bureau, un message crypté s’affiche : « %s, on a un problème. »\n",
		p.Name,
	)
	fmt.Println("Ton supérieur, Jenkins, t’explique qu’il faut discréditer un rival politique. Tu sais que ce genre de mission ne se termine jamais proprement.")
}

func StartCorpo(p *player.Player, inv *inventaire.Inventory) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nQue fais-tu ?")
	fmt.Println("1 - Accepter sans poser de questions")
	fmt.Println("2 - [Corpo] Utiliser ton réseau pour saboter ton rival avant même la mission")
	fmt.Println("3 - Refuser")

	fmt.Print("\nChoix : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("Jenkins t’accorde sa confiance, mais tu te retrouves sur le terrain face à des opposants.")
		ennemi := combat.Ncpd
		fmt.Printf("\n%s apparaît !\n", ennemi.Name)
		fmt.Printf("HP : %d   |    ATK : %d\n", ennemi.HP, ennemi.Attack)
		fmt.Println("\nUn combat commence ! Prépare-toi à riposter.")

		combat.LancerCombat(p, ennemi, inv)

		// Récompense après combat
		fmt.Println("\n🎁 Tu fouilles la zone et trouves un Maxdoc et 100 eddies !")
		inv.AddItem(inventaire.Item{
			Nom:         "Maxdoc M.K 1",
			Description: "Soin rapide de 50 PV",
			Type:        "soin",
			Effet:       50,
			Consommable: true,
		})
		p.Eddies.Add(100)

	case "2":
		fmt.Println("Grâce à tes contacts, ton rival est neutralisé avant même d’avoir commencé. Tu évites un affrontement direct.")

	case "3":
		fmt.Println("Refuser un ordre direct de Jenkins est dangereux... Tu sens que cela ne va pas bien se terminer.")
		ennemi := combat.Agentcorpo
		fmt.Printf("\n%s apparaît !\n", ennemi.Name)
		fmt.Printf("HP : %d   |    ATK : %d\n", ennemi.HP, ennemi.Attack)
		fmt.Println("\nUn combat commence ! Prépare-toi à riposter.")

		combat.LancerCombat(p, ennemi, inv)

		// Récompense après combat
		fmt.Println("\n🎁 Tu fouilles la zone et trouves un Maxdoc et 100 eddies !")
		inv.AddItem(inventaire.Item{
			Nom:         "Maxdoc M.K 1",
			Description: "Soin rapide de 50 PV",
			Type:        "soin",
			Effet:       50,
			Consommable: true,
		})
		p.Eddies.Add(100)

	default:
		fmt.Println("Choix invalide.")
	}
}
