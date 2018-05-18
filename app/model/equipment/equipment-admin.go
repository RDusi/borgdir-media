package equipment

type EquipmentAdmin struct {
	ID          int
	Name        string
	Inhalt      string
	Anzahl      int
	Entliehen   string
	Rueckgabe   string
	EntliehenAn string
}

type EquipmentAdminData struct {
	Items          []EquipmentAdmin
	Benutzername   string
	BenutzerStatus string
}

func EquipmentAdminListeDummy() EquipmentAdminData {
	liste := EquipmentAdminData{
		Benutzername:   "Peter Dieter",
		BenutzerStatus: "Verleiher",
		Items: []EquipmentAdmin{
			{ID: 1, Name: "Kamera 1", Inhalt: "Beschreibung", Rueckgabe: "21.04.2015", EntliehenAn: "Nutzer 1"},
			{ID: 2, Name: "Kamera 2", Inhalt: "Beschreibung", Rueckgabe: "16.04.2015", EntliehenAn: "Nutzer 2"},
			{ID: 3, Name: "Kamera 3", Inhalt: "Beschreibung", Rueckgabe: "07.04.2015", EntliehenAn: "Nutzer 3"},
		},
	}
	return liste
}