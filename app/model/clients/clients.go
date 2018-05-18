package clients

type Client struct {
	ID          int
	ClientName  string
	ClientState string
	ActiveUntil string
	Equipment   string
}

type ClientListe struct {
	Clients        []Client
	Benutzername   string
	BenutzerStatus string
}

type ClientDummy struct {
	ID             int
	ClientState    string
	Benutzername   string
	BenutzerStatus string
}

func ClientListeDummy() ClientListe {
	clientliste := ClientListe{
		Benutzername:   "Peter Dieter",
		BenutzerStatus: "Verleiher",
		Clients: []Client{
			{ID: 1, ClientName: "Client 1", Equipment: "Kamera 1, Kamera 2", ClientState: "Benutzer", ActiveUntil: "12.04.2016"},
			{ID: 2, ClientName: "Client 2", Equipment: "Kamera 1, Kamera 2", ClientState: "Benutzer", ActiveUntil: "15.04.2016"},
			{ID: 3, ClientName: "Client 3", Equipment: "Kamera 1, Kamera 2", ClientState: "Benutzer", ActiveUntil: "25.04.2016"},
		},
	}
	return clientliste
}

func CreateClientDummy() ClientDummy {
	client := ClientDummy{
		Benutzername:   "Peter Dieter",
		BenutzerStatus: "Verleiher",
		ID:             5,
		ClientState:    "Benutzer",
	}
	return client
}
