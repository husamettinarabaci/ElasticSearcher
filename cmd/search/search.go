package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/golobby/container/v3"
	iasaq "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/adapter/query"
	iass "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/service"
	idss "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/service"
	pdiae "github.com/husamettinarabaci/ElasticSearcher/pkg/driven/infra/adapter/elastic"
	pdug "github.com/husamettinarabaci/ElasticSearcher/pkg/driving/usecase/grpc"
	pdur "github.com/husamettinarabaci/ElasticSearcher/pkg/driving/usecase/restapi"
	"gopkg.in/yaml.v2"
)

const configPath = "config/application/search.yml"

type Config struct {
	Debug   bool `yaml:"debug"`
	Restapi struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"restapi"`
	GrpcAPI struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"grpcapi"`
	ElasticEngine struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"elastic"`
}

var AppConfig *Config
var wg sync.WaitGroup

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

	cont := container.New()

	CreateDomainContainer(&cont)

	CreateInfraContainer(&cont)

	CreateApplicationContainer(&cont)

	CreatePresentationContainer(&cont)

	fmt.Println("Started Application")
	wg.Wait()
}

func CreateDomainContainer(cont *container.Container) {

	err := cont.Singleton(func() idss.CoreService {
		return idss.NewCoreService()
	})
	if err != nil {
		panic(err)
	}

}

func CreateInfraContainer(cont *container.Container) {
	err := cont.Singleton(func() pdiae.ElasticSearchEngine {
		return pdiae.NewElasticSearchEngine(AppConfig.ElasticEngine.Host, AppConfig.ElasticEngine.Port)
	})
	if err != nil {
		panic(err)
	}
}

func CreateApplicationContainer(cont *container.Container) {
	err := cont.Singleton(func(s pdiae.ElasticSearchEngine) iass.EventListener {
		return iass.NewEventListener(s)
	})
	if err != nil {
		panic(err)
	}

	err = cont.Singleton(func(s idss.CoreService, e iass.EventListener) iass.CoreService {
		return iass.NewCoreService(s, e)
	})
	if err != nil {
		panic(err)
	}

	err = cont.Singleton(func(s iass.CoreService) iass.QueryHandler {
		return iass.NewQueryHandler(s)
	})
	if err != nil {
		panic(err)
	}

	err = cont.Singleton(func(s iass.QueryHandler) iasaq.QueryAdapter {
		return iasaq.NewQueryAdapter(s)
	})
	if err != nil {
		panic(err)
	}
}

func CreatePresentationContainer(cont *container.Container) {

	var queryPort iasaq.QueryAdapter
	err := cont.Resolve(&queryPort)
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go pdur.NewRestAPI(AppConfig.Restapi.Host, AppConfig.Restapi.Port, queryPort)
	wg.Add(1)
	go pdug.NewGrpcAPI(AppConfig.GrpcAPI.Host, AppConfig.GrpcAPI.Port, queryPort)

}
