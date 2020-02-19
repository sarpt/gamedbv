package config

import "github.com/sarpt/gamedbv/pkg/platform"

// DefaultConfig is used when no overrides are present
var DefaultConfig App = App{
	BaseDirectory: "gamedbv",
}

// DefaultConfigsPerPlatform are default configuration values used for platforms database files when no overwrites are present
var DefaultConfigsPerPlatform = map[string]Platform{
	platform.Wii: Platform{
		appConfig: DefaultConfig,
		directory: "wii",
		name:      platform.Wii,
		index: Index{
			directory: "wii_bleve",
			docType:   "gametdb/game",
			variant:   "bleve",
		},
		source: Source{
			archived:        true,
			archiveFilename: "wiidb.zip",
			filename:        "wiitdb.xml",
			forceDownload:   false,
			format:          "xml",
			name:            "GameTDB",
			url:             "https://www.gametdb.com/wiitdb.zip",
		},
		database: Database{
			path:    "/store.db",
			variant: "sqlite3",
		},
	},
	platform.Ps3: Platform{
		appConfig: DefaultConfig,
		directory: "ps3",
		name:      platform.Ps3,
		index: Index{
			directory: "ps3_bleve",
			docType:   "gametdb/game",
			variant:   "bleve",
		},
		source: Source{
			name:            "GameTDB",
			url:             "https://www.gametdb.com/ps3tdb.zip",
			filename:        "ps3tdb.xml",
			forceDownload:   false,
			format:          "xml",
			archived:        true,
			archiveFilename: "ps3db.zip",
		},
		database: Database{
			path:    "/store.db",
			variant: "sqlite3",
		},
	},
	platform.Wiiu: Platform{
		appConfig: DefaultConfig,
		directory: "wiiu",
		name:      platform.Wiiu,
		index: Index{
			directory: "wiiu_bleve",
			docType:   "gametdb/game",
			variant:   "bleve",
		},
		source: Source{
			archived:        true,
			archiveFilename: "wiiu.zip",
			filename:        "wiiutdb.xml",
			forceDownload:   false,
			format:          "xml",
			name:            "GameTDB",
			url:             "https://www.gametdb.com/wiiutdb.zip",
		},
		database: Database{
			path:    "/store.db",
			variant: "sqlite3",
		},
	},
	platform.Nds: Platform{
		appConfig: DefaultConfig,
		directory: "nds",
		name:      platform.Nds,
		index: Index{
			directory: "nds_bleve",
			docType:   "gametdb/game",
			variant:   "bleve",
		},
		source: Source{
			archived:        true,
			archiveFilename: "nds.zip",
			filename:        "dstdb.xml",
			forceDownload:   false,
			format:          "xml",
			name:            "GameTDB",
			url:             "https://www.gametdb.com/dstdb.zip",
		},
		database: Database{
			path:    "/store.db",
			variant: "sqlite3",
		},
	},
	platform.N3ds: Platform{
		appConfig: DefaultConfig,
		directory: "3ds",
		name:      platform.N3ds,
		index: Index{
			directory: "3ds_bleve",
			docType:   "gametdb/game",
			variant:   "bleve",
		},
		source: Source{
			archived:        true,
			archiveFilename: "3ds.zip",
			filename:        "3dstdb.xml",
			forceDownload:   false,
			format:          "xml",
			name:            "GameTDB",
			url:             "https://www.gametdb.com/3dstdb.zip",
		},
		database: Database{
			path:    "/store.db",
			variant: "sqlite3",
		},
	},
	platform.Switch: Platform{
		appConfig: DefaultConfig,
		directory: "switch",
		name:      platform.Switch,
		index: Index{
			directory: "switch_bleve",
			docType:   "gametdb/game",
			variant:   "bleve",
		},
		source: Source{
			archiveFilename: "switch.zip",
			filename:        "switchtdb.xml",
			name:            "GameTDB",
			archived:        true,
			forceDownload:   false,
			format:          "xml",
			url:             "https://www.gametdb.com/switchtdb.zip",
		},
		database: Database{
			path:    "/store.db",
			variant: "sqlite3",
		},
	},
}
