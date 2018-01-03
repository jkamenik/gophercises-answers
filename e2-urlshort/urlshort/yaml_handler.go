package urlshort

import (
	"errors"
	"net/http"
)

func yamlHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return nil, errors.New("Foo")
}

func parseYAML(yml []byte) (map[string]string, error) {
	return nil, nil
}
