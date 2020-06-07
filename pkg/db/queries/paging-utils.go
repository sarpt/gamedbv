package queries

import "math"

func limitForPage(page, pagesCount, limit, maxLimit int) int {
	lastPage := page+1 >= pagesCount
	if !lastPage {
		return maxLimit
	}

	if limit <= maxLimit {
		return limit
	}

	return limit - maxLimit*page
}

func offsets(limit, maxLimit, chosenPage int) []int {
	if (limit <= maxLimit) || maxLimit == 0 {
		return []int{pageOffset(limit, chosenPage)}
	}

	var offsets []int
	neededPagesCount := int(math.Ceil(float64(limit) / float64(maxLimit)))
	for page := 0; page < neededPagesCount; page++ {
		offsets = append(offsets, pageOffset(maxLimit, page))
	}

	return offsets
}

func pageOffset(limit, page int) int {
	return page * limit
}
