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
