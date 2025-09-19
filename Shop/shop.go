package shop

import (
	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/monnaie"
	"bufio"
	"fmt"
	"strings"
	"time"
)

type Item struct {
	Nom         string
	Prix        int
	Description string
	Consommable bool
	Effet       int
	Type        string
}

var Maxdoc = Item{
	Nom:         "Maxdoc M.K 1",
	Prix:        50,
	Description: "Le MaxDoc est un inhalateur fabriqué par la Trauma Team. Il permet à l'utilisateur de restaurer 50 PV lorsqu'il est utilisé.",
	Consommable: true,
	Effet:       50,
	Type:        "soin",
}

var Revitalisant = Item{
	Nom:         "Revitalisant M.K 1",
	Prix:        60,
	Description: "Le revitalisant est une injection permettant de restaurer instantanément 30 PV, puis de restaurer 30 PV de nouveau au tour suivant.",
	Consommable: true,
}

var Frag = Item{
	Nom:         "Grenade à fragmentation F-GX",
	Prix:        120,
	Description: "Cette grenade inflige des dégâts modérés à tous les ennemis.",
	Consommable: true,
	Effet:       30,
	Type:        "dégâts",
}

var Flash = Item{
	Nom:         "Grenade aveuglante X-22",
	Prix:        150,
	Description: "Aveugle tous les ennemis, les empêchant d'attaquer au tour suivant. Inefficace contre le boss final.",
	Consommable: true,
	Effet:       0,
	Type:        "controle",
}

var Redemarrage = Item{
	Nom:         "Redémarrage optique",
	Prix:        500,
	Description: "Coupe la vision d'un ennemi, l'empêchant d'attaquer au tour suivant. Inefficace contre le boss final.",
	Consommable: false,
}

var Surchauffe = Item{
	Nom:         "Surchauffe",
	Prix:        500,
	Description: "Fait surchauffer les implants d'un ennemi, infligeant des dégâts légers sur 3 tours.",
	Consommable: false,
}

var Circuit = Item{
	Nom:         "Court-circuit",
	Prix:        1000,
	Description: "Crée un court-circuit chez l'ennemi. Dégâts élevés.",
	Consommable: false,
}

var Composant1 = Item{
	Nom:         "Composant Niveau 1",
	Prix:        10,
	Description: "Composant basique pour une arme de niveau 1.",
	Consommable: false,
}
var Composant2 = Item{
	Nom:         "Composant Niveau 2",
	Prix:        30,
	Description: "Composant intermédiaire pour une arme de niveau 2.",
	Consommable: false,
}
var Composant3 = Item{
	Nom:         "Composant Niveau 3",
	Prix:        50,
	Description: "Composant avancé pour une arme de niveau 3.",
	Consommable: false,
}
var Composant4 = Item{
	Nom:         "Composant Niveau 4",
	Prix:        70,
	Description: "Composant expert pour une arme de niveau 4.",
	Consommable: false,
}
var Composant5 = Item{
	Nom:         "Composant Niveau 5",
	Prix:        90,
	Description: "Composant ultime pour une arme de niveau 5.",
	Consommable: false,
}

func AfficherMenuBoutique(reader *bufio.Reader, eddies *monnaie.Eddies, inventory *inventaire.Inventory, delay time.Duration) {
	items := []Item{Maxdoc, Revitalisant, Frag, Flash, Redemarrage, Surchauffe, Circuit}

	for {
		fmt.Println("\n===== BOUTIQUE =====")
		for i, item := range items {
			fmt.Printf("%d. %s - %d eddies\n", i+1, item.Nom, item.Prix)
		}
		fmt.Println("C. Crafter une arme")
		fmt.Println("R. Revenir au menu")

		fmt.Print("Votre choix : ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(strings.ToUpper(choice))

		switch choice {
		case "C":
			CraftArme(reader, eddies, inventory, int(delay))
		case "R":
			return
		default:
			index := -1
			fmt.Sscanf(choice, "%d", &index)
			if index >= 1 && index <= len(items) {
				item := items[index-1]
				if eddies.Spend(item.Prix) {
					// Crée l’item pour l’inventaire
					newItem := inventaire.Item{
						Nom:         item.Nom,
						Description: item.Description,
						Type:        item.Type,
						Effet:       item.Effet,
						Consommable: item.Consommable,
					}
					inventory.AddItem(newItem)
					fmt.Printf("✅ Vous avez acheté %s pour %d eddies.\n", item.Nom, item.Prix)
				} else {
					fmt.Println("❌ Vous n’avez pas assez d’eddies pour cet achat.")
				}
			} else {
				fmt.Println("❌ Choix invalide.")
			}
		}
	}
}
