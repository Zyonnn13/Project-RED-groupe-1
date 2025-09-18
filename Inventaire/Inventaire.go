package inventaire

import (
	"Project-RED-groupe-1/player"
	"fmt"
)

type Item struct {
	Nom         string
	Description string
	Type        string
	Effet       int
	Consommable bool
}

type Inventory struct {
	Items []Item
}

func NewInventory() *Inventory {
	return &Inventory{Items: []Item{}}
}

func (inv *Inventory) AddItem(item Item) {
	inv.Items = append(inv.Items, item)
	fmt.Printf("📦 Tu as ajouté %s à ton inventaire.\n", item.Nom)
}

func (inv *Inventory) ShowInventory() {
	fmt.Println("\n📦 Inventaire :")
	if len(inv.Items) == 0 {
		fmt.Println("  (vide)")
		return
	}
	for i, item := range inv.Items {
		fmt.Printf("  %d. %s - %s\n", i+1, item.Nom, item.Description)
	}
}

func (inv *Inventory) UseItem(nom string, p *player.Player) bool {
	for i, item := range inv.Items {
		if item.Nom == nom {
			switch item.Type {
			case "soin":
				p.HP += item.Effet
				if p.HP > p.MaxHP {
					p.HP = p.MaxHP
				}
				fmt.Printf("🧪 %s utilise %s (+%d PV)\n", p.Name, item.Nom, item.Effet)
				p.AfficherBarreDeVie("compact")
			case "boost":
				p.Attack += item.Effet
				fmt.Printf("💪 %s utilise %s (+%d ATK)\n", p.Name, item.Nom, item.Effet)
			default:
				fmt.Println("❌ Effet inconnu pour", item.Nom)
				return false
			}
			if item.Consommable {
				inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			}
			return true
		}
	}
	fmt.Println("❌ Objet non trouvé dans l'inventaire.")
	return false
}
