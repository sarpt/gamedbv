package dbdownload

// DefaultDbInfosByPlatform are default values used for database fetching when no overwrites present
var DefaultDbInfosByPlatform = map[string]DbInfo{
	"wii": DbInfo{
		DbArchiveFileName: "wiidb.zip",
		DbRemoteURL:       "https://www.gametdb.com/wiitdb.zip",
		ForceDbDownload:   false,
		LocalDirectory:    "wii",
	},
	"ps3": DbInfo{
		DbArchiveFileName: "ps3db.zip",
		DbRemoteURL:       "https://www.gametdb.com/ps3tdb.zip",
		ForceDbDownload:   false,
		LocalDirectory:    "ps3",
	},
	"wiiu": DbInfo{
		DbArchiveFileName: "wiiu.zip",
		DbRemoteURL:       "https://www.gametdb.com/wiiutdb.zip",
		ForceDbDownload:   false,
		LocalDirectory:    "wiiu",
	},
	"nds": DbInfo{
		DbArchiveFileName: "nds.zip",
		DbRemoteURL:       "https://www.gametdb.com/dstdb.zip",
		ForceDbDownload:   false,
		LocalDirectory:    "nds",
	},
	"3ds": DbInfo{
		DbArchiveFileName: "3ds.zip",
		DbRemoteURL:       "https://www.gametdb.com/3dstdb.zip",
		ForceDbDownload:   false,
		LocalDirectory:    "3ds",
	},
	"switch": DbInfo{
		DbArchiveFileName: "switch.zip",
		DbRemoteURL:       "https://www.gametdb.com/switchtdb.zip",
		ForceDbDownload:   false,
		LocalDirectory:    "switch",
	},
}
