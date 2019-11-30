package monday

type State struct {
	state string
}

var (
	allState      = State{"all"}
	activeState   = State{"active"}
	archivedState = State{"archived"}
	deletedState  = State{"deleted"}
)

func AllState() State {
	return allState
}

func ActiveState() State {
	return activeState
}

func ArchivedState() State {
	return archivedState
}

func DeletedState() State {
	return deletedState
}
