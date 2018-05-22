package equipment

type Equipment struct {
	ID     int
	Name   string
	Inhalt string
	Anzahl int
	Status int
	Bild   string
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
			{ID: 1, Name: "Kamera 1", Inhalt: "Beschreibung", Anzahl: 123, Status: 2, Bild: "../../../static/images/kamera1_150x150.jpg"},
			{ID: 2, Name: "Stativ 1", Inhalt: "Beschreibung", Anzahl: 10, Status: 2, Bild: "../../../static/images/stativ1_150x150.jpg"},
			{ID: 3, Name: "Mikro 1", Inhalt: "Beschreibung", Anzahl: 200, Status: 2, Bild: "../../../static/images/mikro1_150x150.jpg"},
			{ID: 4, Name: "Objektiv 1", Inhalt: "Beschreibung", Anzahl: 200, Status: 2, Bild: "../../../static/images/objektiv1_150x150.jpg"},
		},
	}
	return liste
}
