package histoire

import (
	"Project-RED-groupe-1/combat"
	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/player"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Histoire Nomade
func NomadeHistoire() {
	fmt.Println("\n=== Histoire Nomade ===")
	fmt.Println("Spawn : miroir dans un garage de voiture V, un ancien membre d’un clan nomade des Badlands, décide de quitter sa vie itinérante pour tenter sa chance à Night City.")
	fmt.Println("L’histoire débute dans la ville poussiéreuse de Yucca, où V tente de réparer son véhicule pour rejoindre la métropole.")
}

// Début de l’aventure Nomade interactive
func StartNomade(p *player.Player, inv *inventaire.Inventory) {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)

	fmt.Println("\nTu es dans ton garage à Yucca. Ton véhicule est à moitié démonté.")
	fmt.Println("1 - Continuer à réparer le véhicule")
	fmt.Println("2 - Aller voir Jackie Welles pour discuter du plan de contrebande")
	fmt.Println("3 - Explorer la ville avant de partir")

	fmt.Print("\nChoix : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("\nTu passes des heures à réparer ton véhicule. Le moteur ronronne enfin.")
		p.Eddies.Add(20)
		fmt.Println("🎁 Tu gagnes 20 eddies pour avoir vendu quelques pièces de ton ancien véhicule.")
	case "2":
		fmt.Println("\nTu retrouves Jackie Welles. Il te confie la mission : faire passer un colis à travers la frontière.")
		fmt.Println("Ensemble, vous soudoyez les gardes-frontières.")

		if chance < 70 { // 70% de réussite
			fmt.Println("\nVotre plan fonctionne parfaitement ! Les gardes-frontières sont soudoyés sans problème.")
			p.Eddies.Add(50) // petite récompense
		} else {
			fmt.Println("\nLe plan échoue partiellement ! Les gardes remarquent quelque chose et vous devez combattre.")
			combat.LancerCombat(p, combat.Valentinos, inv)
		}
		// Éventuel mini-combat ou test de chance
		combat.LancerCombat(p, combat.Valentinos, inv)
		fmt.Println("\nVous réussissez à semer les gardes d’Arasaka !")
	case "3":
		fmt.Println("\nEn explorant Yucca, tu découvres un iguane vivant dans un garage abandonné.")
		inv.AddItem(inventaire.Item{
			Nom:         "Iguane vivant",
			Description: "Un mystérieux iguane trouvé dans un garage",
			Type:        "objet",
			Effet:       0,
			Consommable: false,
		})
		p.Eddies.Add(10)
		fmt.Println("🎁 Tu gagnes 10 eddies en vendant quelques pièces de récupération.")

	default:
		fmt.Println("Choix invalide. Tu restes à réfléchir dans ton garage...")
	}

	StartNomadeSuite(p, inv)
}

// Suite de l’aventure Nomade
func StartNomadeSuite(p *player.Player, inv *inventaire.Inventory) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nAprès avoir échappé aux gardes, vous avancez vers Night City.")
	fmt.Println("Vous devez maintenant décider de la prochaine étape :")
	fmt.Println("1 - Rejoindre la métropole directement")
	fmt.Println("2 - Chercher un endroit sûr pour cacher le colis")
	fmt.Println("3 - Faire une pause et améliorer ton véhicule")

	fmt.Print("\nChoix : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("\nVous prenez la route vers Night City, le vent dans les cheveux et le moteur qui rugit.")

		// Nouveau test de chance pour une embuscade
		rand.Seed(time.Now().UnixNano())
		chance := rand.Intn(100)
		if chance < 60 {
			fmt.Println("\nVous parvenez à passer sans problème. Aucun bandit sur votre route.")
			p.Eddies.Add(30)
		} else {
			fmt.Println("\nUn petit groupe de bandits tente de vous bloquer ! Préparez-vous au combat.")
			combat.LancerCombat(p, combat.Agentarasaka, inv)
			fmt.Println("\nVous continuez votre route sain et sauf.")
		}

	case "2":
		fmt.Println("\nVous trouvez un garage abandonné pour cacher le colis. C’est temporaire, mais sûr.")
		inv.AddItem(inventaire.Item{
			Nom:         "Garage sécurisé",
			Description: "Permet de cacher temporairement le colis",
			Type:        "objet",
			Effet:       0,
			Consommable: false,
		})

	case "3":
		fmt.Println("\nTu travailles sur ton véhicule, améliorant moteur et suspension.")
		p.Eddies.Add(30)
		fmt.Println("🎁 Tu gagnes 30 eddies en vendant de vieilles pièces.")

	default:
		fmt.Println("Choix invalide. Vous perdez du temps et devez improviser votre route...")
	}

	// Étape finale vers la métropole
	fmt.Println("\nAprès avoir pris les décisions précédentes, vous approchez enfin de Night City.")
	fmt.Println("Mais Arasaka est sur vos traces. Que faites-vous ?")
	fmt.Println("1 - Contacter Jackie pour un plan rapide")
	fmt.Println("2 - Prendre un raccourci risqué mais plus rapide")
	fmt.Println("3 - Ralentir et éviter les zones surveillées")

	fmt.Print("\nChoix : ")
	finalChoice, _ := reader.ReadString('\n')
	finalChoice = strings.TrimSpace(finalChoice)

	switch finalChoice {
	case "1":
		fmt.Println("\nJackie organise une diversion parfaite. Vous entrez dans Night City sans encombre.")
		p.Eddies.Add(50)
	case "2":
		fmt.Println("\nLe raccourci est dangereux ! Vous êtes attaqués par une patrouille d’Arasaka.")
		combat.LancerCombat(p, combat.Agentarasaka, inv)
		fmt.Println("\nVous survivez et atteignez Night City avec le colis.")
	case "3":
		fmt.Println("\nVous évitez les zones surveillées et arrivez plus tard, mais indemne. Prudence payée !")
		p.Eddies.Add(20)
	}

	fmt.Println("\nFélicitations ! Vous avez complété la première mission Nomade à Night City.")
	fmt.Println("L’histoire peut continuer avec d’autres missions ou combats.")
}
