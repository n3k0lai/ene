package Plugins

type IPlugin interface {
	Test(query string) bool
}

type Plugin struct {
}
