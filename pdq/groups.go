package pdq

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	. "github.com/di-wu/monday"
)

type Group struct {
	Id, Title string
}

func (g Group) equals(other Group) bool {
	if g.Id != other.Id || g.Title != other.Title {
		return false
	}
	return true
}

// EnsureGroup creates a group with the given title if it not already exists.
func (c SimpleClient) EnsureGroup(boardID int, title string) (Group, bool, error) {
	groups, err := c.GetGroups(boardID)
	if err != nil {
		return Group{}, false, err
	}
	var hit bool
	var group Group
	for _, g := range groups {
		if g.Title == title {
			hit = true
			group = g
			break
		}
	}
	if hit {
		return group, false, nil
	}
	group, err = c.CreateGroup(boardID, title)
	if err != nil {
		return Group{}, false, err
	}
	return group, true, nil
}

// CreateGroup creates a group with given name.
func (c SimpleClient) CreateGroup(boardID int, name string) (Group, error) {
	resp, err := c.Exec(context.Background(), NewMutationPayload(
		Groups.Create(boardID, name, []GroupsField{
			GroupsIDField(),
			GroupsTitleField(),
		}),
	))
	if err != nil {
		return Group{}, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Group{}, err
	}
	var body struct {
		Data struct {
			Group Group `json:"create_group"`
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return Group{}, err
	}
	return body.Data.Group, nil
}

// GetGroupWithID returns the group with given identifier.
func (c SimpleClient) GetGroupWithID(boardID int, groupID string) (Group, error) {
	resp, err := c.Exec(context.Background(), NewQueryPayload(
		Boards.List(
			[]BoardsField{
				NewBoardsGroupsFields(
					[]GroupsField{
						GroupsIDField(),
						GroupsTitleField(),
					},
					[]GroupsArgument{
						NewGroupsIDsArgument([]string{groupID}),
					},
				),
			},
			NewBoardsIDsArgument([]int{boardID}),
		),
	))
	if err != nil {
		return Group{}, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Group{}, err
	}
	var body struct {
		Data struct {
			Boards []struct {
				Groups []Group
			}
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return Group{}, err
	}
	if len(body.Data.Boards) != 1 {
		return Group{}, fmt.Errorf("no boards returned for id %d: %s", boardID, string(raw))
	}
	if len(body.Data.Boards[0].Groups) != 1 {
		return Group{}, fmt.Errorf("no groups returned for id %s in board %d: %s", groupID, boardID, string(raw))
	}
	return body.Data.Boards[0].Groups[0], nil
}

// GetGroups returns all the groups.
func (c SimpleClient) GetGroups(boardID int) ([]Group, error) {
	resp, err := c.Exec(context.Background(), NewQueryPayload(
		Boards.List(
			[]BoardsField{
				NewBoardsGroupsFields(
					[]GroupsField{
						GroupsIDField(),
						GroupsTitleField(),
					},
					nil,
				),
			},
			NewBoardsIDsArgument([]int{boardID}),
		),
	))
	if err != nil {
		return nil, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var body struct {
		Data struct {
			Boards []struct {
				Groups []Group
			}
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	if len(body.Data.Boards) != 1 {
		return nil, fmt.Errorf("no boards returned for id %d: %s", boardID, string(raw))
	}
	return body.Data.Boards[0].Groups, nil
}
