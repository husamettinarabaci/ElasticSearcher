package pkg_driving_usecase_restapi

import (
	"github.com/gin-gonic/gin"
	iaspq "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/port/query"
	pdmd "github.com/husamettinarabaci/ElasticSearcher/pkg/driving/model/dto"
)

type RestAPI struct {
	ginEngine *gin.Engine
	queryPort iaspq.QueryPort
}

func NewRestAPI(host string, port string, qp iaspq.QueryPort) *RestAPI {
	api := &RestAPI{
		ginEngine: gin.Default(),
		queryPort: qp,
	}
	if qp != nil {
		api.ginEngine.GET("/api/v1/search", api.Search)
	}

	err := api.ginEngine.Run(host + ":" + port)
	if err != nil {
		panic(err)
	}
	//TODO: wg.Done()
	return api
}

func (api *RestAPI) Search(c *gin.Context) {
	var reqDto pdmd.Request
	err := c.BindJSON(&reqDto)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	results, err := api.queryPort.Search(c, reqDto.ToEntity())
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var resDtos []pdmd.Result
		for _, result := range results {
			resDtos = append(resDtos, pdmd.FromResultEntity(result))
		}

		c.JSON(200, gin.H{
			"results": resDtos,
		})
		return
	}
}
