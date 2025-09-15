package player

import "strings"

type Player struct {
	Name  string
	Class string
}

func (p *Player) ChooseClass(choice string) {
	switch strings.TrimSpace(choice) {
	case "1":
		p.Class = "Corpo"
	case "2":
		p.Class = "Nomade"
	case "3":
		p.Class = "Gosse des rues"
	}
}
