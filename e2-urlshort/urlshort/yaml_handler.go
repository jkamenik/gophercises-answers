package urlshort

import (
	"errors"
	"fmt"
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

func parseYAML(yml []byte) (map[string]string, error) {
	var data []urlPath

	err := yaml.Unmarshal(yml, data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("yaml data: %v\n", data)
	return nil, nil
}
