package shop

import (
	"bufio"
	"fmt"
	"strings"

	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/monnaie"
)

// VendreObjet vend un objet de l'inventaire et ajoute les eddies
func VendreObjet(reader *bufio.Reader, eddies *monnaie.Eddies, inventory *inventaire.Inventory, items []Item) {
	fmt.Println("Votre inventaire :")
	inventory.ShowInventory()

	fmt.Print("Entrez le nom exact de l’objet à vendre : ")
	objet, _ := reader.ReadString('\n')
	objet = strings.TrimSpace(objet)

	// Vérifier que l'objet est dans l'inventaire
	if !inventory.HasItem(objet) {
		fmt.Println("Objet introuvable dans l’inventaire.")
		return
	}

	// Déterminer le prix de vente (50% du prix d'achat)
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

	// Retirer l'objet de l'inventaire et ajouter les eddies
	inventory.RemoveItem(objet)
	eddies.Add(prixVente)

	fmt.Printf("Vous avez vendu %s pour %d eddies.\n", objet, prixVente)
}
