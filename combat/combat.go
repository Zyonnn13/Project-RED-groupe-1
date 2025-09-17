package combat

import (
	"Project-RED-groupe-1/player"
	"fmt"
	"math/rand"
	"time"
)

func LancerCombat(joueur *player.Designplayer, ennemi Ennemis) {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Un combat commence contre %s !\n", ennemi.Name)

	for joueur.HP > 0 && ennemi.Hp > 0 {
		// Tour du joueur
		damage := calcDamage(joueur.Attack)
		ennemi.Hp -= damage
		if ennemi.Hp < 0 {
			ennemi.Hp = 0
		}
		fmt.Printf("%s attaque et inflige %d dégâts. HP ennemi : %d/%d\n",
			joueur.Name, damage, ennemi.Hp, ennemi.MaxHp)

		if ennemi.Hp <= 0 {
			fmt.Printf("%s est vaincu !\n", ennemi.Name)
			break
		}

		// Tour de l'ennemi
		damage = calcDamage(ennemi.Attaque)
		joueur.HP -= damage
		if joueur.HP < 0 {
			joueur.HP = 0
		}
		fmt.Printf("%s attaque et inflige %d dégâts. HP joueur : %d/%d\n",
			ennemi.Name, damage, joueur.HP, joueur.MaxHP)

		if joueur.HP <= 0 {
			fmt.Printf("Vous avez été vaincu par %s...\n", ennemi.Name)
			break
		}
	}
}

func calcDamage(base int) int {
	variance := base / 10
	if variance < 1 {
		variance = 1
	}
	return base + rand.Intn(variance*2+1) - variance
}
