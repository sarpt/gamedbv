package platform

import "github.com/sarpt/gamedbv/pkg/gamedbv"

// DefaultConfigsPerPlatform are default configuration values used for platforms database files when no overwrites are present
var DefaultConfigsPerPlatform = map[string]Config{
	Wii: Config{
		AppConfig:         gamedbv.DefaultConfig,
		ArchiveFileName:   "wiidb.zip",
		ContentFileName:   "wiitdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/wiitdb.zip",
		ForceDbDownload:   false,
		IndexDir:          "wii_bleve",
		PlatformDirectory: "wii",
	},
	Ps3: Config{
		AppConfig:         gamedbv.DefaultConfig,
		ArchiveFileName:   "ps3db.zip",
		ContentFileName:   "ps3tdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/ps3tdb.zip",
		ForceDbDownload:   false,
		IndexDir:          "ps3_bleve",
		PlatformDirectory: "ps3",
	},
	Wiiu: Config{
		AppConfig:         gamedbv.DefaultConfig,
		ArchiveFileName:   "wiiu.zip",
		ContentFileName:   "wiiutdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/wiiutdb.zip",
		ForceDbDownload:   false,
		IndexDir:          "wiiu_bleve",
		PlatformDirectory: "wiiu",
	},
	Nds: Config{
		AppConfig:         gamedbv.DefaultConfig,
		ArchiveFileName:   "nds.zip",
		ContentFileName:   "dstdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/dstdb.zip",
		ForceDbDownload:   false,
		IndexDir:          "nds_bleve",
		PlatformDirectory: "nds",
	},
	N3ds: Config{
		AppConfig:         gamedbv.DefaultConfig,
		ArchiveFileName:   "3ds.zip",
		ContentFileName:   "3dstdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/3dstdb.zip",
		ForceDbDownload:   false,
		IndexDir:          "3ds_bleve",
		PlatformDirectory: "3ds",
	},
	Switch: Config{
		AppConfig:         gamedbv.DefaultConfig,
		ArchiveFileName:   "switch.zip",
		ContentFileName:   "switchtdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/switchtdb.zip",
		ForceDbDownload:   false,
		IndexDir:          "switch_bleve",
		PlatformDirectory: "switch",
	},
}
