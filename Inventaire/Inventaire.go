package inventaire

import (
	"fmt"
)

type Inventory struct {
	Items []string
}

func NewInventory() *Inventory {
	return &Inventory{
		Items: []string{},
	}
}

func (inv *Inventory) Additem(item string) {
	inv.Items = append(inv.Items, item)
	fmt.Println("Tu as ajouté", item)
}

func (inv *Inventory) Showinventory() {
	fmt.Println("\nInventaire :")
	if len(inv.Items) == 0 {
		fmt.Println("  (vide)")
	} else {
		for i, item := range inv.Items {
			fmt.Printf("  %d. %s\n", i+1, item)
		}
	}
}
