package data

type Repository interface {
	GetPlcs(machine string, time string, limit int) ([]*Plc, error)
	CreateState(state []*State) error
}
