package idx

import (
	"errors"

	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/gametdb"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
	"github.com/sarpt/gamedbv/pkg/parser"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/zip"
)

// IndexConfig instructs how unziping, parsing and indexing should be performed
type IndexConfig struct {
	IndexFilepath   string
	IndexVariant    string
	Name            string
	DocType         string
	SourceFilepath  string
	ArchiveFilepath string
	SourceFilename  string
}

const bleveCreator string = "bleve"

var (
	ErrDatabaseNotOpen = errors.New("database not open")
)

// PreparePlatform unzips and parses source file, creates Index related to the platfrom and populates the database.
func (s *Server) PreparePlatform(variant platform.Variant, notifier progress.Notifier) error {
	cfg := s.cfg.Indexes[variant]
	notifier.NextStatus(newPlatformUnzipStatus(variant))
	err := zip.UnzipPlatformSource(mapToZipConfig(cfg))
	if err != nil {
		notifier.NextError(err)
		return err
	}

	notifier.NextStatus(newPlatformParsingStatus(variant))
	gametdbModelProvider, err := parsePlatformSource(mapToParser(cfg))
	if err != nil {
		notifier.NextError(err)
		return err
	}

	gametdbAdapter := NewGameTDBAdapter(variant.ID(), gametdbModelProvider)

	notifier.NextStatus(newPlatformIndexingStatus(variant))
	err = indexPlatform(mapToIndexConfig(cfg), gametdbAdapter)
	if err != nil {
		notifier.NextError(err)
		return err
	}

	if s.db == nil {
		notifier.NextError(ErrDatabaseNotOpen)
		return ErrDatabaseNotOpen
	}

	notifier.NextStatus(newDatabasePopulateStatus(variant))
	err = s.db.ProvidePlatformData(gametdbAdapter.PlatformProvider())
	if err != nil {
		notifier.NextError(err)
	}
	return err
}

func parsePlatformSource(cfg parser.Config) (gametdb.ModelProvider, error) {
	gametdbModelProvider := gametdb.ModelProvider{}
	err := parser.ParseSourceFile(cfg, &gametdbModelProvider)

	return gametdbModelProvider, err
}

func indexPlatform(platformConfig index.Config, gametdbAdapter GameTDBAdapter) error {
	creators := map[string]index.Creator{
		bleveCreator: bleve.Creator{},
	}

	return index.PrepareIndex(creators, platformConfig, gametdbAdapter.GameSources())
}

func mapToZipConfig(cfg IndexConfig) zip.Config {
	return zip.Config{
		ArchiveFilepath: cfg.ArchiveFilepath,
		SourceFilename:  cfg.SourceFilename,
		Name:            cfg.Name,
		OutputFilepath:  cfg.SourceFilepath,
	}
}

func mapToParser(cfg IndexConfig) parser.Config {
	return parser.Config{
		Filepath: cfg.SourceFilepath,
	}
}

func mapToIndexConfig(cfg IndexConfig) index.Config {
	return index.Config{
		Filepath: cfg.IndexFilepath,
		Variant:  cfg.IndexVariant,
		Name:     cfg.Name,
		DocType:  cfg.DocType,
	}
}
