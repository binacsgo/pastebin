package pastebin

import (
	"bytes"
	"fmt"

	"github.com/russross/blackfriday/v2"

	emoji "github.com/yuin/goldmark-emoji"
	mermaid "go.abhg.dev/goldmark/mermaid"

	// toc "github.com/abhinav/goldmark-toc"
	// highlighting "github.com/yuin/goldmark-highlighting"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var (
	gMarkDownImpl *MarkDownImpl
	goldMarkImpl  goldmark.Markdown
)

func init() {
	gMarkDownImpl = new(MarkDownImpl)
	goldMarkImpl = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Footnote,
			extension.DefinitionList,
			extension.NewTypographer(
				extension.WithTypographicSubstitutions(extension.TypographicSubstitutions{
					extension.LeftSingleQuote:  []byte("&sbquo;"),
					extension.RightSingleQuote: nil, // nil disables a substitution
				}),
			),
			emoji.Emoji,
			&mermaid.Extender{},
			// highlighting.Highlighting,
			// &toc.Extender{},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
}

// ParseContentMD parse the markdown syntax directly
func ParseContentMD(input string) (string, error) {
	return gMarkDownImpl.ParseContent(input)
}

// MarkDownImpl the markdown impl.
type MarkDownImpl struct {
	useBlackFriday bool
}

// Syntax return the syntax
func (impl *MarkDownImpl) Syntax() string {
	return "markdown"
}

// ParseContent return the
func (impl *MarkDownImpl) ParseContent(input string) (string, error) {
	if impl.useBlackFriday {
		output := blackfriday.Run(
			bytes.Replace([]byte(input), []byte("\r"), nil, -1),
			blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.HardLineBreak),
		)
		return string(output), nil
	}
	var buf bytes.Buffer
	if err := goldMarkImpl.Convert([]byte(input), &buf); err != nil {
		return fmt.Sprint(err), err
	}
	return buf.String(), nil
}
