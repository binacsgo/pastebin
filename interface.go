package pastebin

// PasteBin the interface def.
type PasteBin interface {
	Syntax() string
	ParseContent(string) (string, error)
}
