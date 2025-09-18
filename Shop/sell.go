package shop

import (
	"bufio"
	"fmt"
	"strings"
	"Project-RED-groupe-1/Shop/shop"
	"Project-RED-groupe-1/monnaie"
	"Project-RED-groupe-1/Inventaire"
)

func VendreObjet(reader *bufio.Reader, eddies *monnaie.Eddies, inventory *inventaire.Inventory, items []shop.Item) {
	fmt.Println("Votre inventaire :")
	inventory.Showinventory()

	fmt.Print("Entrez le nom exact de l’objet à vendre : ")
	objet, _ := reader.ReadString('\n')
	objet = strings.TrimSpace(objet)

	if !inventory.HasItem(objet) {
		fmt.Println("Objet introuvable dans l’inventaire.")
		return
	}

	prixVente := 0
	for _, it := range items {
		if it.Nom == objet {
			prixVente = it.Prix / 2
			break
		}
	}

	if prixVente <= 0 {
		fmt.Println("Cet objet ne peut pas être vendu.")
		return
	}

	inventory.Removeitem(objet)
	eddies.Add(prixVente)
	fmt.Printf("Vous avez vendu %s pour %d eddies.\n", objet, prixVente)
}