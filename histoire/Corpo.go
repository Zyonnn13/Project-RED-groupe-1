package histoire

import (
	"Project-RED-groupe-1/inventaire"
	"bufio"
	"fmt"
	"os"
)

func CorpoHistoire() {
	fmt.Println("\n=== Histoire Corpo ===")
	fmt.Println("Les néons de Night City se reflètent sur les vitres teintées de la tour Arasaka. Dans ton bureau, un message crypté s’affiche : « V, on a un problème. »")
	fmt.Println(" Ton supérieur, Jenkins, t’explique qu’il faut discréditer un rival politique. Tu sais que ce genre de mission ne se termine jamais proprement.")
}

func StartScenario() {
	reader := bufio.NewReader(os.Stdin)
	inv := inventaire.NewInventory() //inventaire joueur

	fmt.Println("")

}
