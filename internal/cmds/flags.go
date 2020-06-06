package cmds

const (
	// PlatformFlag is used to specify console platforms
	PlatformFlag = "platform"

	// JSONFlag is used to specify whether json output should be used
	JSONFlag = "json"

	// TextFlag is used to specify text to be searched in games titles or descriptions
	TextFlag = "text"

	// LanguageFlag is used to specify language code
	LanguageFlag = "language"

	// RegionFlag is used to specify region (PAL, NTSC-U, etc.)
	RegionFlag = "region"

	// PageFlag is used to specify page of paging mechanism (games etc.)
	PageFlag = "page"

	// PageLimitFlag is used to specify maximum number of entries per page in paging mechanism
	PageLimitFlag = "page-limit"

	// DebugFlag is used to specify whether command should run in development debug mode. Specific behavior depends on the command
	DebugFlag = "debug"

	// InterfaceFlag is used to specify network interface used for binary to listen on
	InterfaceFlag = "interface"
)
