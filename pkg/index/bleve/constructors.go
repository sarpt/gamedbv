package bleve

import (
	"github.com/blevesearch/bleve"
	"github.com/sarpt/gamedbv/pkg/index"
)

// NewSearcher initializes bleve implementation of Searcher
func NewSearcher(configs []index.PlatformConfig) (*Searcher, []index.PlatformConfig) {
	var ignored []index.PlatformConfig
	s := &Searcher{
		indexes: make(map[string]bleve.Index),
	}

	for _, conf := range configs {
		err := s.AddIndex(conf)
		if err != nil {
			ignored = append(ignored, conf)
		}
	}

	return s, ignored
}
