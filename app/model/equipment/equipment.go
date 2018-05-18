package equipment

type Equipment struct {
	ID     int
	Name   string
	Inhalt string
	Anzahl int
	Status int
}

type EquipmentData struct {
	Items          []Equipment
	Benutzername   string
	BenutzerStatus string
}

func EuqipmentListeDummy() EquipmentData {
	liste := EquipmentData{
		Benutzername:   "Erica Mustermann",
		BenutzerStatus: "Benutzer",
		Items: []Equipment{
			{ID: 1, Name: "Kamera 1", Inhalt: "Beschreibung", Anzahl: 123, Status: 2},
			{ID: 2, Name: "Kamera 2", Inhalt: "Beschreibung", Anzahl: 10, Status: 2},
			{ID: 3, Name: "Kamera 3", Inhalt: "Beschreibung", Anzahl: 200, Status: 2},
		},
	}
	return liste
}
