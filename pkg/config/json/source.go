package json

// Source includes information about platform source which should provide identifying data neccessary for
// downloading, parsing, indexing etc.
type Source struct {
	Name            string
	Archived        bool
	ArchiveFilename string
	Filename        string
	Format          string
	ForceDownload   bool
	URL             string
}
