package shop

type Item struct {
	Nom         string
	Prix        int
	Description string
	Consommable bool
}

var Maxdoc = Item{
	Nom:         "Maxdoc M.K 1",
	Prix:        50,
	Description: "Le MaxDoc est un Inhalateur fabriqué par la Trauma Team. Il permet à l'utilisateur de restaurer 50 PV lorsuq'il est utilisé.",
	Consommable: true,
}

var Revitalisant = Item{
	Nom:         "Revitalisant M.K 1",
	Prix:        60,
	Description: "Le revitalisant est une injection permetant à son utilisateur de restaurer instantanément 30 PV, puis de restaurer 30 PV de nouveau au tour suivant",
	Consommable: true,
}

var Frag = Item{
	Nom:         "La grenade à fragmentation F-GX",
	Prix:        120,
	Description: "Cette grenade permet à son utilisateur d'infliger des dégâts à tous les ennemis. Dégâts : modérés",
	Consommable: true,
}

var Flash = Item{
	Nom:         "Grenade aveuglante X-22",
	Prix:        150,
	Description: "Cette grenade permet à son utilisateur d'aveugler tous les ennemis ce qui les empêche d'attaquer au tour suivant. Les ennemis sont capable d'annnuler cet effet et il est inefficace contre le boss final.",
	Consommable: true,
}

var Redémarrage = Item{
	Nom:         "Redémarrage optique",
	Prix:        500,
	Description: "Ce hack permet à son utilisateur de couper la vison d'un ennemi ce qui l'empêche d'attaquer au tour suivant. Cet effet est inefficace contre le boss final.",
	Consommable: false,
}

var Surchauffe = Item{
	Nom:         "Surchauffe",
	Prix:        500,
	Description: "Ce hack permet à son utilisateur de faire surchauffer les implants de son ennemi lui infligeant des dégâts sur 3 tours. Dégâts : légers.",
	Consommable: false,
}

var Circuit = Item{
	Nom:         "Court circuit",
	Prix:        1000,
	Description: "Ce hack permet à son utilisateur de créer un court circuit chez l'ennemi. Dégâts : élevé.",
	Consommable: false,
}

//craft d'arme et pas d'achat en boutique objet de craft composant de niveau 1 2 3 4 ou 5
