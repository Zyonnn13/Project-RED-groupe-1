package player

import (
	armes "Project-RED-groupe-1/armes"
	"Project-RED-groupe-1/monnaie"
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Player struct {
	Name   string
	Hair   string
	Eyes   string
	Height string
	MaxHP  int
	HP     int
	Attack int
	Class  string
	Arme   armes.Arme
	Eddies *monnaie.Eddies
}


func NewPlayer() *Player {
	reader := bufio.NewReader(os.Stdin)
	p := &Player{}
	p.Eddies = monnaie.NewEddies(0)

	fmt.Println("                                                                                    Création de ton personnage                                                                ")
	fmt.Println()

	
	for {
		fmt.Print(" Avant toutes choses tu dois choisir ton nom de personnage : ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		valid := true
		for _, r := range name {
			if !unicode.IsLetter(r) {
				valid = false
				break
			}
		}

		if valid && len(name) > 0 {
			name = strings.ToLower(name)
			name = strings.Title(name)
			p.Name = name
			break
		} else {
			fmt.Println(" Le nom doit contenir uniquement des lettres !")
		}
	}

	
	for {
		fmt.Println("\n  Choisis ta couleur de cheveux.")
		fmt.Println(" 1 : Noir ")
		fmt.Println(" 2 : Rouge ")
		fmt.Println(" 3 : Marron ")
		fmt.Println(" 4 : Violet ")
		fmt.Println(" 5 : Bleu ")
		fmt.Println(" 6 : Blond ")
		fmt.Println(" 7 : Rose ")
		fmt.Println(" 8 : Vert ")
		HairInput, _ := reader.ReadString('\n')
		switch strings.TrimSpace(HairInput) {
		case "1":
			p.Hair = "Noir"
		case "2":
			p.Hair = "Rouge"
		case "3":
			p.Hair = "Marron"
		case "4":
			p.Hair = "Violet"
		case "5":
			p.Hair = "Bleu"
		case "6":
			p.Hair = "Blond"
		case "7":
			p.Hair = "Rose"
		case "8":
			p.Hair = "Vert"
		default:
			fmt.Println("choix invalide.Merci dentree un nombre valide")
			continue
		}
		break
	}
	
	for {
		fmt.Println("\n Choisis la couleur de tes yeux.")
		fmt.Println(" 1 : Noir ")
		fmt.Println(" 2 : Rouge ")
		fmt.Println(" 3 : Marron ")
		fmt.Println(" 4 : Bleu ")
		fmt.Println(" 5 : Rose ")
		fmt.Println(" 6 : Vert ")
		EyesInput, _ := reader.ReadString('\n')
		switch strings.TrimSpace(EyesInput) {
		case "1":
			p.Eyes = "Noir"
		case "2":
			p.Eyes = "Rouge"
		case "3":
			p.Eyes = "Marron"
		case "4":
			p.Eyes = "Bleu"
		case "5":
			p.Eyes = "Rose"
		case "6":
			p.Eyes = "Vert"
		default:
			fmt.Println("Choix invalide. Merci d'entrer un nombre valide!")
			continue
		}
		break
	}
	
	for {
		fmt.Println("\n Choisis ta taille.")
		fmt.Println(" 1 : Petit ")
		fmt.Println(" 2 : Moyen ")
		fmt.Println(" 3 : Grand ")
		HeightInput, _ := reader.ReadString('\n')
		switch strings.TrimSpace(HeightInput) {
		case "1":
			p.Height = "Petit"
		case "2":
			p.Height = "Moyen"
		case "3":
			p.Height = "Grand"
		default:
			fmt.Println("Choix invalide. Merci d'entrer un nombre valide!")
			continue
		}
		break
	}
	return p
}


func (p *Player) ChooseClass(choice string) {
	switch strings.TrimSpace(choice) {
	case "1":
		p.Class = "Corpo"
		p.MaxHP = 100
		p.Attack = 20
		p.Arme = armes.Pistolet1
	case "2":
		p.Class = "Nomade"
		p.MaxHP = 100
		p.Attack = 30
		p.Arme = armes.Pistolet1
	case "3":
		p.Class = "Gosse des rues"
		p.MaxHP = 100
		p.Attack = 30
		p.Arme = armes.Pistolet1
	}
	p.HP = p.MaxHP

}
func (p *Player) AfficherBarreDeVie(style string) {
	barLength := 20
	filled := (p.HP * barLength) / p.MaxHP

	var bar string
	switch style {
	case "compact":
		bar = strings.Repeat("█", filled) + strings.Repeat("░", barLength-filled)
		fmt.Printf("HP: [%s] %d/%d\n", bar, p.HP, p.MaxHP)
	case "classic":
		fmt.Printf("\nSanté : %d/%d\n", p.HP, p.MaxHP)
		fmt.Print("[")
		for i := 0; i < barLength; i++ {
			if i < filled {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("]")
	default:
		fmt.Println("Style inconnu.")
	}
}
