package monday

type Payload struct {
	queries []Query
}

func NewQueryPayload(queries ...Query) Payload {
	return Payload{queries: queries}
}

type Query struct {
	str string
}

func NewBoardsQuery(boards Boards) Query {
	return Query{str: boards.stringify()}
}

func NewComplexityQuery(complexity Complexity) Query {
	return Query{str: complexity.stringify()}
}
