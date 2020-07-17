package api

import (
	"net/http"
	"strconv"

	"github.com/sarpt/gamedbv/internal/status"
)

const (
	textFilterQuery string = "q"
	pageQuery       string = "_page"
	limitQuery      string = "_limit"
	platformQuery   string = "platform"
	regionQuery     string = "region"
	uidQuery        string = "uid"
	indexedQuery    string = "indexed"
)

func getTextQuery(r *http.Request) string {
	return r.URL.Query().Get(textFilterQuery)
}

func getCurrentPageQuery(r *http.Request) (int, error) {
	page := r.URL.Query().Get(pageQuery)
	if page == "" {
		return 0, nil
	}

	return strconv.Atoi(page)
}

func getPageLimitQuery(r *http.Request) (int, error) {
	limit := r.URL.Query().Get(limitQuery)
	if limit == "" {
		return -1, nil
	}

	return strconv.Atoi(limit)
}

func getPlatformsQuery(r *http.Request) []string {
	query := r.URL.Query()
	if platforms, ok := query[platformQuery]; ok {
		return platforms
	}

	return []string{}
}

func getRegionsQuery(r *http.Request) []string {
	query := r.URL.Query()
	if regions, ok := query[regionQuery]; ok {
		return regions
	}

	return []string{}
}

func getUIDQuery(r *http.Request) string {
	return r.URL.Query().Get(uidQuery)
}

func getIndexedQuery(r *http.Request) (status.FilterIndexing, error) {
	indexed := r.URL.Query().Get(indexedQuery)
	if indexed == "" {
		return status.AllPlatforms, nil
	}

	onlyIndexed, err := strconv.ParseBool(indexed)
	if err != nil {
		return status.AllPlatforms, nil
	}

	if onlyIndexed {
		return status.WithIndex, nil
	}

	return status.WithoutIndex, nil
}
