package profil

type Profil struct {
	ID             int
	Benutzername   string
	Email          string
	Passwort       string
	BenutzerStatus string
}

func (u *Profil) SetCurrentUser(ID int, Benutzername string, Email string, Passwort string, BenutzerStatus string) {
	u.ID = ID
	u.Benutzername = Benutzername
	u.Email = Email
	u.Passwort = Passwort
	u.BenutzerStatus = BenutzerStatus
}

func (u *Profil) GetCurrentUser() Profil {
	benutzer := Profil{
		ID:             u.ID,
		Benutzername:   u.Benutzername,
		Email:          u.Email,
		Passwort:       u.Passwort,
		BenutzerStatus: u.BenutzerStatus,
	}
	return benutzer
}

func ProfilDummyDataUser() Profil {
	liste := Profil{Benutzername: "Erica Mustermann", BenutzerStatus: "Benutzer"}
	return liste
}

func ProfilDummyDataAdmin() Profil {
	liste := Profil{Benutzername: "Peter Dieter", BenutzerStatus: "Verleiher"}
	return liste
}
