package pastebin

import "time"

// PasteObj the definition
type PasteObj struct {
	Poster     string
	Syntax     string // enum
	Content    string
	TinyURL    string
	Expiration time.Duration
}
