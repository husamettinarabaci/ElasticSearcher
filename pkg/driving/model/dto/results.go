package pkg_driving_dto

import pdmps "github.com/husamettinarabaci/ElasticSearcher/pkg/driving/model/proto/search"

type Results []Result

// TODO: implement this method
func (r Results) ToProtoResults() *pdmps.Results {

	var results []*pdmps.Result
	for _, res := range r {
		results = append(results, res.ToProtoResult())
	}

	return &pdmps.Results{
		Results: results,
	}

}
