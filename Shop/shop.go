package shop

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

//craft d'arme et pas d'achat en boutique objet de craft composant de niveau 1 2 3 4 ou 5
