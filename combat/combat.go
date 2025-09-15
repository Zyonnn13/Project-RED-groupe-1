package combat

import (
	"Project-RED-groupe-1/player"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Monster struct {
	Name      string
	MaxHp     int
	Hp        int
	Attaque   int
	CurrentHP int
}

func IniAgentCorpo() *Monster {
	return &Monster{
		Name:    "Gobelin",
		MaxHp:   40,
		Hp:      40,
		Attaque: 5,
	}

}

func GoblinPattern(player *player.Designplayer) {
	goblin := IniAgentCorpo()
	turn := 1

	fmt.Println("\n Début du combat d'entraînement contre le Gobelin !")

	for player.HP > 0 {
		var damage int
		if turn%3 == 0 {
			damage = goblin.Attaque * 2
		} else {
			damage = goblin.Attaque
		}

		player.HP -= damage
		if player.HP < 0 {
			player.HP = 0
		}

		fmt.Printf("Tour %d : %s inflige à %s %d de dégâts.\n", turn, goblin.Name, player.Name, damage)
		fmt.Printf("PV de %s : %d/%d\n", player.Name, player.HP, player.MaxHP)

		if player.HP == 0 {
			fmt.Printf("\n %s est KO. Fin du combat.\n", player.Name)
			break
		}
		turn++
	}
}

func CharacterTurn(p *player.Designplayer, m *Monster) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n🎮 Ton tour de jeu !")
	fmt.Println("1 - Attaquer")
	fmt.Println("2 - Inventaire")

	fmt.Print("Choix : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		// Attaque basique
		damage := 5
		m.CurrentHP -= damage
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}

		fmt.Printf("\n %s utilise Attaque basique et inflige %d dégâts à %s !\n", p.Name, damage, m.Name)
		fmt.Printf("PV de %s : %d/%d\n", m.Name, m.CurrentHP, m.MaxHp)

		if m.CurrentHP == 0 {
			fmt.Printf("✅ %s est vaincu !\n", m.Name)
			return
		}

		// Tour du monstre
		fmt.Println("\n Tour du monstre !")
		p.HP -= m.Attaque
		if p.HP < 0 {
			p.HP = 0
		}
		fmt.Printf("%s inflige %d dégâts à %s.\n", m.Name, m.Attaque, p.Name)
		fmt.Printf("PV de %s : %d/%d\n", p.Name, p.HP, p.MaxHP)

		if p.HP == 0 {
			fmt.Printf(" %s est KO...\n", p.Name)
		}

	case "2":
		fmt.Println("\n Inventaire : (fonctionnalité à venir...)")
	default:
		fmt.Println(" Choix invalide.")
	}
}
