package main

import (
	armes "Project-RED-groupe-1/armes"
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
	character := player.NewPlayer()

	character.ChooseClass(choice)
	printlnSlow(fmt.Sprintf("\nTu es un %s !", character.Class), delay)

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
	character.AfficherBarreDeVie("compact")

	printlnSlow(fmt.Sprintf("\nTu commences l’aventure avec %d eddies en poche. Utilise-les avec sagesse ! Bien évidemment on t'a donné une arme attention elle a un taux de precision, et un hack pour bien commencer ton aventure. Bonne chance", eddies.GetBalance()), delay)

	printlnSlow("\nAppuie sur Entrée pour démarrer l'histoire...", delay)
	reader.ReadString('\n')
	inventory := inventaire.NewInventory()
	switch choice {
	case "1":
		histoire.CorpoHistoire(character)
		histoire.StartCorpo(character, inventory)
		combat.LancerCombat(character, combat.Agentcorpo, inventory)
		histoire.Acte1_Relic(character, inventory)

	case "2":
		histoire.NomadeHistoire()
		histoire.StartNomade(character, inventory)
		combat.LancerCombat(character, combat.Ncpd, inventory)
		histoire.Acte1_Relic(character, inventory)

	case "3":
		histoire.GosseHistoire()
		histoire.StartGosse(character, inventory)
		combat.LancerCombat(character, combat.Adam, inventory)
		histoire.Acte1_Relic(character, inventory)

		if choice == "1" || choice == "2" || choice == "3" {
		}
		printlnSlow("Veuillez entrer 1, 2 ou 3", delay)
	}
	AfficherMenu(character, eddies, inventory, delay, reader)
	inventory.ShowInventory()

}
