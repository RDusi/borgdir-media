package cart

import "github.com/jhoefker/borgdir-media/app/model/equipment"

type CartItem struct {
	ID        int
	Equipment equipment.Equipment
	Anzahl    int
	Rueckgabe string
}
