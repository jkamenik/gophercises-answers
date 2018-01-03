package urlshort

import (
	"reflect"
	"testing"
)

type yamlTestInput struct {
	yaml   string
	err    error
	parsed map[string]string
}

// If both the default handler and the map has no value then throw an error
func TestYAMLToMap(t *testing.T) {
	yaml := [...]yamlTestInput{
		yamlTestInput{
			`this is invalid`,
			ErrYAMLParseError,
			nil,
		},
		yamlTestInput{
			`
- url: http://google.com
	path: /foo
			`,
			nil,
			map[string]string{
				"url":  "http://google.com",
				"path": "/foo",
			},
		},
		yamlTestInput{
			`
- url: http://yahoo.com
  path: /bar
  ignored: true
	    `,
			nil,
			map[string]string{
				"url":  "http://yahoo.com",
				"path": "/bar",
			},
		},
		yamlTestInput{
			`
- URL: bad
			`,
			ErrYAMLParseError,
			nil,
		},
	}

	for _, y := range yaml {
		maps, err := parseYAML([]byte(y.yaml))

		if err != y.err {
			t.Errorf("error '%v' is not '%v' for '%v'", err, y.err, y.yaml)
		}

		if !reflect.DeepEqual(maps, y.parsed) {
			t.Errorf("map '%v' is not '%v' for '%v'", maps, y.parsed, y.yaml)
		}
	}
}
