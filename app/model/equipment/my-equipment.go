package equipment

type MyEquipment struct {
	EuqipmentItem Equipment
	Entliehen     string
	Rueckgabe     string
}

type MyEquipmentData struct {
	Items          []MyEquipment
	Benutzername   string
	BenutzerStatus string
}

func MyEquipmentListeDummy() MyEquipmentData {
	liste := MyEquipmentData{
		Benutzername:   "Erica Mustermann",
		BenutzerStatus: "Benutzer",
		Items: []MyEquipment{
			{EuqipmentItem: EuqipmentListeDummy().Items[0], Entliehen: "20.04.2015", Rueckgabe: "21.04.2015"},
			{EuqipmentItem: EuqipmentListeDummy().Items[1], Entliehen: "20.04.2015", Rueckgabe: "21.04.2015"},
			{EuqipmentItem: EuqipmentListeDummy().Items[2], Entliehen: "20.04.2015", Rueckgabe: "21.04.2015"},
			{EuqipmentItem: EuqipmentListeDummy().Items[3], Entliehen: "20.04.2015", Rueckgabe: "21.04.2015"},
		},
	}
	return liste
}
