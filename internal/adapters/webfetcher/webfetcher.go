package webfetcher

import "net/http"

type WebFetcher interface {
	FetchUrl(string) (http.Response, error)
}
