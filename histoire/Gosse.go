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

// Histoire Gosse des rues
func GosseHistoire() {
	fmt.Println("\n=== Histoire Gosse des rues ===")
	fmt.Println("Spawn : El Coyote Cojo, un bar du district de Heywood, connu pour ses gangs.")
	fmt.Println("Ton personnage a le nez cassé, que vous replacez avant de commencer.")
	fmt.Println("Le patron du bar vous demande de l’aider car il doit de l’argent à Kirk.")
	fmt.Println("Kirk vous propose de voler une voiture en échange de le laisser tranquille.")
}

// Début de l’aventure interactive
func StartGosse(p *player.Player, inv *inventaire.Inventory) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nQue fais-tu ?")
	fmt.Println("1 - Accepter de voler la voiture Rayfield Aerondight")
	fmt.Println("2 - Refuser et rester au bar")
	fmt.Println("3 - Explorer le quartier pour récupérer des infos")

	fmt.Print("\nChoix : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("\nVous vous rendez à la voiture pour le vol.")
		fmt.Println("Le gadget ne fonctionne pas et la voiture déclenche l’alarme !")

		// Combat contre la police
		fmt.Println("La police arrive sur le champ ! Combat engagé !")
		combat.LancerCombat(p, combat.Ncpd, inv)

		// Apparition de Jackie
		fmt.Println("\nUn mystérieux allié apparaît : Jackie !")
		fmt.Println("Grâce à lui, le combat se réengage et vous parvenez à battre la police.")
		combat.LancerCombat(p, combat.Ncpd, inv)

		// Récompenses et progression
		fmt.Println("\nVous et Jackie êtes désormais partenaires. Vous menez plusieurs contrats pendant les 6 prochains mois,")
		fmt.Println("et votre réputation de rue atteint le niveau 1.")

		p.Eddies.Add(100)
		inv.AddItem(inventaire.Item{
			Nom:         "Rayfield Aerondight",
			Description: "Voiture volée lors de la première mission",
			Type:        "vehicule",
			Effet:       0,
			Consommable: false,
		})

	case "2":
		fmt.Println("\nVous décidez de rester au bar. Rien ne se passe, mais vous restez en sécurité.")
	case "3":
		fmt.Println("\nEn explorant Heywood, vous récupérez quelques informations sur les gangs locaux.")
		p.Eddies.Add(20)
		fmt.Println("🎁 Tu gagnes 20 eddies en récupérant ces infos.")

	default:
		fmt.Println("Choix invalide. Le temps passe et l’opportunité s’éloigne...")
	}

	StartGosseSuite(p, inv)
}

func StartGosseSuite(p *player.Player, inv *inventaire.Inventory) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nAprès la première mission, tu dois décider de la prochaine étape :")
	fmt.Println("1 - Accepter un contrat de gang pour voler une cargaison")
	fmt.Println("2 - Rechercher un allié pour renforcer ton équipe")
	fmt.Println("3 - Prendre du repos et améliorer ton équipement")

	fmt.Print("\nChoix : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("\nVous partez voler une cargaison sous la protection de Jackie.")
		combat.LancerCombat(p, combat.Agentarasaka, inv)
		p.Eddies.Add(50)
	case "2":
		fmt.Println("\nVous recrutez un nouvel allié dans le quartier pour renforcer votre équipe.")
		inv.AddItem(inventaire.Item{
			Nom:         "Nouvel allié",
			Description: "Partenaire de mission pour les prochains contrats",
			Type:        "allie",
			Effet:       0,
			Consommable: false,
		})
	case "3":
		fmt.Println("\nVous améliorez votre équipement et récupérez des eddies supplémentaires.")
		p.Eddies.Add(30)
	default:
		fmt.Println("Choix invalide. Tu perds du temps à Heywood...")
	}

	fmt.Println("\nFélicitations ! Tu as complété la première série de missions de Gosse des rues.")
	fmt.Println("Ton aventure dans le district de Heywood ne fait que commencer...")
}
