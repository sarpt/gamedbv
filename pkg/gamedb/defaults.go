package gamedb

import "github.com/sarpt/gamedbv/pkg/platform"

// DefaultDbInfosByPlatform are default values used for database fetching when no overwrites present
var DefaultDbInfosByPlatform = map[string]Info{
	platform.Wii: Info{
		ArchiveFileName:   "wiidb.zip",
		ContentFileName:   "wiitdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/wiitdb.zip",
		ForceDbDownload:   false,
		PlatformDirectory: "wii",
	},
	platform.Ps3: Info{
		ArchiveFileName:   "ps3db.zip",
		ContentFileName:   "ps3tdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/ps3tdb.zip",
		ForceDbDownload:   false,
		PlatformDirectory: "ps3",
	},
	platform.Wiiu: Info{
		ArchiveFileName:   "wiiu.zip",
		ContentFileName:   "wiiutdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/wiiutdb.zip",
		ForceDbDownload:   false,
		PlatformDirectory: "wiiu",
	},
	platform.Nds: Info{
		ArchiveFileName:   "nds.zip",
		ContentFileName:   "dstdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/dstdb.zip",
		ForceDbDownload:   false,
		PlatformDirectory: "nds",
	},
	platform.N3ds: Info{
		ArchiveFileName:   "3ds.zip",
		ContentFileName:   "3dstdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/3dstdb.zip",
		ForceDbDownload:   false,
		PlatformDirectory: "3ds",
	},
	platform.Switch: Info{
		ArchiveFileName:   "switch.zip",
		ContentFileName:   "switchtdb.xml",
		DatabaseType:      "GameTDB",
		URL:               "https://www.gametdb.com/switchtdb.zip",
		ForceDbDownload:   false,
		PlatformDirectory: "switch",
	},
}
