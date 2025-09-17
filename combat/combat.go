package combat

import (
	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/player"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func LancerCombat(joueur *player.Player, ennemi Ennemis, inv *inventaire.Inventory) {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n Un combat commence contre %s !\n", ennemi.Name)

	for joueur.HP > 0 && ennemi.HP > 0 {
		afficherStats(joueur, ennemi)

		// Tour du joueur
		fmt.Println("\nQue veux-tu faire ?")
		fmt.Println("1 - Attaquer avec votre armes")
		fmt.Println("2 - Inventaire")
		fmt.Println("3 - Fuir")
		fmt.Print("Choix : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			// Attaque
			degat := joueur.Arme.Attaque()
			ennemi.HP -= degat
			if ennemi.HP < 0 {
				ennemi.HP = 0
			}
			fmt.Printf("Tu attaques %s avec ton %s et infliges %d dégâts !\n", ennemi.Name, joueur.Arme.Nom, degat)

		case "2":
			// Inventaire
			inv.ShowInventory()
			fmt.Print("Quel objet veux-tu utiliser ? ")
			itemChoice, _ := reader.ReadString('\n')
			itemChoice = strings.TrimSpace(itemChoice)
			success := inv.UseItem(itemChoice, joueur)
			if success {
				continue
			}

		case "3":
			fmt.Println("Tu t’enfuis du combat !")
			return

		default:
			fmt.Println("Choix invalide, ton tour est perdu !")
		}

		if ennemi.HP <= 0 {
			fmt.Printf("🎉 %s est vaincu !\n", ennemi.Name)
			break
		}

		fmt.Printf("\n--- Tour de %s ---\n", ennemi.Name)
		degat := calcDamage(ennemi.Attack)
		joueur.HP -= degat
		if joueur.HP < 0 {
			joueur.HP = 0
		}
		fmt.Printf("%s attaque et inflige %d dégâts !\n", ennemi.Name, degat)

		if joueur.HP <= 0 {
			fmt.Println("💀 Tu as été vaincu...")
			break
		}
	}

	fmt.Println("Fin du combat.")
}

func afficherStats(joueur *player.Player, ennemi Ennemis) {
	fmt.Printf("\n--- Statut ---\n")
	fmt.Printf("❤️ %s : %d/%d HP\n", joueur.Name, joueur.HP, joueur.MaxHP)
	fmt.Printf("💀 %s : %d/%d HP\n", ennemi.Name, ennemi.HP, ennemi.MaxHP)
}

func calcDamage(base int) int {
	variance := base / 5
	if variance < 1 {
		variance = 1
	}
	return base + rand.Intn(variance*2+1) - variance
}
