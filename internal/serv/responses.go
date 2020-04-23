package serv

type descriptionResponse struct {
	Language languageResponse `json:"language"`
	Title    string           `json:"title"`
	Synopsis string           `json:"synopsis"`
}

type regionResponse struct {
	Code string `json:"code"`
}

type gameResponse struct {
	UID          string                `json:"uid"`
	SerialNumber string                `json:"serialNumber"`
	Region       regionResponse        `json:"region"`
	Platform     platformResponse      `json:"platform"`
	Descriptions []descriptionResponse `json:"descriptions"`
}

type languageResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type platformResponse struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

type gamesResponse struct {
	Total int            `json:"total"`
	Games []gameResponse `json:"games"`
}

type languagesResponse struct {
	Languages []languageResponse `json:"languages"`
}

type platformsResponse struct {
	Platforms []platformResponse `json:"platforms"`
}

type regionsResponse struct {
	Regions []regionResponse `json:"regions"`
}
