package histoire

import (
	"Project-RED-groupe-1/combat"
	"Project-RED-groupe-1/inventaire"
	"Project-RED-groupe-1/player"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Acte1_Relic(p *player.Player, inv *inventaire.Inventory) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- ACTE 1 : LA RELIC ---")
	fmt.Println("Après ton prologue, tu retrouves Jackie Welles.")
	fmt.Println("Il te propose un gros coup : travailler pour Dexter DeShawn, un fixer influent.")
	fmt.Println("Dex veut que vous voliez une puce prototype : la Relic, au Konpeki Plaza.")

	fmt.Println("\nChoix :")
	fmt.Println("1 - Accepter directement")
	fmt.Println("2 - Négocier ta part avec Dex")
	fmt.Println("3 - Refuser la mission")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("\nTu acceptes sans discuter. Jackie devient ton partenaire.")
	case "2":
		fmt.Println("\nTu tentes de négocier... Dex accepte à contre-cœur. Tu gagnes +500 eddies mais il se méfie de toi.")
		p.Eddies.Add(500)
	case "3":
		fmt.Println("\nJackie insiste, mais tu perds un peu de sa confiance...")
	}

	// --- Préparation avec Evelyn Parker ---
	fmt.Println("\nDex te présente Evelyn Parker. Elle t’explique que la Relic est dans la suite de Yorinobu Arasaka.")
	fmt.Println("Tu dois analyser la suite via une braindance...")

	fmt.Println("\nChoix :")
	fmt.Println("1 - Se concentrer sur la sécurité (caméras, gardes)")
	fmt.Println("2 - Chercher des accès cachés (conduits d’aération)")
	fmt.Println("3 - Étudier la Relic (stabilité, température)")
	bdChoice, _ := reader.ReadString('\n')
	bdChoice = strings.TrimSpace(bdChoice)

	switch bdChoice {
	case "1":
		fmt.Println("\nTu découvres plusieurs caméras cachées. Ça pourra t’aider plus tard.")
	case "2":
		fmt.Println("\nTu trouves un conduit d’aération qui mène directement à la suite.")
	case "3":
		fmt.Println("\nTu apprends que la Relic doit rester à température contrôlée... ça peut poser problème.")
	}

	// --- Braquage Konpeki ---
	fmt.Println("\nToi et Jackie infiltrez le Konpeki Plaza...")
	fmt.Println("Tout semble sous contrôle, jusqu’à ce que Yorinobu Arasaka assassine son père sous vos yeux !")
	fmt.Println("L’alarme retentit, les gardes rappliquent !")

	fmt.Println("\nUn combat commence contre les agents d’Arasaka !")
	combat.LancerCombat(p, combat.Agentcorpo, inv) // à définir dans combat.go

	// --- Jackie blessé ---
	fmt.Println("\nVous réussissez à vous enfuir, mais Jackie est gravement blessé.")
	fmt.Println("Dans ses derniers instants, il te confie la Relic...")
	fmt.Println("Tu l’insères dans ton crâne... et perds connaissance.")

	// --- Apparition de Johnny ---
	fmt.Println("\nQuand tu te réveilles... Johnny Silverhand est dans ta tête !")
	fmt.Println("Rockeur-terroriste légendaire, mort depuis 50 ans... et maintenant coincé avec toi.")

	fmt.Println("\nChoix :")
	fmt.Println("1 - L’ignorer (Johnny devient hostile)")
	fmt.Println("2 - L’écouter (il te révèle des infos sur Arasaka)")
	fmt.Println("3 - Le provoquer (relation conflictuelle)")
	johnnyChoice, _ := reader.ReadString('\n')
	johnnyChoice = strings.TrimSpace(johnnyChoice)

	switch johnnyChoice {
	case "1":
		fmt.Println("\nJohnny t’insulte et promet de te pourrir la vie.")
	case "2":
		fmt.Println("\nJohnny t’explique qu’Arasaka manipule tout. Tu gagnes des infos précieuses.")
	case "3":
		fmt.Println("\nJohnny se vexe... mais il respecte ton caractère.")
	}

	fmt.Println("\n--- Fin de l’Acte 1 : La Relic ---")
	fmt.Println("Tu es désormais lié à Johnny. La Relic te tue lentement, et il faut trouver un moyen de l’extraire...")
	Acte2_Alliances(p, inv)
}
