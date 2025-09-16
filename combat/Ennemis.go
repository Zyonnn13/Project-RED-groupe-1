package combat

type Valentinos struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type Opposants struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type NCPD struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type AgentCorpo struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type AgentRivaux struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type Kirk struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type Agentsecurité struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type AgentAraska struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type Drone struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type Wraiths struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type TygerClaws struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type ChefsClaws struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type Dronelourd struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

type TourelleAutomatique struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}
type Sniper struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}
type Netrunner struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}
type AdamSmacher struct {
	Name    string
	MaxHp   int
	Hp      int
	Attaque int
}

func IniValentinos() *Valentinos {
	return &Valentinos{
		Name:    "Valentinos",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}

func IniOpposants() *Opposants {
	return &Opposants{
		Name:    "Opposants",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}

func IniNCPD() *NCPD {
	return &NCPD{
		Name:    "NCPD",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}

func IniDrone() *Drone {
	return &Drone{
		Name:    "Drone",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}
func IniAgentCorpo() *AgentCorpo {
	return &AgentCorpo{
		Name:    "AgentCorpo",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}
func IniKirk() *Kirk {
	return &Kirk{
		Name:    "Kirk",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}
func IniSecurite() *Agentsecurité {
	return &Agentsecurité{
		Name:    "AgentSecurité",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}
func IniAgentArasaka() *AgentAraska {
	return &AgentAraska{
		Name:    "AgentArasaka",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}
func IniWraiths() *Wraiths {
	return &Wraiths{
		Name:    "Wraiths",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}
func IniTygerClaws() *TygerClaws {
	return &TygerClaws{
		Name:    "TygerClaws",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}
func IniChefsClaws() *ChefsClaws {
	return &ChefsClaws{
		Name:    "ChefsClaws",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}

func IniDronelourd() *Dronelourd {
	return &Dronelourd{
		Name:    "Dronelourd",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}

func IniTourelleauto() *TourelleAutomatique {
	return &TourelleAutomatique{
		Name:    "TourelleAutomatique",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}

func IniSniper() *Sniper {
	return &Sniper{
		Name:    "Sniper",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}

func IniNetrunner() *Netrunner {
	return &Netrunner{
		Name:    "Netrunner",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}

func IniAdam() *AdamSmacher {
	return &AdamSmacher{
		Name:    "AdamSmacher",
		MaxHp:   100,
		Hp:      100,
		Attaque: 25,
	}
}
