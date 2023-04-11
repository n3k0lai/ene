package core

type IPlugin interface {
	TestTrigger(query string)
}

func TestTrigger(query string) {

}