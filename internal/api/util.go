package api

import (
	"net/http"

	"github.com/sarpt/gamedbv/pkg/platform"
)

func getPlatformVariants(r *http.Request) []platform.Variant {
	platforms := getPlatformsQuery(r)

	if len(platforms) == 0 {
		return platform.All()
	}

	return platform.ByIds(platforms)
}
