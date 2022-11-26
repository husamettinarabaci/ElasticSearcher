package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	iasac "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/adapter/cqhandler"
	iaspc "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/port/cqhandler"
	iass "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/service"
	idss "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/service"
	pdur "github.com/husamettinarabaci/ElasticSearcher/pkg/driving/usecase/restapi"
	"gopkg.in/yaml.v2"
)

const configPath = "config/application/search.yml"

type Config struct {
	Debug bool `yaml:"debug"`
	Cli   struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"cli"`
	Restapi struct {
		Enabled      bool   `yaml:"enabled"`
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		QueryHandler bool   `yaml:"queryHandler"`
	} `yaml:"restapi"`
}

var AppConfig *Config

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ReadConfig()
	var err error
	c := gin.Default()

	cont := container.New()

	err = cont.Singleton(func() idss.SearchService {
		return idss.NewSearchService()
	})
	if err != nil {
		panic(err)
	}

	err = cont.Singleton(func(s idss.SearchService) iass.SearchService {
		return iass.NewSearchService(s)
	})
	if err != nil {
		panic(err)
	}

	err = cont.Singleton(func(s iass.SearchService) iasac.QueryHandler {
		return iasac.NewQueryHandler(s)
	})
	if err != nil {
		panic(err)
	}

	if AppConfig.Restapi.Enabled {

		var queryHandler iaspc.QueryHandler
		if AppConfig.Restapi.QueryHandler {
			err = cont.Resolve(&queryHandler)
			if err != nil {
				panic(err)
			}
		}
		pdur.NewRestAPI(c, queryHandler)
		c.Run(AppConfig.Restapi.Host + ":" + AppConfig.Restapi.Port)
	}
}
