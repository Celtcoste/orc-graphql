package graphql_tools

import (
	graphqlClient "github.com/machinebox/graphql"
)

func GraphqlRequest(requestInformation *GraphQLRequest) (interface{}, error) {
	client := graphqlClient.NewClient(requestInformation.url)

	// make a request
	req := graphqlClient.NewRequest(requestInformation.request)

	// set any variables
	for i := 0; i < len(requestInformation.variables); i++ {
		req.Var(requestInformation.variables[i].name, requestInformation.variables[i].value)
	}
	//req.Var("playerUid", playerUID)
	// set header fields
	for i := 0; i < len(requestInformation.headers); i++ {
		req.Header.Set(requestInformation.headers[i].name, requestInformation.headers[i].value)
	}
	//req.Header.Set("Cache-Control", "no-cache")

	/*type PlayerInfo struct {
		PlayerInformation struct {
			Golds            int `json:"golds"`
			Diamonds         int `json:"diamonds"`
			InventorySize    int `json:"inventorySize"`
			InventorySizeMax int `json:"inventorySizeMax"`
		} `json:"playerInformation"`
	}*/

	if err := client.Run(requestInformation.ctx, req, requestInformation.resp); err != nil {
		return nil, err
	}
	return requestInformation.resp, nil
}