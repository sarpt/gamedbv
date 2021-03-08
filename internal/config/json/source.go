package json

// Source includes information about platform source which should provide identifying data neccessary for
// downloading, parsing, indexing etc.
type Source struct {
	Name            string `json:"Name"`
	Archived        bool   `json:"Archived"`
	ArchiveFilename string `json:"ArchiveFilename"`
	Filename        string `json:"Filename"`
	Format          string `json:"Format"`
	ForceDownload   bool   `json:"ForceDownload"`
	URL             string `json:"URL"`
}
