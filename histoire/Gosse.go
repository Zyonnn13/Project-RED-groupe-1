package histoire

import (
	"Project-RED-groupe-1/player"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GosseHistoire() {
	fmt.Println("\n========= Histoire Gosse des rues ========")
	fmt.Println(" Tu as grandi dans les ruelles de Heywood. Tu connais les gangs, les deals, et comment survivre. Mais aujourd’hui, un vieil ami te demande un service dangereux.")
	fmt.Println("La pluie tombe sur les trottoirs crasseux. Dans un bar enfumé, ton ami Pepe t’explique qu’il doit de l’argent à un gang : les Valentinos. Il te supplie d’aller négocier avec leur chef, Kirk.")
}

func StartGosse(p *player.Designplayer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(" 1- Accepter de négocier")
	fmt.Println(" 2- Proposer de rembourser toi-même ")
	fmt.Println(" 3- Tendre un piège à Kirk ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("Tu rencontres Kirk dans une arrière-salle du bar. Il t’écoute, mais les Valentinos ne sont pas réputés pour leur patience.")
	}
}
