package json

// API specifies Api binary behavior
type API struct {
	IPAddress      string `json:"IPAddress"`
	Port           string `json:"Port"`
	NetInterface   string `json:"NetInterface"`
	Debug          bool   `json:"Debug"`
	DlRPCAddress   string `json:"DlRPCAddress"`
	DlRPCPort      string `json:"DlRPCPort"`
	IdxRPCAddress  string `json:"IdxRPCAddress"`
	IdxRPCPort     string `json:"IdxRPCPort"`
	ReadTimeout    string `json:"ReadTimeout"`
	RPCDialTimeout string `json:"RPCDialTimeout"`
	WriteTimeout   string `json:"WriteTimeout"`
}
