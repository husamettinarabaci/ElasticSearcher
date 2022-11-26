package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
	idsmo "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/object"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
)

func main() {

	cfg := elasticsearch7.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, err := elasticsearch7.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	var queryString = `
	{
		"query": {
		  "bool": {
			"should": [
			  {
				"wildcard": {
				  "brand": {
					"value": "*iss*"
				  }
				}
			  },
			  {
				"wildcard": {
				  "brand": {
					"value": "*iat*"
				  }
				}
			  }
			],
			"minimum_should_match": 1
		  }
		},
		"_source": [
		  "brand",
		  "model"
		]
	  }`

	var b strings.Builder
	b.WriteString(queryString)
	read := strings.NewReader(b.String())

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("postgre_vehicle"),
		es.Search.WithBody(read),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var result idsme.Result
		result.Id = idsmo.NewID(hit.(map[string]interface{})["_id"].(string))
		//result.Score = idsmo.NewScore(hit.(map[string]interface{})["_score"].(float64))
		//result.Source = idsmo.NewSource(hit.(map[string]interface{})["_source"].(map[string]interface{}))
		fmt.Println(result)
	}

}
