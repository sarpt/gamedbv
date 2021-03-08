package json

// API specifies Api binary behavior
type API struct {
	IPAddress    string `json:"IPAddress"`
	Port         string `json:"Port"`
	NetInterface string `json:"NetInterface"`
	Debug        bool   `json:"Debug"`
	ReadTimeout  string `json:"ReadTimeout"`
	WriteTimeout string `json:"WriteTimeout"`
}
