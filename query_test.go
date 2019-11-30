package monday

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
)

var boardID int
var mondayAPIToken string
var client *Client

func TestMain(m *testing.M) {
	id, ok := os.LookupEnv("BOARD_ID")
	if !ok {
		log.Println("could not get board id from env")
		return
	}
	boardID, _ = strconv.Atoi(id)
	mondayAPIToken, ok = os.LookupEnv("MONDAY_API_TOKEN")
	if !ok {
		log.Println("could not get monday api token from env")
		return
	}
	client = NewClient(mondayAPIToken, nil)
	m.Run()
}

func TestSimpleBoardsQuery(t *testing.T) {
	payload := NewQueryPayload(
		NewBoardsQuery(
			NewBoardsWithArguments(
				[]BoardsField{
					BoardsIDField(),
				},
				[]BoardsArgument{
					NewIDsBoardsArg([]int{boardID}),
				},
			),
		),
	)
	resp, err := client.Exec(context.Background(), payload)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("non-200 status code returned: %d\n", resp.StatusCode)
	}

	raw, _ := ioutil.ReadAll(resp.Body)
	var response struct {
		Data struct {
			Boards []struct {
				Id string
			}
			Complexity struct {
				After  int
				Before int
				Query  int
			}
		}
	}
	if err := json.Unmarshal(raw, &response); err != nil {
		t.Error(err)
	}

	if len(response.Data.Boards) != 1 {
		t.Error("did not get board\n")
		return
	}

	id, err := strconv.Atoi(response.Data.Boards[0].Id)
	if err != nil {
		t.Error(err)
	}
	if id != boardID {
		t.Errorf("got: %d, expected: %d\n", id, boardID)
	}
}
