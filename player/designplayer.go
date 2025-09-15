package player

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Designplayer struct {
	Name   string
	Hair   string
	Eyes   string
	Height string
}

func NewDesignplayer() *Designplayer {
	reader := bufio.NewReader(os.Stdin)
	p := &Designplayer{}

	fmt.Println("                                                                                    Création de ton personnage                                                                ")
	fmt.Println()
	fmt.Println()

	//Nom
	for {
		fmt.Print(" Avant toutes choses tu dois choisir ton nom de personnage : ")
		name, _ := reader.ReadString('\n')
		p.Name = strings.TrimSpace(name)

		if len(name) > 0 && unicode.IsUpper(rune(name[0])) {
			p.Name = name
			break
		} else {
			fmt.Println(" Le nom doit commencer par une Majuscule!!!!!")
		}
	}

	//Couleur de cheveux

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
	Hair := strings.TrimSpace(HairInput)
	switch Hair {
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
	}

	//Couleur des yeux

	fmt.Println("\n Choisis la couleur de tes yeux.")
	fmt.Println(" 1 : Noir ")
	fmt.Println(" 2 : Rouge ")
	fmt.Println(" 3 : Marron ")
	fmt.Println(" 4 : Bleu ")
	fmt.Println(" 5 : Rose ")
	fmt.Println(" 6 : Vert ")
	EyesInput, _ := reader.ReadString('\n')
	Eyes := strings.TrimSpace(EyesInput)
	switch Eyes {
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
	}

	// Taille

	fmt.Println("\n Choisis ta taille.")
	fmt.Println(" 1 : Petit ")
	fmt.Println(" 2 : Moyen ")
	fmt.Println(" 3 : Grand ")
	HeightInput, _ := reader.ReadString('\n')
	Height := strings.TrimSpace(HeightInput)
	switch Height {
	case "1":
		p.Height = "Petit"
	case "2":
		p.Height = "Moyen"
	case "3":
		p.Height = "Grand"
	}
	return p
}
