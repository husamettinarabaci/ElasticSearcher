package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	pdmps "github.com/husamettinarabaci/ElasticSearcher/pkg/driving/model/proto/search"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pdmps.NewSearchClient(conn)

	//TestSearch(client)
	TestSearchWildCards(client)
}

func TestSearch(client pdmps.SearchClient) {

	req := &pdmps.Request{
		Uid:   uuid.NewString(),
		Index: "postgre_vehicle",
		Query: "{\"query\":{\"bool\":{\"should\":[{\"wildcard\":{\"brand\":{\"value\":\"*iss*\"}}},{\"wildcard\":{\"brand\":{\"value\":\"*iat*\"}}}],\"minimum_should_match\":1}},\"_source\":[\"brand\",\"model\"]}",
	}

	results, err := client.Search(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling Search: %s", err)
	}

	log.Printf("%v", results)
}

func TestSearchWildCards(client pdmps.SearchClient) {

	req := &pdmps.WildCardRequest{
		Uid:   uuid.NewString(),
		Index: "postgre_vehicle",
		Wildcards: []*pdmps.WildCard{
			{
				Field: "brand",
				Value: "*iss*",
			},
			{
				Field: "brand",
				Value: "*iat*",
			},
		},
		Sources: []*pdmps.Source{
			{
				Source: "brand",
			},
			{
				Source: "model",
			},
		},
	}

	results, err := client.SearchWildCards(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling Search: %s", err)
	}

	log.Printf("%v", results)
}
