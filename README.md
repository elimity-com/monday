# monday 
[![GoDoc](https://godoc.org/github.com/di-wu/monday?status.svg)](https://godoc.org/github.com/di-wu/monday)

monday is a go client library for accessing the [monday api v2](https://monday.com/developers/v2)

## querying monday.com entities
[GraphiQL](https://monday.com/developers/v2/try-it-yourself) (an in-browser tool for writing, validating, and testing GraphQL queries

### simple example
```
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

resp, err := NewClient(mondayAPIToken, nil).Exec(context.Background(), payload)
if err != nil {
    t.Error(err)
}
```