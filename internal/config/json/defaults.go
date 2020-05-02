package json

import "github.com/sarpt/gamedbv/pkg/platform"

// DefaultConfigsPerPlatform are default configuration values used for platforms database files when no overwrites are present
var DefaultConfigsPerPlatform = map[string]Platform{
	platform.Wii.ID(): {
		Directory: "wii",
		Name:      platform.Wii.ID(),
		Index: Index{
			Directory: "wii_bleve",
			DocType:   "gametdb/game",
			Variant:   "bleve",
		},
		Source: Source{
			Archived:        true,
			ArchiveFilename: "wiidb.zip",
			Filename:        "wiitdb.xml",
			ForceDownload:   false,
			Format:          "xml",
			Name:            "GameTDB",
			URL:             "https://www.gametdb.com/wiitdb.zip",
		},
	},
	platform.Ps3.ID(): {
		Directory: "ps3",
		Name:      platform.Ps3.ID(),
		Index: Index{
			Directory: "ps3_bleve",
			DocType:   "gametdb/game",
			Variant:   "bleve",
		},
		Source: Source{
			Name:            "GameTDB",
			URL:             "https://www.gametdb.com/ps3tdb.zip",
			Filename:        "ps3tdb.xml",
			ForceDownload:   false,
			Format:          "xml",
			Archived:        true,
			ArchiveFilename: "ps3db.zip",
		},
	},
	platform.Wiiu.ID(): {
		Directory: "wiiu",
		Name:      platform.Wiiu.ID(),
		Index: Index{
			Directory: "wiiu_bleve",
			DocType:   "gametdb/game",
			Variant:   "bleve",
		},
		Source: Source{
			Archived:        true,
			ArchiveFilename: "wiiu.zip",
			Filename:        "wiiutdb.xml",
			ForceDownload:   false,
			Format:          "xml",
			Name:            "GameTDB",
			URL:             "https://www.gametdb.com/wiiutdb.zip",
		},
	},
	platform.Nds.ID(): {
		Directory: "nds",
		Name:      platform.Nds.ID(),
		Index: Index{
			Directory: "nds_bleve",
			DocType:   "gametdb/game",
			Variant:   "bleve",
		},
		Source: Source{
			Archived:        true,
			ArchiveFilename: "nds.zip",
			Filename:        "dstdb.xml",
			ForceDownload:   false,
			Format:          "xml",
			Name:            "GameTDB",
			URL:             "https://www.gametdb.com/dstdb.zip",
		},
	},
	platform.N3ds.ID(): {
		Directory: "3ds",
		Name:      platform.N3ds.ID(),
		Index: Index{
			Directory: "3ds_bleve",
			DocType:   "gametdb/game",
			Variant:   "bleve",
		},
		Source: Source{
			Archived:        true,
			ArchiveFilename: "3ds.zip",
			Filename:        "3dstdb.xml",
			ForceDownload:   false,
			Format:          "xml",
			Name:            "GameTDB",
			URL:             "https://www.gametdb.com/3dstdb.zip",
		},
	},
	platform.Switch.ID(): {
		Directory: "switch",
		Name:      platform.Switch.ID(),
		Index: Index{
			Directory: "switch_bleve",
			DocType:   "gametdb/game",
			Variant:   "bleve",
		},
		Source: Source{
			ArchiveFilename: "switch.zip",
			Filename:        "switchtdb.xml",
			Name:            "GameTDB",
			Archived:        true,
			ForceDownload:   false,
			Format:          "xml",
			URL:             "https://www.gametdb.com/switchtdb.zip",
		},
	},
}

// DefaultConfig is used when no overrides are present
var DefaultConfig App = App{
	API: API{
		Address:      "127.0.0.1:3001",
		Debug:        false,
		WriteTimeout: "15s",
		ReadTimeout:  "15s",
	},
	Directory: "gamedbv",
	Database: Database{
		FileName: "/store.db",
		Variant:  "sqlite3",
		MaxLimit: 999,
	},
}
