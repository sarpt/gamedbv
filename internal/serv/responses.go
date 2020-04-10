package serv

type description struct {
	Language language `json:"language"`
	Title    string   `json:"title"`
	Synopsis string   `json:"synopsis"`
}

type game struct {
	UID          string        `json:"uid"`
	SerialNumber string        `json:"serialNumber"`
	Region       string        `json:"region"`
	Platform     string        `json:"platform"`
	Descriptions []description `json:"descriptions"`
}

type language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type platform struct {
	Name string `json:"name"`
}

type gamesResponse struct {
	Total int    `json:"total"`
	Games []game `json:"games"`
}

type languagesResponse struct {
	Languages []language `json:"languages"`
}

type platformsResponse struct {
	Platforms []platform `json:"platforms"`
}
