package entities

type Benutzer struct {
	ID          int
	Username    string
	Userstate   string
	Usertype    string
	Email       string
	ActiveUntil string
	Picture     string
	Password    string
}

type ListBenutzer struct {
	Benutzer []Benutzer
}
