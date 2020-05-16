package platform

var (
	// Wii is Nintendo Wii platform
	Wii = Variant{
		id:   "wii",
		name: "Nintendo Wii",
	}
	// Wiiu is Nintendo WiiU platform
	Wiiu = Variant{
		id:   "wiiu",
		name: "Nintendo WiiU",
	}
	// Ps3 is PlayStation 3 platform
	Ps3 = Variant{
		id:   "ps3",
		name: "PlayStation 3",
	}
	// Nds is Nintendo DS platform
	Nds = Variant{
		id:   "nds",
		name: "Nintendo DS",
	}
	// N3ds is Nintendo 3DS platform
	N3ds = Variant{
		id:   "n3ds",
		name: "Nintendo 3DS",
	}
	// Switch is Nintendo Switch platfrom
	Switch = Variant{
		id:   "switch",
		name: "Nintendo Switch",
	}
)

var variants = map[string]Variant{
	Wii.id:    Wii,
	Wiiu.id:   Wiiu,
	Ps3.id:    Ps3,
	Nds.id:    Nds,
	N3ds.id:   N3ds,
	Switch.id: Switch,
}
