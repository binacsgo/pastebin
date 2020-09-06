package pastebin

import (
	"bytes"

	"github.com/russross/blackfriday/v2"
)

var gMarkDownImpl *MarkDownImpl

func init() {
	gMarkDownImpl = new(MarkDownImpl)
}

// ParseContentMD parse the markdown syntax directly
func ParseContentMD(input string) (string, error) {
	return gMarkDownImpl.ParseContent(input)
}

// MarkDownImpl the markdown impl.
type MarkDownImpl struct{}

// Syntax return the syntax
func (impl *MarkDownImpl) Syntax() string {
	return "markdown"
}

// ParseContent return the
func (impl *MarkDownImpl) ParseContent(input string) (string, error) {
	output := blackfriday.Run(
		bytes.Replace([]byte(input), []byte("\r"), nil, -1),
		blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.HardLineBreak),
	)
	return string(output), nil
}
