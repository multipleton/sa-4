package engine

type Handler interface {
	Post(command Command)
}
