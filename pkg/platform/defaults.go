package platform

import "github.com/sarpt/gamedbv/pkg/gamedbv"

// DefaultConfigsPerPlatform are default configuration values used for platforms database files when no overwrites are present
var DefaultConfigsPerPlatform = map[string]Config{
	Wii: Config{
		appConfig: gamedbv.DefaultConfig,
		directory: "wii",
		name:      Wii,
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
	},
	Ps3: Config{
		appConfig: gamedbv.DefaultConfig,
		directory: "ps3",
		name:      Ps3,
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
	},
	Wiiu: Config{
		appConfig: gamedbv.DefaultConfig,
		directory: "wiiu",
		name:      Wiiu,
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
	},
	Nds: Config{
		appConfig: gamedbv.DefaultConfig,
		directory: "nds",
		name:      Nds,
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
	},
	N3ds: Config{
		appConfig: gamedbv.DefaultConfig,
		directory: "3ds",
		name:      N3ds,
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
	},
	Switch: Config{
		appConfig: gamedbv.DefaultConfig,
		directory: "switch",
		name:      Switch,
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
	},
}
