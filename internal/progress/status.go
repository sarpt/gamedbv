package progress

// Status provides information about step and message of the process in question
type Status struct {
	Platform string `json:"platform,omitempty"`
	Process  string `json:"process,omitempty"`
	Step     string `json:"step,omitempty"`
	Message  string `json:"message,omitempty"`
}
