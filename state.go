package monday

// The possible states for a board or item.
type State struct {
	state string
}

var (
	allState      = State{"all"}
	activeState   = State{"active"}
	archivedState = State{"archived"}
	deletedState  = State{"deleted"}
)

// Active, Archived and Deleted.
func AllState() State {
	return allState
}

// Active only (Default).
func ActiveState() State {
	return activeState
}

// Archived only.
func ArchivedState() State {
	return archivedState
}

// Deleted only.
func DeletedState() State {
	return deletedState
}
