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

func SuiteCorpoChoix1(p *player.Player, inv *inventaire.Inventory) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nTu avances dans les couloirs du complexe Arasaka, encore sous tension après le combat.")
	fmt.Println("Tout est silencieux, mais tu sais que ce n’est pas fini…")
	fmt.Println("Soudain, tu entends des bruits métalliques derrière une porte entrouverte.")
	fmt.Println("1 - Ouvrir la porte et affronter ce qui se cache à l’intérieur")
	fmt.Println("2 - Passer discrètement et ignorer le bruit")
	fmt.Println("3 - T’enfermer dans un bureau proche pour observer")

	fmt.Print("\nChoix : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("\nUne escouade de gardes arrive ! Prépare-toi au combat.")
		ennemi := combat.Agentarasaka
		combat.LancerCombat(p, ennemi, inv)

		// Récompense spécifique au choix 1
		fmt.Println("\n🎁 Après le combat, tu trouves un Cyberdeck et 100 eddies.")
		inv.AddItem(inventaire.Item{
			Nom:         "Cyberdeck Mk I",
			Description: "Permet d’effectuer des hacks rapides",
			Type:        "hack",
			Effet:       0,
			Consommable: false,
		})
		p.Eddies.Add(100)

		// Suite narration spécifique
		fmt.Println("\nEn explorant la zone, tu trouves un terminal sécurisé contenant des infos sensibles.")
		fmt.Println("1 - Télécharger les données pour Jenkins")
		fmt.Println("2 - Les effacer pour éviter que d'autres ne s’en servent")

		fmt.Print("\nChoix : ")
		subChoice, _ := reader.ReadString('\n')
		subChoice = strings.TrimSpace(subChoice)

		if subChoice == "1" {
			fmt.Println("\nTu télécharges les données et gagnes la confiance de Jenkins. Ta réputation monte !")
			p.Eddies.Add(50)
		} else if subChoice == "2" {
			fmt.Println("\nTu effaces les données. Cela évite un scandale, mais Jenkins sera mécontent...")
			p.Eddies.Add(10)
		} else {
			fmt.Println("Choix invalide. Tu restes sur place, le temps passe...")
		}

	case "2":
		fmt.Println("\nTu passes discrètement, évitant tout combat. Tu gagnes du temps et quelques eddies pour la prudence.")
		p.Eddies.Add(20)

	case "3":
		fmt.Println("\nTu observes depuis le bureau. Un garde passe et laisse tomber un paquet.")
		inv.AddItem(inventaire.Item{
			Nom:         "Pack de soins mineur",
			Description: "Restaure 20 PV",
			Type:        "soin",
			Effet:       20,
			Consommable: true,
		})

	default:
		fmt.Println("Choix invalide. Tu restes immobile, perdu dans tes pensées...")
	}
}
