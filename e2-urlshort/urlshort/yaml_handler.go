package urlshort

import (
	"errors"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func yamlHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return nil, errors.New("Foo")
}

type urlPath struct {
	URL  string `yaml: url`
	Path string `yaml: path`
}

func parseYAML(yml []byte) (urls map[string]string, err error) {
	var data []urlPath
	urls = make(map[string]string)

	err = yaml.Unmarshal(yml, &data)
	if err != nil {
		return
	}

	for _, item := range data {
		if item.Path == "" || item.URL == "" {
			continue
		}
		urls[item.Path] = item.URL
	}
	return
}
