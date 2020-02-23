package config

// Platform groups information used for platform database handling
type Platform struct {
	dirPath string
	name    string
	index   Index
	source  Source
}

// ArchiveFilepath returns the absolute filepath related to the platform's database archive file
func (conf Platform) ArchiveFilepath() string {
	return conf.source.archiveFilename
}

// Filepath returns the absolute filepath related to the platform's database content file
func (conf Platform) Filepath() string {
	return conf.source.filename
}

// Filename returns filename of source file containing titles database
func (conf Platform) Filename() string {
	return conf.source.filename
}

// IndexFilepath returns absolute path of Index file
func (conf Platform) IndexFilepath() string {
	return conf.index.path
}

// ForceSourceDownload spcifies if the source should be redownloaded, even in the case of source already existing in the filesystem
func (conf Platform) ForceSourceDownload() bool {
	return conf.source.forceDownload
}

// URL returns url associated with the source of the titles databases that should be fetched before parsing
func (conf Platform) URL() string {
	return conf.source.url
}

// IndexVariant returns index type (eg. bleve, solr etc.)
func (conf Platform) IndexVariant() string {
	return conf.index.variant
}

// DocType returns document identifier used for indexes documents matching
func (conf Platform) DocType() string {
	return conf.index.docType
}

// Name returns name of the platform whose information is presented in config
func (conf Platform) Name() string {
	return conf.source.name
}

// DirectoryPath returns directory name related to the specified platform
func (conf Platform) DirectoryPath() string {
	return conf.dirPath
}
