package serv

import (
	"net/http"
	"strconv"
)

func getTextQueryFromRequest(r *http.Request) string {
	return r.URL.Query().Get("q")
}

func getCurrentPageFromRequest(r *http.Request) (int, error) {
	page := r.URL.Query().Get("_page")
	if page == "" {
		return 0, nil
	}

	return strconv.Atoi(page)
}

func getPageLimitFromRequest(r *http.Request) (int, error) {
	limit := r.URL.Query().Get("_limit")
	if limit == "" {
		return -1, nil
	}

	return strconv.Atoi(limit)
}

func getPlatformsFromRequest(r *http.Request) []string {
	query := r.URL.Query()
	if platforms, ok := query["platform"]; ok {
		return platforms
	}

	return []string{}
}

func getRegionsFromRequest(r *http.Request) []string {
	query := r.URL.Query()
	if regions, ok := query["region"]; ok {
		return regions
	}

	return []string{}
}
