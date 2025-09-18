package inventaire

import (
	"Project-RED-groupe-1/player"
	"fmt"
)

<<<<<<< HEAD
// Définition du type Item
type Item struct {
	Nom         string
	Prix        int
	Description string
	Consommable bool
}

// Inventaire contenant des objets complets
=======
type Item struct {
	Nom         string
	Description string
	Type        string
	Effet       int
	Consommable bool
}

>>>>>>> 6201e347e5deb15fd7f486b9883ced1449981a1b
type Inventory struct {
	Items []Item
}

// Constructeur
func NewInventory() *Inventory {
	return &Inventory{Items: []Item{}}
}

<<<<<<< HEAD
// Ajouter un objet
func (inv *Inventory) AddItem(obj Item) {
	inv.Items = append(inv.Items, obj)
	fmt.Println("✅ Tu as ajouté", obj.Nom, "à ton inventaire.")
=======
func (inv *Inventory) AddItem(item Item) {
	inv.Items = append(inv.Items, item)
	fmt.Printf("📦 Tu as ajouté %s à ton inventaire.\n", item.Nom)
>>>>>>> 6201e347e5deb15fd7f486b9883ced1449981a1b
}

// Afficher l'inventaire
func (inv *Inventory) ShowInventory() {
	fmt.Println("\n📦 Inventaire :")
	if len(inv.Items) == 0 {
		fmt.Println("  (vide)")
		return
	}
<<<<<<< HEAD
	for i, obj := range inv.Items {
		fmt.Printf("  %d. %s - %s\n", i+1, obj.Nom, obj.Description)
	}
}

// Utiliser un objet consommable
func (inv *Inventory) UseItem(nom string, p *player.Player) bool {
	for i, obj := range inv.Items {
		if obj.Nom == nom {
			if !obj.Consommable {
				fmt.Println("❌ Cet objet n'est pas consommable.")
				return false
			}

			switch obj.Nom {
			case "Potion de soin":
				soin := 20
				p.HP += soin
				if p.HP > p.MaxHP {
					p.HP = p.MaxHP
				}
				fmt.Printf("🧪 %s utilise %s (+%d HP)\n", p.Name, obj.Nom, soin)
				fmt.Printf("❤️ PV de %s : %d / %d\n", p.Name, p.HP, p.MaxHP)

			case "Boost d'attaque":
				p.Attack += 5
				fmt.Printf("💪 %s utilise %s (+5 ATK)\n", p.Name, obj.Nom)

			default:
				fmt.Println("❌ Effet inconnu pour", obj.Nom)
				return false
			}

			// Retirer l'objet après utilisation
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
=======
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
>>>>>>> 6201e347e5deb15fd7f486b9883ced1449981a1b
			return true
		}
	}
	fmt.Println("❌ Objet non trouvé dans l'inventaire.")
	return false
}

// Retirer un objet manuellement
func (inv *Inventory) RemoveItem(nom string) bool {
	for i, obj := range inv.Items {
		if obj.Nom == nom {
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			fmt.Println("🗑️", obj.Nom, "a été retiré de l'inventaire.")
			return true
		}
	}
	fmt.Println("❌ Impossible de retirer", nom, ": objet non trouvé.")
	return false
}

// Afficher uniquement les consommables
func (inv *Inventory) ShowConsumables() {
	fmt.Println("\n🧪 Objets consommables :")
	found := false
	for i, obj := range inv.Items {
		if obj.Consommable {
			fmt.Printf("  %d. %s - %s\n", i+1, obj.Nom, obj.Description)
			found = true
		}
	}
	if !found {
		fmt.Println("  Aucun objet consommable.")
	}
}

// Compter les objets
func (inv *Inventory) Count() int {
	return len(inv.Items)
}
