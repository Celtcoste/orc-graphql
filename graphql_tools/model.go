package graphql_tools

import "context"

type GraphQLRequest struct {
	url string
	request string
	variables []content
	headers []content
	ctx context.Context
	resp *interface{}
}

type content struct {
	name string
	value string
}