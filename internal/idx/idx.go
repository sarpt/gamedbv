package idx

import (
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/gametdb"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
	"github.com/sarpt/gamedbv/pkg/parser"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/progress"
	"github.com/sarpt/gamedbv/pkg/zip"
)

const bleveCreator string = "bleve"

// Config instructs how unziping, parsing and indexing should be performed
type Config struct {
	IndexFilepath   string
	IndexVariant    string
	Name            string
	DocType         string
	SourceFilepath  string
	ArchiveFilepath string
	SourceFilename  string
}

// PreparePlatform unzips and parses source file, creates Index related to the platfrom and populates the database
func PreparePlatform(cfg Config, platformVariant platform.Variant, printer progress.Notifier, database db.Database) {
	platformName := platformVariant.String()

	printer.NextStatus(newPlatformUnzipStatus(platformName))
	err := zip.UnzipPlatformSource(getZipConfig(cfg))
	if err != nil {
		printer.NextError(err)
		return
	}

	printer.NextStatus(newPlatformParsingStatus(platformName))
	gametdbModelProvider, err := parsePlatformSource(getParserConfig(cfg))
	if err != nil {
		printer.NextError(err)
		return
	}

	gametdbAdapter := NewGameTDBAdapter(platformVariant.ID(), gametdbModelProvider)

	printer.NextStatus(newPlatformIndexingStatus(platformName))
	err = indexPlatform(getIndexConfig(cfg), gametdbAdapter)
	if err != nil {
		printer.NextError(err)
		return
	}

	printer.NextStatus(newDatabasePopulateStatus(platformName))
	err = database.ProvidePlatformData(gametdbAdapter.PlatformProvider())
	if err != nil {
		printer.NextError(err)
	}
}

func parsePlatformSource(cfg parser.Config) (gametdb.ModelProvider, error) {
	gametdbModelProvider := gametdb.ModelProvider{}
	err := parser.ParseSourceFile(cfg, &gametdbModelProvider)

	return gametdbModelProvider, err
}

func indexPlatform(platformConfig index.Config, gametdbAdapter GameTDBAdapter) error {
	creators := map[string]index.Creator{
		bleveCreator: bleve.BleveCreator{},
	}

	return index.PrepareIndex(creators, platformConfig, gametdbAdapter.GameSources())
}

func getZipConfig(cfg Config) zip.Config {
	return zip.Config{
		ArchiveFilepath: cfg.ArchiveFilepath,
		SourceFilename:  cfg.SourceFilename,
		Name:            cfg.Name,
		OutputFilepath:  cfg.SourceFilepath,
	}
}

func getParserConfig(cfg Config) parser.Config {
	return parser.Config{
		Filepath: cfg.SourceFilepath,
	}
}

func getIndexConfig(cfg Config) index.Config {
	return index.Config{
		Filepath: cfg.IndexFilepath,
		Variant:  cfg.IndexVariant,
		Name:     cfg.Name,
		DocType:  cfg.DocType,
	}
}
