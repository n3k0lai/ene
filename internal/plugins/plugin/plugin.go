package Plugin

type IPlugin interface {
	Test(query string) bool
	Reset()
}

type Plugin struct {
}
