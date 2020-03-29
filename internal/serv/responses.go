package serv

type description struct {
	Language string `json:"language"`
	Title    string `json:"title"`
	Synopsis string `json:"synopsis"`
}

type game struct {
	UID          string        `json:"uid"`
	SerialNumber string        `json:"serialNumber"`
	Region       string        `json:"region"`
	Platform     string        `json:"platform"`
	Descriptions []description `json:"descriptions"`
}

type gamesResponse struct {
	Total int    `json:"total"`
	Games []game `json:"games"`
}
