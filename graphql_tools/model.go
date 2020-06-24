package graphql_tools

import "context"

type GraphQLStruct struct {
	Url string
	Request string
	Variables []Content
	Headers []Content
	Ctx context.Context
	Resp *interface{}
}

type Content struct {
	Name string
	Value string
}