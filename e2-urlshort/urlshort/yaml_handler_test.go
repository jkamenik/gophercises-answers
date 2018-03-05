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
  path: /foo`,
			nil,
			map[string]string{
				"/foo": "http://google.com",
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
				"/bar": "http://yahoo.com",
			},
		},
		yamlTestInput{
			`
- URL: bad
`,
			nil,
			map[string]string{},
		},
	}

	for _, y := range yaml {
		maps, err := parseYAML([]byte(y.yaml))

		if y.err != nil {
			// we expect the input to be a failure
			if err == nil {
				t.Error("Expected an error, but none found")
			}
		}

		if y.parsed != nil {
			// expect the values to equal
			if !reflect.DeepEqual(maps, y.parsed) {
				t.Errorf("map '%v' is not '%v' for '%v'", maps, y.parsed, y.yaml)
			}
		}
	}
}
