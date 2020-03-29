package serv

import (
	"net/http"
	"strconv"
)

const (
	textFilterQuery string = "q"
	pageQuery       string = "_page"
	limitQuery      string = "_limit"
	platformQuery   string = "platform"
	regionQuery     string = "region"
)

func getTextQueryFromRequest(r *http.Request) string {
	return r.URL.Query().Get(textFilterQuery)
}

func getCurrentPageFromRequest(r *http.Request) (int, error) {
	page := r.URL.Query().Get(pageQuery)
	if page == "" {
		return 0, nil
	}

	return strconv.Atoi(page)
}

func getPageLimitFromRequest(r *http.Request) (int, error) {
	limit := r.URL.Query().Get(limitQuery)
	if limit == "" {
		return -1, nil
	}

	return strconv.Atoi(limit)
}

func getPlatformsFromRequest(r *http.Request) []string {
	query := r.URL.Query()
	if platforms, ok := query[platformQuery]; ok {
		return platforms
	}

	return []string{}
}

func getRegionsFromRequest(r *http.Request) []string {
	query := r.URL.Query()
	if regions, ok := query[regionQuery]; ok {
		return regions
	}

	return []string{}
}
