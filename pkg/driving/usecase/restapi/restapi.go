package pkg_driving_usecase_restapi

import (
	"github.com/gin-gonic/gin"
	iaspq "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/port/cqhandler"
	pdd "github.com/husamettinarabaci/ElasticSearcher/pkg/driving/dto"
)

type RestAPI struct {
	ginEngine    *gin.Engine
	queryHandler iaspq.QueryHandler
}

func NewRestAPI(c *gin.Engine, qh iaspq.QueryHandler) *RestAPI {
	api := &RestAPI{
		ginEngine:    c,
		queryHandler: qh,
	}
	if qh != nil {
		api.ginEngine.GET("/api/v1/search", api.GetResultByRequest)
	}
	return api
}

func (api *RestAPI) GetResultByRequest(c *gin.Context) {
	var reqDto pdd.Request
	err := c.BindJSON(&reqDto)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	results, err := api.queryHandler.GetResultByRequest(c, reqDto.ToEntity())
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var resDtos []pdd.Result
		for _, result := range results {
			resDtos = append(resDtos, pdd.FromResultEntity(result))
		}

		c.JSON(200, gin.H{
			"results": resDtos,
		})
		return
	}
}
