package hack

import (
	"Project-RED-groupe-1/combat/Ennemis"
)

func Surchauffe() {
	Surchauffeturn := 3
	if Surchauffeturn > 0 {
		Ennemis.Hp -= 10
		Surchauffeturn -= 1
	}
}
