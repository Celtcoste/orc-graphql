package graphql_tools

import "context"

type GraphQLRequest struct {
	url string
	request string
	variables []struct {
		name string
		value string
	}
	headers []struct {
		name string
		value string
	}
	ctx context.Context
	resp *interface{}
}
