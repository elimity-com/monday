package pdq

import . "github.com/di-wu/monday"

// GetBoardWithID returns a query that gets the name and description of the board with given identifier.
func GetBoardWithID(id int) Query {
	return NewBoardsWithArguments(
		[]BoardsField{
			BoardsNameField(),
			BoardsDescriptionField(),
		},
		[]BoardsArgument{
			NewBoardsIDsArgument([]int{id}),
		},
	)
}
