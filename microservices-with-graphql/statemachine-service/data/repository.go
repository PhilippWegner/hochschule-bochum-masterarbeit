package data

type Repository interface {
	GetStates(machine string, limit int) ([]*State, error)
	GetPlcs(machine string, time string, limit int) ([]*Plc, error)
	CreateState(state []*CreateStatesInput) error
}
