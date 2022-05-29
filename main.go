package main

import (
	"context"
	"fmt"

	"github.com/shurcooL/graphql"
)

func main() {
	/**
	 * This is the Kenshi Deep Index endpoint for GraphQL
	 */
	endpoint := "https://api.kenshi.io/index/graphql"

	/**
	 * Create a GraphQL client to connect to the GraphQL endpoint
	 */

	client := graphql.NewClient(endpoint, nil)

	/**
	 * Define your GraphQL query here
	 * Note: Refer to the Kenshi GraphQL schema documentation located at
	 * https://docs.kenshi.io/services/deep-index/graphql/schema.html
	 */
	var query struct {
		GetEntries []struct {
			Event struct {
				Name graphql.String
			}
			Block struct {
				Number graphql.Int
			}
		} `graphql:"getEntries(blockchain: \"binance-testnet\", apikey: \"API_KEY\", owner: \"API_KEY_OWNER\")"`
	}

	/**
	 * Send the query to the GraphQL server
	 */
	err := client.Query(context.Background(), &query, nil)

	/**
	 * Check for errors and print the retrieved data
	 */
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(query.GetEntries)
	}
}
