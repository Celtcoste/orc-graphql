package graphql_tools

import "context"

type GraphQLRequest struct {
	url string
	request string
	variables []Content
	headers []Content
	ctx context.Context
	resp *interface{}
}

type Content struct {
	name string
	value string
}