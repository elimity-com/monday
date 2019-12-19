package pdq

import (
	"context"

	. "github.com/di-wu/monday"
)

func (c SimpleClient) EnsureColumnValue(boardID int, itemID int, value ColumnValue) error {
	_, err := c.Exec(context.Background(), NewMutationPayload(
		ChangeColumnValue(
			itemID, value.ID(), boardID, value, nil,
		),
	))
	return err
}
