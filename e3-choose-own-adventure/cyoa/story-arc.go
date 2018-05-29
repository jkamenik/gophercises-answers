package cyoa

import (
	"encoding/json"
	"io"
)

// Story is the main story listing all Chapters.
type Story map[string]Chapter

// Chapter is part of the main story type.  It contains the title, the story
// text (as paragraphs), and the other stories that can be selected at the end.
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option contains the text to show to the user and next chapter title to choose
// when selected.
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// ParseJSON parses a JSON like story and returns a story.
func ParseJSON(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story

	err := d.Decode(&story)
	return story, err
}
