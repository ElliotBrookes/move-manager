package webfetcher

import (
	"ElliotBrookes/move-manager/internal/adapters/webfetcher"
	"ElliotBrookes/move-manager/internal/adapters/webfetcher/mapbox"
)

const (
	MapBoxInstance = iota
	NinjaApiInstance
)

func NewWebFetcherFactory(iType int) (webfetcher.WebFetcher, error) {
    switch iType {
    case MapBoxInstance:

        return mapbox.

    case NinjaApiInstance:
        return 
}
