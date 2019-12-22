# monday 
[![GoDoc](https://godoc.org/github.com/di-wu/monday?status.svg)](https://godoc.org/github.com/di-wu/monday)

monday is a go client library for accessing the [monday api v2](https://monday.com/developers/v2)

## using graphql
[GraphiQL](https://monday.com/developers/v2/try-it-yourself) (an in-browser tool for writing, validating, and testing GraphQL queries

## querying monday.com entities
```go
NewClient(mondayAPIToken, nil).Exec(context.Background(), NewQueryPayload(
    NewBoardsWithArguments(
        nil,
        NewBoardsIDsArgument([]int{boardID}),
    ),
))
```
the code above executes the following query
```graphql
query {
    boards(ids: boardID) {
       id
    }
}
```
## mutating monday.com entities
```go
NewClient(mondayAPIToken, nil).Exec(context.Background(), NewMutationPayload(
    Boards.Create(boardName, PublicBoardsKind(), nil),
))
```
the code above executes the following mutation
```graphql
mutation {
    create_board(board_name: boardName, board_kind: public) {
       id
    }
}
```