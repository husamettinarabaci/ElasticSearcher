package pkg_driven_infra_adapter_elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	_ "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/port/searchengine"
	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
	idsmo "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/object"
)

type ElasticSearchEngine struct {
	config elasticsearch7.Config
	engine *elasticsearch7.Client
}

func NewElasticSearchEngine(host string, port string) ElasticSearchEngine {

	cfg := elasticsearch7.Config{
		Addresses: []string{
			"http://" + host + ":" + port,
		},
	}

	es, err := elasticsearch7.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	elastic := ElasticSearchEngine{
		config: cfg,
		engine: es,
	}

	return elastic
}

func (ese ElasticSearchEngine) Search(ctx context.Context, req idsme.Request) ([]idsme.Result, error) {
	var results []idsme.Result
	var b strings.Builder
	b.WriteString(req.Query.String())
	read := strings.NewReader(b.String())

	res, err := ese.engine.Search(
		ese.engine.Search.WithContext(context.Background()),
		ese.engine.Search.WithIndex(req.Index.String()),
		ese.engine.Search.WithBody(read),
		ese.engine.Search.WithTrackTotalHits(true),
		ese.engine.Search.WithPretty(),
	)
	if err != nil {
		return results, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return results, err
		} else {
			return results, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return results, err
	}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var result idsme.Result
		result.Id = idsmo.NewID(hit.(map[string]interface{})["_id"].(string))
		result.Score = idsmo.NewScore(hit.(map[string]interface{})["_score"].(float64))
		result.Source = idsmo.NewSource(hit.(map[string]interface{})["_source"].(map[string]interface{}))

		results = append(results, result)
	}

	return results, nil
}
