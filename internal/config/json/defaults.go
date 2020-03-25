package json

import "github.com/sarpt/gamedbv/pkg/platform"

// DefaultConfigsPerPlatform are default configuration values used for platforms database files when no overwrites are present
var DefaultConfigsPerPlatform = map[string]Platform{
	platform.Wii: Platform{
		Directory: "wii",
		Name:      platform.Wii,
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
	platform.Ps3: Platform{
		Directory: "ps3",
		Name:      platform.Ps3,
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
	platform.Wiiu: Platform{
		Directory: "wiiu",
		Name:      platform.Wiiu,
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
	platform.Nds: Platform{
		Directory: "nds",
		Name:      platform.Nds,
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
	platform.N3ds: Platform{
		Directory: "3ds",
		Name:      platform.N3ds,
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
	platform.Switch: Platform{
		Directory: "switch",
		Name:      platform.Switch,
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
	Directory: "gamedbv",
	Database: Database{
		FileName: "/store.db",
		Variant:  "sqlite3",
		MaxLimit: 999,
	},
}
