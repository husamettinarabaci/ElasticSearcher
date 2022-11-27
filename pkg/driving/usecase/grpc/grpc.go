package pkg_driving_usecase_grpc

import (
	"context"
	"net"

	iaspq "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/port/query"
	pdmps "github.com/husamettinarabaci/ElasticSearcher/pkg/driving/model/proto/search"
	"google.golang.org/grpc"

	pdmd "github.com/husamettinarabaci/ElasticSearcher/pkg/driving/model/dto"
)

type GrpcAPI struct {
	queryPort  iaspq.QueryPort
	gRpcServer *grpc.Server
	pdmps.UnimplementedSearchServer
}

func NewGrpcAPI(host string, port string, qp iaspq.QueryPort) *GrpcAPI {

	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}

	api := &GrpcAPI{
		queryPort:  qp,
		gRpcServer: grpc.NewServer(),
	}

	pdmps.RegisterSearchServer(api.gRpcServer, api)

	if err := api.gRpcServer.Serve(lis); err != nil {
		panic(err)
	}

	//TODO: wg.Done()
	return api
}

func (api GrpcAPI) Search(ctx context.Context, req *pdmps.Request) (*pdmps.Results, error) {

	reqDto := pdmd.FromProtoRequest(req)
	results, err := api.queryPort.Search(ctx, reqDto.ToEntity())
	if err != nil {
		return nil, err
	} else {
		var resDtos pdmd.Results
		for _, result := range results {
			resDtos = append(resDtos, pdmd.FromResultEntity(result))
		}
		return resDtos.ToProtoResults(), nil
	}
}

func (api GrpcAPI) SearchWildCards(ctx context.Context, wdilReq *pdmps.WildCardRequest) (*pdmps.Results, error) {

	wildStr := ""
	for _, wd := range wdilReq.Wildcards {
		wildStr += "{\"wildcard\":{\"" + wd.Field + "\":{\"value\":\"" + wd.Value + "\"}}},"
	}
	sourceStr := ""
	for _, sr := range wdilReq.Sources {
		sourceStr += "\"" + sr.Source + "\","
	}
	if len(wildStr) > 0 {
		wildStr = wildStr[:len(wildStr)-1]
	}
	if len(sourceStr) > 0 {
		sourceStr = sourceStr[:len(sourceStr)-1]
	}
	qStr := "{\"query\":{\"bool\":{\"should\":["
	qStr += wildStr
	qStr += "],\"minimum_should_match\":1}},\"_source\":["
	qStr += sourceStr
	qStr += "]}"

	req := pdmps.Request{
		Uid:   wdilReq.Uid,
		Index: wdilReq.Index,
		Query: qStr,
	}

	reqDto := pdmd.FromProtoRequest(&req)
	results, err := api.queryPort.Search(ctx, reqDto.ToEntity())
	if err != nil {
		return nil, err
	} else {
		var resDtos pdmd.Results
		for _, result := range results {
			resDtos = append(resDtos, pdmd.FromResultEntity(result))
		}
		return resDtos.ToProtoResults(), nil
	}
}
