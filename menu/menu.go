package menu

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	shop "Project-RED-groupe-1/Shop"
	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/monnaie"
	"Project-RED-groupe-1/player"
)

func printlnSlow(text string, delay time.Duration) {
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(delay)
	}
	fmt.Println()
}

func AfficherMenu(character *player.Player, eddies *monnaie.Eddies, inventory *inventaire.Inventory, delay time.Duration, reader *bufio.Reader) {
	for {
		fmt.Println("\n===== MENU PRINCIPAL =====")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l’inventaire")
		fmt.Println("3. Accéder à la Boutique")
		fmt.Println("4. Revenir au jeu")
		fmt.Print("Votre choix : ")

		menuChoice, _ := reader.ReadString('\n')
		menuChoice = strings.TrimSpace(strings.ToUpper(menuChoice))

		switch menuChoice {
		case "1":
			fmt.Printf("\nNom : %s\nClasse : %s\nSanté : %d/%d\nEddies : %d\n",
				character.Name, character.Class, character.HP, character.MaxHP, eddies.GetBalance())
			fmt.Println("Appuie sur Entrée pour revenir au menu.")
			reader.ReadString('\n')

		case "2":
			inventory.ShowInventory()
			fmt.Println("Appuie sur Entrée pour revenir au menu.")
			reader.ReadString('\n')

		case "3":
			ouvrirBoutique(reader, eddies, inventory, delay)

		case "4":
			return

		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}

func ouvrirBoutique(reader *bufio.Reader, eddies *monnaie.Eddies, inventory *inventaire.Inventory, delay time.Duration) {
	items := []shop.Item{
		shop.Maxdoc,
		shop.Revitalisant,
		shop.Frag,
		shop.Flash,
		shop.Redemarrage,
		shop.Surchauffe,
		shop.Circuit,
	}

	for {
		printlnSlow("\n===== MENU BOUTIQUE =====", delay)
		for i, item := range items {
			fmt.Printf("%d. %s - %d eddies\n", i+1, item.Nom, item.Prix)
		}
		fmt.Println("\nA. Afficher les détails d’un objet")
		fmt.Println("B. Acheter un objet")
		fmt.Println("R. Revenir au menu principal")
		fmt.Println("C. Crafter une arme")
		fmt.Print("Votre choix : ")

		shopChoice, _ := reader.ReadString('\n')
		shopChoice = strings.TrimSpace(strings.ToUpper(shopChoice))

		switch shopChoice {
		case "A":
			fmt.Print("Entrez le numéro de l’objet pour voir les détails : ")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			index := -1
			fmt.Sscanf(numStr, "%d", &index)
			if index >= 1 && index <= len(items) {
				item := items[index-1]
				fmt.Printf("\nNom : %s\nPrix : %d eddies\nDescription : %s\n", item.Nom, item.Prix, item.Description)
				if item.Consommable {
					fmt.Println("Type : Consommable")
				} else {
					fmt.Println("Type : Hack")
				}
			} else {
				fmt.Println("Numéro invalide.")
			}
			fmt.Println("Appuie sur Entrée pour continuer.")
			reader.ReadString('\n')

		case "B":
			fmt.Print("Entrez le numéro de l’objet à acheter : ")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			index := -1
			fmt.Sscanf(numStr, "%d", &index)
			if index >= 1 && index <= len(items) {
				item := items[index-1]
				if eddies.Spend(item.Prix) {
					// ✅ Ajouter l’objet directement à l’inventaire
					inventory.AddItem(inventaire.Item{
						Nom:         item.Nom,
						Description: item.Description,
						Type:        item.Type,
						Effet:       item.Effet,
						Consommable: item.Consommable,
					})
					printlnSlow(fmt.Sprintf("Vous avez acheté %s pour %d eddies.", item.Nom, item.Prix), delay)
					fmt.Printf("Eddies restants : %d eddies\n", eddies.GetBalance())
				} else {
					printlnSlow("Vous n’avez pas assez d’eddies pour cet achat.", delay)
				}
			} else {
				fmt.Println("Numéro invalide.")
			}
			fmt.Println("Appuie sur Entrée pour continuer.")
			reader.ReadString('\n')

		case "C":
			shop.CraftArme(reader, eddies, inventory, int(delay))
			fmt.Println("Appuie sur Entrée pour continuer.")
			reader.ReadString('\n')

		case "R":
			return

		default:
			printlnSlow("Choix invalide. Veuillez réessayer.", delay)
		}
	}
}
