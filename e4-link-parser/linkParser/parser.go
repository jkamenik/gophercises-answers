package linkParser

import (
	"errors"
	"io"
)

type Link struct {
	Href string
	Text string
}

func Parse(doc io.Reader) ([]Link, error) {
	return nil, errors.New("TBD")
}
