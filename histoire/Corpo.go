package histoire

import (
	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/player"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CorpoHistoire() {
	fmt.Println("\n=== Histoire Corpo ===")
	fmt.Println("Les néons de Night City se reflètent sur les vitres teintées de la tour Arasaka. Dans ton bureau, un message crypté s’affiche : « V, on a un problème. »")
	fmt.Println(" Ton supérieur, Jenkins, t’explique qu’il faut discréditer un rival politique. Tu sais que ce genre de mission ne se termine jamais proprement.")
}

func StartCorpo(p *player.Designplayer) {
	reader := bufio.NewReader(os.Stdin)
	inv := inventaire.NewInventory() //inventaire joueur

	fmt.Println(" 1- Accepter sans poser de questions ")
	fmt.Println(" 2- [Corpo] Utiliser ton réseau pour saboter ton rival avant même la mission ")
	fmt.Println(" 3- Refuser ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

}
