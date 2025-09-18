package main

import (
	armes "Project-RED-groupe-1/Armes"
	shop "Project-RED-groupe-1/Shop"
	"Project-RED-groupe-1/combat"
	"Project-RED-groupe-1/histoire"
	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/monnaie"
	"Project-RED-groupe-1/player"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func printlnSlow(text string, delay time.Duration) {
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(delay)
	}
	fmt.Println()
}

func sound() {
	for {
		f, err := os.Open("sound/musique.mp3")
		if err != nil {
			fmt.Println("Erreur ouverture fichier audio :", err)
			return
		}
		streamer, format, err := mp3.Decode(f)
		if err != nil {
			fmt.Println("Erreur décodage audio :", err)
			f.Close()
			return
		}
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Millisecond*50))
		done := make(chan bool)
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			streamer.Close()
			f.Close()
			done <- true
		})))
		<-done
	}
}

func soundplay() {
	go sound()
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
	soundplay()
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

	character := player.NewPlayer()
	character.Arme = armes.Pistolet1

	printlnSlow("\n╔════════════════════════════════════════════════════════════════════╗", 5*time.Millisecond)
	printlnSlow("║                        Ton personnage final                        ║", 5*time.Millisecond)
	fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %-15s │ %-15s │ %-15s │ %-15s \n", "Nom", "Cheveux", "Yeux", "Taille")
	fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %-15s │ %-15s │ %-15s │ %-15s \n",
		character.Name, character.Hair, character.Eyes, character.Height)
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")

	eddies := monnaie.NewEddies(200)
	character.HP = 100
	character.MaxHP = 100
	character.AfficherBarreDeSante()

	printlnSlow(fmt.Sprintf("\nTu commences l’aventure avec %d eddies en poche. Utilise-les avec sagesse ! Bien évidemment on t'a donné une arme de débarquement avec 40 pourcent de précision. Bonne chance", eddies.GetBalance()), delay)

	printlnSlow("\nAppuie sur Entrée pour démarrer l'histoire...", delay)
	reader.ReadString('\n')

	switch choice {
	case "1":
		histoire.CorpoHistoire()
		histoire.StartCorpo(character)
		combat.LancerCombat(p, combat.Agentcorpo, &inventaire.Inventory{})

	case "2":
		histoire.NomadeHistoire()
		histoire.StartNomade(character)
		combat.LancerCombat(p, combat.Ncpd, &inventaire.Inventory{})
	case "3":
		histoire.GosseHistoire()
		histoire.StartGosse(character)
		combat.LancerCombat(p, combat.Adam, &inventaire.Inventory{})

		if choice == "1" || choice == "2" || choice == "3" {
			break
		}
		printlnSlow("Veuillez entrer 1, 2 ou 3", delay)
	}

	inventory := inventaire.NewInventory()
	inventory.AddItem("maxdoc")
	inventory.ShowInventory()

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
			fmt.Printf("Nom : %s\n", p.Name)
			fmt.Printf("Classe : %s\n", p.Class)
			fmt.Printf("Santé : %d/%d\n", character.HP, character.MaxHP)
			fmt.Printf("Eddies :%d\n ", eddies.GetBalance())
			printlnSlow("Appuie sur Entrée pour revenir au menu.", delay)
			reader.ReadString('\n')

		case "2":
			printlnSlow("\n=== INVENTAIRE ===", delay)
			inventory := inventaire.NewInventory()
			inventory.ShowInventory()
			printlnSlow("Appuie sur Entrée pour revenir au menu.", delay)
			reader.ReadString('\n')

		case "3":
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
				// 👇 Ajoute ce bloc juste après l'affichage des objets
				fmt.Print("\nEntrez le numéro de l’objet à acheter (ou appuyez sur Entrée pour annuler) : ")
				achatInput, _ := reader.ReadString('\n')
				achatInput = strings.TrimSpace(achatInput)

				if achatInput != "" {
					index := -1
					fmt.Sscanf(achatInput, "%d", &index)
					if index >= 1 && index <= len(items) {
						item := items[index-1]
						if eddies.Spend(item.Prix) {
							printlnSlow(fmt.Sprintf("Vous avez acheté %s pour %d crédits.", item.Nom, item.Prix), delay)
						} else {
							printlnSlow("Vous n’avez pas assez de crédits.", delay)
						}
					} else {
						printlnSlow("Numéro invalide.", delay)
					}
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
						if eddies.Spend(item.Prix) {
							inventory.AddItem(item.Nom)
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

				case "R":
					fmt.Println("Retour au menu principal...")

				default:
					fmt.Println("Choix invalide. Veuillez réessayer.")
				}
				if shopChoice == "R" {
					break
				}
			}

			// 👇 Ajoute ce bloc juste après l'affichage des objets
			fmt.Print("\nEntrez le numéro de l’objet à acheter (ou appuyez sur Entrée pour annuler) : ")
			achatInput, _ := reader.ReadString('\n')
			achatInput = strings.TrimSpace(achatInput)

			if achatInput != "" {
				index := -1
				fmt.Sscanf(achatInput, "%d", &index)
				if index >= 1 && index <= len(items) {
					item := items[index-1]
					if eddies.Spend(item.Prix) {
						printlnSlow(fmt.Sprintf("Vous avez acheté %s pour %d crédits.", item.Nom, item.Prix), delay)
					} else {
						printlnSlow("Vous n’avez pas assez de crédits.", delay)
					}
				} else {
					printlnSlow("Numéro invalide.", delay)
				}
			}

			printlnSlow("\nAppuie sur Entrée pour revenir au menu.", delay)
			reader.ReadString('\n')

		case "4":
			printlnSlow("\n=== QUITTER ===", delay)
			printlnSlow("Appuie sur Entrée pour quitter le jeu.", delay)
			reader.ReadString('\n')
			return

		default:
			printlnSlow("Choix invalide. Veuillez réessayer.", delay)
		}
	}

}
