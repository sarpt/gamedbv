package platform

// All returns all possible variants of platform databases that could be downloaded
func All() []Variant {
	var allVariants []Variant

	for _, variant := range variants {
		allVariants = append(allVariants, variant)
	}

	return allVariants
}

// ByIds returns all variants that match provided ids
func ByIds(ids []string) []Variant {
	var found []Variant

	for _, id := range ids {
		if variant, ok := variants[id]; ok {
			found = append(found, variant)
		}
	}

	return found
}
