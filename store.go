package telegram

type Store interface {
	SetState(user int, state string, field []string) error
	getState(user int) State
}

