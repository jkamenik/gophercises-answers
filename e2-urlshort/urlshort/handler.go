package urlshort

import (
	"errors"
	"net/http"
)

var (
	// ErrEmptyMapNoDefault is returned if the map length is zero and there
	// is no default handler.  To prevent situtations where there is no default
	// behavior always provide a fallback.
	ErrEmptyMapNoDefault = errors.New("urlshort: Empty Map with no Default Handler")
	ErrYAMLParseError    = errors.New("urlshort: YAML parse error")
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) (http.HandlerFunc, error) {
	return mapHandler(pathsToUrls, fallback)
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return yamlHandler(yml, fallback)
}
