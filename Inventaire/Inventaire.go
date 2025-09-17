package inventaire

import (
	"Project-RED-groupe-1/player"
	"fmt"
)

type Inventory struct {
	Items []string
}

func NewInventory() *Inventory {
	return &Inventory{Items: []string{}}
}

func (inv *Inventory) AddItem(item string) {
	inv.Items = append(inv.Items, item)
	fmt.Println("Tu as ajouté", item)
}

func (inv *Inventory) ShowInventory() {
	fmt.Println("\n📦 Inventaire :")
	if len(inv.Items) == 0 {
		fmt.Println("  (vide)")
		return
	}
	for i, obj := range inv.Items {
		fmt.Printf("  %d. %s\n", i+1, obj)
	}
}

func (inv *Inventory) UseItem(item string, p *player.Player) bool {
	for i, obj := range inv.Items {
		if obj == item {
			switch item {
			case "Potion de soin":
				soin := 20
				p.HP += soin
				if p.HP > p.MaxHP {
					p.HP = p.MaxHP
				}

				fmt.Printf("🧪 %s utilise Potion de soin (+%d HP)\n", p.Name, soin)

				fmt.Println("🧪 Vous utilisez un Maxdoc M.K 1")
				fmt.Printf("❤️ PV de %s : %d / %d\n", p.Name, p.HP, p.MaxHP)

			case "Boost d'attaque":
				p.Attack += 5
				fmt.Printf("💪 %s utilise Boost d'attaque (+5 ATK)\n", p.Name)

			default:
				fmt.Println("❌ Effet inconnu pour", item)
				return false
			}
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			return true
		}
	}
	fmt.Println("❌ Objet non trouvé dans l'inventaire.")
	return false
}
