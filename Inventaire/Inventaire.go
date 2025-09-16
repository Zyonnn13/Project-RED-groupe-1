package inventaire

import (
	"Project-RED-groupe-1/player"
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

func (inv *Inventory) UseItem(item string, player *player.Designplayer) bool {
	for i, obj := range inv.Items {
		if obj == item {
			// Appliquer l'effet
			switch item {
			case "Potion de soin":
				player.HP += 20
				if player.HP > player.MaxHP {
					player.HP = player.MaxHP
				}
				fmt.Println("🧪 Vous utilisez Potion de soin")
				fmt.Printf("❤️ PV de %s : %d / %d\n", player.Name, player.HP, player.MaxHP)

			case "Boost d'attaque":
				player.Attack += 5
				fmt.Println("💪 Vous utilisez Boost d'attaque")
				fmt.Printf("⚔️ Attaque de %s : %d\n", player.Name, player.Attack)

			default:
				fmt.Println("❌ Objet inconnu, aucun effet.")
				return false
			}

			// Supprimer l'objet utilisé
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			return true
		}
	}
	fmt.Println("❌ Objet non trouvé dans l'inventaire.")
	return false
}
