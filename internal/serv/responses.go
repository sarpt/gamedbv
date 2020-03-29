package serv

type description struct {
	Language    string
	Title       string
	Description string
}

type game struct {
	UID          string
	SerialNumber string
	Region       string
	Platform     string
	Descriptions []description
}

type gamesResponse struct {
	Total int
	Games []game
}
