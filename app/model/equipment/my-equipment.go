package equipment

type MyEquipment struct {
	ID        int
	Name      string
	Inhalt    string
	Anzahl    int
	Entliehen string
	Rueckgabe string
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
			{ID: 1, Name: "Kamera 1", Inhalt: "Beschreibung", Entliehen: "20.04.2015", Rueckgabe: "21.04.2015"},
			{ID: 2, Name: "Kamera 2", Inhalt: "Beschreibung", Entliehen: "15.04.2015", Rueckgabe: "16.04.2015"},
			{ID: 3, Name: "Kamera 3", Inhalt: "Beschreibung", Entliehen: "06.04.2015", Rueckgabe: "07.04.2015"},
		},
	}
	return liste
}
