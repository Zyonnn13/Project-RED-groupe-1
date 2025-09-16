package main

import (
	shop "Project-RED-groupe-1/Shop"
	"Project-RED-groupe-1/histoire"
	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/player"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func printlnSlow(text string, delay time.Duration) {
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(delay)
	}
	fmt.Println()
}

func printlnLineByLine(text string, delay time.Duration) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fmt.Println(line)
		time.Sleep(delay)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	delay := 20 * time.Millisecond

	ascii := `
________/\\\\\\\\\________________/\\\___________________________________________________________________________________________________________________________                       
 _____/\\\////////________________\/\\\________________________________________________________________________________/\\\_______________________________________                      
  ___/\\\/______________/\\\__/\\\_\/\\\_______________________________________/\\\\\\\\\______________________________\/\\\_______________________________________                     
   __/\\\_______________\//\\\/\\\__\/\\\____________/\\\\\\\\___/\\/\\\\\\\___/\\\/////\\\__/\\\____/\\\__/\\/\\\\\\___\/\\\\\\\\__________________________________                    
    _\/\\\________________\//\\\\\___\/\\\\\\\\\____/\\\/////\\\_\/\\\/////\\\_\/\\\\\\\\\\__\/\\\___\/\\\_\/\\\////\\\__\/\\\////\\\________________________________                   
     _\//\\\________________\//\\\____\/\\\////\\\__/\\\\\\\\\\\__\/\\\___\///__\/\\\//////___\/\\\___\/\\\_\/\\\__\//\\\_\/\\\\\\\\/_________________________________                  
      __\///\\\___________/\\_/\\\_____\/\\\__\/\\\_\//\\///////___\/\\\_________\/\\\_________\/\\\___\/\\\_\/\\\___\/\\\_\/\\\///\\\_________________________________                 
       ____\////\\\\\\\\\_\//\\\\/______\/\\\\\\\\\___\//\\\\\\\\\\_\/\\\_________\/\\\_________\//\\\\\\\\\__\/\\\___\/\\\_\/\\\_\///\\\_______________________________                
        _______\/////////___\////________\/////////_____\//////////__\///__________\///___________\/////////___\///____\///__\///____\///________________________________               
 ________________________________________________________________________________________________________/\\\\\\\\\_________/\\\\\\\_____/\\\\\\\\\\\\\\\__/\\\\\\\\\\\\\\\_            
  ______________________________________________________________________________________________________/\\\///////\\\_____/\\\/////\\\__\/////////////\\\_\/////////////\\\_           
   _____________________________________________________________________________________________________\///______\//\\\___/\\\____\//\\\____________/\\\/_____________/\\\/__          
    _______________________________________________________________________________________________________________/\\\/___\/\\\_____\/\\\__________/\\\/_____________/\\\/____         
     ____________________________________________________________________________________________________________/\\\//_____\/\\\_____\/\\\________/\\\/_____________/\\\/______        
      _________________________________________________________________________________________________________/\\\//________\/\\\_____\/\\\______/\\\/_____________/\\\/________       
       _______________________________________________________________________________________________________/\\\/___________\//\\\____/\\\_____/\\\/_____________/\\\/__________      
        ______________________________________________________________________________________________________/\\\\\\\\\\\\\\\__\///\\\\\\\/____/\\\/_____________/\\\/____________ 
         _____________________________________________________________________________________________________\///////////////_____\///////_____\///______________\///______________
`
	yellow := "\033[33m"
	reset := "\033[0m"
	green := "\033[32m"
	printlnLineByLine(yellow+ascii+reset, 100*time.Millisecond)

	printlnSlow("==== Bienvenue dans Cyberpunk 2077 ====", delay)
	printlnSlow("Choisis ta classe au sein de Night City", delay)

	p := &player.Player{}

	printlnSlow("1 - Corpo", delay)
	printlnSlow(green+"  Tu es un employé ambitieux d’Arasaka, spécialisé dans la sécurité interne. Tu as accès à des informations sensibles, mais ton supérieur te confie une mission qui pourrait te coûter ta carrière… ou ta vie."+reset, delay)

	printlnSlow("2 - Nomade", delay)
	printlnSlow(green+" Tu viens des Badlands, loin de la corruption de la ville. Ton clan t’a confié une mission : faire passer une cargaison illégale à travers les checkpoints de Night City."+reset, delay)

	printlnSlow("3 - Gosse de rue", delay)
	printlnSlow(green+" Tu as grandi dans les ruelles de Heywood. Tu connais les gangs, les deals, et comment survivre. Mais aujourd’hui, un vieil ami te demande un service dangereux."+reset, delay)

	var choice string
	for {
		fmt.Print("Ton choix: ")
		choice, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "1" || choice == "2" || choice == "3" {
			break
		}
		printlnSlow("Veuillez entrer 1, 2 ou 3", delay)
	}

	p.ChooseClass(choice)
	printlnSlow(fmt.Sprintf("\nTu es un %s !", p.Class), delay)

	character := player.NewDesignplayer()

	printlnSlow("\n╔════════════════════════════════════════════════════════════════════╗", 5*time.Millisecond)
	printlnSlow("║                        Ton personnage final                        ║", 5*time.Millisecond)
	fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %-15s │ %-15s │ %-15s │ %-15s \n", "Nom", "Cheveux", "Yeux", "Taille")
	fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %-15s │ %-15s │ %-15s │ %-15s \n",
		character.Name, character.Hair, character.Eyes, character.Height)
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")

	joueur := player.Joueur{
		Nom:      character.Name,
		Sante:    100,
		SanteMax: 100,
	}

	joueur.AfficherBarreDeSante()

	printlnSlow("\nAppuie sur Entrée pour démarrer l'histoire...", delay)
	reader.ReadString('\n')

	switch choice {
	case "1":
		histoire.CorpoHistoire()
	case "2":
		histoire.NomadeHistoire()
	case "3":
		histoire.GosseHistoire()

		if choice == "1" || choice == "2" || choice == "3" {
			break
		}
		printlnSlow("Veuillez entrer 1, 2 ou 3", delay)
	}

	inventory := inventaire.NewInventory()
	inventory.Additem("maxdoc")
	inventory.Showinventory()

	for {
		printlnSlow("\n===== MENU PRINCIPAL =====", delay)
		printlnSlow("1. Afficher les informations du personnage", delay)
		printlnSlow("2. Accéder au contenu de l’inventaire", delay)
		printlnSlow("3. Accéder à la Boutique", delay)
		printlnSlow("4. Quitter", delay)
		fmt.Print("Votre choix : ")

		menuChoice, _ := reader.ReadString('\n')
		menuChoice = strings.TrimSpace(menuChoice)

		switch menuChoice {
		case "1":
			printlnSlow("\n--- INFOS PERSONNAGE ---", delay)
			fmt.Printf("Nom : %s\n", joueur.Nom)
			fmt.Printf("Classe : %s\n", p.Class)
			fmt.Printf("Santé : %d/%d\n", joueur.Sante, joueur.SanteMax)
			printlnSlow("Appuie sur Entrée pour revenir au menu.", delay)
			reader.ReadString('\n')

		case "2":
			printlnSlow("\n--- INVENTAIRE ---", delay)
			inventory := inventaire.NewInventory()
			inventory.Showinventory()
			printlnSlow("Appuie sur Entrée pour revenir au menu.", delay)
			reader.ReadString('\n')

		case "3":
			printlnSlow("\n--- BOUTIQUE ---", delay)

			items := []shop.Item{
				shop.Maxdoc,
				shop.Revitalisant,
				shop.Frag,
				shop.Flash,
				shop.Redémarrage,
				shop.Surchauffe,
				shop.Circuit,
			}

			for {
				fmt.Println("\n===== MENU BOUTIQUE =====")
				for i, item := range items {
					fmt.Printf("%d. %s - %d eddies\n", i+1, item.Nom, item.Prix)
				}
				fmt.Println("A. Afficher les détails d’un objet")
				fmt.Println("B. Acheter un objet")
				fmt.Println("R. Revenir au menu principal")
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
						fmt.Printf("\nNom : %s\n", item.Nom)
						fmt.Printf("Prix : %d eddies\n", item.Prix)
						fmt.Printf("Description : %s\n", item.Description)
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
						inventory.Additem(item.Nom)
						fmt.Printf("Vous avez acheté %s !\n", item.Nom)
					} else {
						fmt.Println("Numéro invalide.")
					}
					fmt.Println("Appuie sur Entrée pour continuer.")
					reader.ReadString('\n')

				case "R":
					fmt.Println("Retour au menu principal...")

				default:
					fmt.Println("Choix invalide. Veuillez réessayer.")
				}

				if shopChoice == "R" {
					break
				}
			}

		case "4":
			printlnSlow("\n--- QUITTER ---", delay)
			printlnSlow("Appuie sur Entrée pour quitter le jeu.", delay)
			reader.ReadString('\n')
			return

		default:
			printlnSlow("Choix invalide. Veuillez réessayer.", delay)
		}
	}

}
