package shop

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	armes "Project-RED-groupe-1/armes"
	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/monnaie"
)

func CraftArme(reader *bufio.Reader, eddies *monnaie.Eddies, inventory *inventaire.Inventory, delay int) {

	fmt.Print("Entrez le niveau de l’arme à crafter (1-5) : ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	niveau, err := strconv.Atoi(numStr)
	if err != nil || niveau < 1 || niveau > 5 {
		fmt.Println("Niveau invalide.")
		return
	}

	fmt.Println("Choisissez le type d’arme :")
	fmt.Println("1. Pistolet")
	fmt.Println("2. Fusil à pompe")
	fmt.Println("3. Mitraillette")
	fmt.Println("4. Fusil d’assaut")
	fmt.Println("5. Katana")
	fmt.Println("6. Couteau")
	fmt.Print("Votre choix : ")
	typeStr, _ := reader.ReadString('\n')
	typeStr = strings.TrimSpace(typeStr)
	typeChoix, err := strconv.Atoi(typeStr)
	if err != nil || typeChoix < 1 || typeChoix > 6 {
		fmt.Println("Type invalide.")
		return
	}

	cout := 10 + (niveau-1)*20
	componentName := fmt.Sprintf("Composant Niveau %d", niveau)

	if !inventory.HasItem(componentName) {
		fmt.Printf("Vous avez besoin de '%s' pour crafter cette arme.\n", componentName)
		return
	}

	if !eddies.Spend(cout) {
		fmt.Println("Pas assez d’eddies pour crafter.")
		return
	}

	inventory.RemoveItem(componentName)

	var arme armes.Arme
	switch typeChoix {
	case 1:
		switch niveau {
		case 1:
			arme = armes.Pistolet1
		case 2:
			arme = armes.Pistolet2
		case 3:
			arme = armes.Pistolet3
		case 4:
			arme = armes.Pistolet4
		case 5:
			arme = armes.Pistolet5
		}
	case 2:
		switch niveau {
		case 1:
			arme = armes.Fusilspompe1
		case 2:
			arme = armes.Fusilspompe2
		case 3:
			arme = armes.Fusilspompe3
		case 4:
			arme = armes.Fusilspompe4
		case 5:
			arme = armes.Fusilspompe5
		}
	case 3:
		switch niveau {
		case 1:
			arme = armes.Mitraillette
		case 2:
			arme = armes.Mitraillette2
		case 3:
			arme = armes.Mitraillette3
		case 4:
			arme = armes.Mitraillette4
		case 5:
			arme = armes.Mitraillette5
		}
	case 4:
		switch niveau {
		case 1:
			arme = armes.Fusil1
		case 2:
			arme = armes.Fusil2
		case 3:
			arme = armes.Fusil3
		case 4:
			arme = armes.Fusil4
		case 5:
			arme = armes.Fusil5
		}
	case 5:
		switch niveau {
		case 1:
			arme = armes.Katana1
		case 2:
			arme = armes.Katana2
		case 3:
			arme = armes.Katana3
		case 4:
			arme = armes.Katana4
		case 5:
			arme = armes.Katana5
		}
	case 6:
		switch niveau {
		case 1:
			arme = armes.Couteau1
		case 2:
			arme = armes.Couteau2
		case 3:
			arme = armes.Couteau3
		case 4:
			arme = armes.Couteau4
		case 5:
			arme = armes.Couteau5
		}
	}

	item := inventaire.Item{
		Nom:         fmt.Sprintf("%s (Niv %d)", arme.Nom, niveau),
		Type:        "arme",
		Effet:       arme.Degats,
		Consommable: false,
	}

	inventory.AddItem(item)
	fmt.Printf("Vous avez crafté %s (niveau %d) pour %d eddies. Composant consommé: %s\n", arme.Nom, niveau, cout, componentName)
}
