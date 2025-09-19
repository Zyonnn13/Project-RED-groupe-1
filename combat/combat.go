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
		TickSurchauffe(&ennemi)
		if ennemi.HP <= 0 {
			fmt.Printf(" %s est vaincu !\n", ennemi.Name)
			break
		}

		afficherStats(joueur, ennemi)

		actionDone := false
		for !actionDone {
			fmt.Println("\nQue veux-tu faire ?")
			fmt.Println("1 - Attaquer avec votre arme")
			fmt.Println("2 - Inventaire")
			fmt.Println("3 - Utiliser Surchauffe")
			fmt.Print("\nChoix : ")

			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)

			switch choice {
			case "1":
				degat, crit := joueur.Arme.Attaque()
				ennemi.HP -= degat
				if ennemi.HP < 0 {
					ennemi.HP = 0
				}
				if degat > 0 {
					if crit {
						fmt.Printf("\n COUP CRITIQUE ! Tu attaques %s avec ton %s et infliges %d dégâts !\n", ennemi.Name, joueur.Arme.Nom, degat)
					} else {
						fmt.Printf("\nTu attaques %s avec ton %s et infliges %d dégâts.\n", ennemi.Name, joueur.Arme.Nom, degat)
					}
				} else {
					fmt.Printf("\n Tu attaques %s avec ton %s... mais tu rates !\n", ennemi.Name, joueur.Arme.Nom)
				}
				joueur.AfficherBarreDeVie("compact")
				ennemi.AfficherBarreDeVie("compact")
				actionDone = true

			case "2":
				inv.ShowInventory()
				fmt.Print("Quel objet veux-tu utiliser ? ")
				itemChoice, _ := reader.ReadString('\n')
				itemChoice = strings.TrimSpace(itemChoice)
				if itemChoice == "" {
					continue
				}
				if inv.UseItem(itemChoice, joueur) {
					actionDone = true
				}
			case "3":
				ApplySurchauffe(&ennemi)
				actionDone = true

			default:
				fmt.Println("Choix invalide, ton tour est perdu !")
				actionDone = true
			}

			if ennemi.HP <= 0 {
				fmt.Printf(" %s est vaincu !\n", ennemi.Name)
				break
			}

			fmt.Printf("\n--- Tour de %s ---\n", ennemi.Name)
			degat := ennemi.Attaque()
			joueur.HP -= degat
			if joueur.HP < 0 {
				joueur.HP = 0
			}
			fmt.Printf("%s attaque et inflige %d dégâts !\n", ennemi.Name, degat)

			if joueur.HP <= 0 {
				fmt.Println("💀 Tu as été vaincu...")
				os.Exit(0)
			}
		}
	}
	fmt.Println("Fin du combat.")
}

func afficherStats(joueur *player.Player, ennemi Ennemis) {
	fmt.Println("\n--- Statut ---")
	fmt.Printf("%s : %d/%d HP\n", joueur.Name, joueur.HP, joueur.MaxHP)
	joueur.AfficherBarreDeVie("compact")
	fmt.Printf(" %s : %d/%d HP\n", ennemi.Name, ennemi.HP, ennemi.MaxHP)
	ennemi.AfficherBarreDeVie("compact")
}
