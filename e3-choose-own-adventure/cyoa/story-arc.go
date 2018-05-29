package cyoa

// Story is the main story listing all Chapters.
type Story map[string]Chapter

// Chapter is part of the main story type.  It contains the title, the story
// text (as paragraphs), and the other stories that can be selected at the end.
type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

// Option contains the text to show to the user and next chapter title to choose
// when selected.
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
