package pdq

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	. "github.com/di-wu/monday"
)

type Item struct {
	Id, Name, Value string
}

func (i Item) ID() int {
	id, _ := strconv.Atoi(i.Id)
	return id
}

func (c SimpleClient) GetItemColumnValues(itemID int) ([]map[string]interface{}, error) {
	resp, err := c.Exec(context.Background(), NewQueryPayload(
		NewItemsWithArguments(
			[]ItemsField{
				NewItemsColumnValuesField(
					[]ColumnValuesField{
						ColumnValuesValueField(),
					},
					nil,
				),
			},
			[]ItemsArgument{
				NewItemsIDsArgument([]int{itemID}),
			},
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
			Items []struct {
				Values []struct {
					Value string
				} `json:"column_values"`
			}
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	if len(body.Data.Items) != 1 {
		return nil, fmt.Errorf("no items returned for id %d: %s", itemID, string(raw))
	}
	if len(body.Data.Items[0].Values) < 1 {
		return nil, fmt.Errorf("no values returned")
	}
	var values []map[string]interface{}
	for _, value := range body.Data.Items[0].Values {
		var m map[string]interface{}
		if value.Value == "" {
			values = append(values, nil)
			continue
		}
		if err := json.Unmarshal([]byte(value.Value), &m); err != nil {
			return nil, err
		}
		values = append(values, m)
	}
	return values, nil
}

func (i Item) equals(other Item) bool {
	if i.Id != other.Id || i.Name != other.Name {
		return false
	}
	return true
}

// EnsureItem creates an item with the given name if it not already exists.
func (c SimpleClient) EnsureItem(boardID int, groupID string, name string) (Item, bool, error) {
	items, err := c.GetItems(boardID, groupID)
	if err != nil {
		return Item{}, false, err
	}
	var hit bool
	var item Item
	for _, i := range items {
		if i.Name == name {
			hit = true
			item = i
			break
		}
	}
	if hit {
		return item, false, nil
	}
	item, err = c.CreateItem(boardID, groupID, name)
	if err != nil {
		return Item{}, false, err
	}
	return item, true, nil
}

// CreateItem creates an item with the given name.
func (c SimpleClient) CreateItem(boardID int, groupID string, name string) (Item, error) {
	resp, err := c.Exec(context.Background(), NewMutationPayload(
		CreateItem(
			boardID, groupID, name, nil,
			[]ItemsField{
				ItemsIDField(),
				ItemsNameField(),
			},
		),
	))
	if err != nil {
		return Item{}, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Item{}, err
	}
	var body struct {
		Data struct {
			Item Item `json:"create_item"`
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return Item{}, err
	}
	return body.Data.Item, nil
}

func (c SimpleClient) CreateItemWithColumnValues(boardID int, groupID string, name string, columnValues []ColumnValue) (Item, error) {
	resp, err := c.Exec(context.Background(), NewMutationPayload(
		CreateItem(
			boardID, groupID, name, columnValues,
			[]ItemsField{
				ItemsIDField(),
				ItemsNameField(),
			},
		),
	))
	if err != nil {
		return Item{}, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Item{}, err
	}
	var body struct {
		Data struct {
			Item Item `json:"create_item"`
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return Item{}, err
	}
	return body.Data.Item, nil
}

// GetItemWithID return the item with the given identifier.
func (c SimpleClient) GetItemWithID(itemID int) (Item, error) {
	resp, err := c.Exec(context.Background(), NewQueryPayload(
		NewItemsWithArguments(
			[]ItemsField{
				ItemsIDField(),
				ItemsNameField(),
			},
			[]ItemsArgument{
				NewItemsIDsArgument([]int{itemID}),
			},
		),
	))
	if err != nil {
		return Item{}, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Item{}, err
	}
	var body struct {
		Data struct {
			Items []Item
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return Item{}, err
	}
	if len(body.Data.Items) != 1 {
		return Item{}, fmt.Errorf("no items returned for id %d: %s", itemID, string(raw))

	}
	return body.Data.Items[0], nil
}

// GetItems returns all the items.
func (c SimpleClient) GetItems(boardID int, groupID string) ([]Item, error) {
	resp, err := c.Exec(context.Background(), NewQueryPayload(
		NewBoardsWithArguments(
			[]BoardsField{
				NewBoardsGroupsFields(
					[]GroupsField{
						NewGroupsItemsField(
							[]ItemsField{
								ItemsIDField(),
								ItemsNameField(),
							},
							nil,
						),
					},
					[]GroupsArgument{
						NewGroupsIDsArgument([]string{groupID}),
					},
				),
			},
			[]BoardsArgument{
				NewBoardsIDsArgument([]int{boardID}),
			},
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
				Groups []struct {
					Items []Item
				}
			}
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	if len(body.Data.Boards) != 1 {
		return nil, fmt.Errorf("no boards returned for id %d: %s", boardID, string(raw))
	}
	if len(body.Data.Boards[0].Groups) != 1 {
		return nil, fmt.Errorf("no groups returned for id %s: %s", groupID, string(raw))
	}
	return body.Data.Boards[0].Groups[0].Items, nil
}
