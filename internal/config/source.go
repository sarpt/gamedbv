package config

// Source includes information about platform source which should provide identifying data neccessary for
// downloading, parsing, indexing etc.
type Source struct {
	name            string
	archived        bool
	archiveFilepath string
	filepath        string
	filename        string
	format          string
	forceDownload   bool
	url             string
}
