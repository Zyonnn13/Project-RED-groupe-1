package monnaie

// Eddies représente la monnaie virtuelle du joueur
type Eddies struct {
	Balance int
}

// NewEddies initialise la monnaie avec un solde de départ
func NewEddies(initialBalance int) *Eddies {
	return &Eddies{Balance: initialBalance}
}

// Add ajoute des crédits
func (e *Eddies) Add(amount int) {
	e.Balance += amount
}

// Spend tente de dépenser des crédits
func (e *Eddies) Spend(amount int) bool {
	if e.Balance >= amount {
		e.Balance -= amount
		return true
	}
	return false
}

// GetBalance retourne le solde actuel
func (e *Eddies) GetBalance() int {
	return e.Balance
}
