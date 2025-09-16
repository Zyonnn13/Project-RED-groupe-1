package combat

type Ennemis struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

var Opposant = Ennemis{"Opposant", 100, 100, 50}
var Valentinos = Ennemis{"Valentinos", 100, 100, 50}
var Ncpd = Ennemis{"Ncpd", 100, 100, 50}
var Agentcorpo = Ennemis{"Agent Corpo", 100, 100, 50}
var Agentrivaux = Ennemis{"Agent Rival", 100, 100, 50}
var Kirk = Ennemis{"Kirk", 120, 120, 60}
var Agentsecu = Ennemis{"Agent de Sécurité", 100, 100, 50}
var Agentarasaka = Ennemis{"Agent Arasaka", 100, 100, 50}
var Drone = Ennemis{"Drone", 75, 75, 40}
var Wraiths = Ennemis{"Wraiths", 100, 100, 50}
var Tygerclaws = Ennemis{"Tyger Claws", 100, 100, 50}
var Chefclaws = Ennemis{"Chef des Tyger Claws", 150, 150, 100}
var Dronelourd = Ennemis{"Drone Lours", 90, 90, 50}
var Tourelleauto = Ennemis{"Tourelle Automatique", 10, 10, 200}
var Sniper = Ennemis{"Sniper", 50, 50, 100}
var Netrunner = Ennemis{"Netrunner", 50, 50, 100}
var Adam = Ennemis{"Adam Smasher", 200, 200, 200}
