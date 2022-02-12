package data

import (
	"fmt"
	"log"

	"github.com/TwiN/go-color"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/settings"
)

var ElasticSearch *elasticsearch.Client

func init() {
	configuration := elasticsearch.Config{
		Addresses: []string{
			settings.Configuration.GetString("database.elasticsearch.connectionstring"),
		},
	}

	var err error
	ElasticSearch, err = elasticsearch.NewClient(configuration)

	if err != nil {
		panic(fmt.Errorf("error creating the client: %s", err.Error()))
	}

	response, err := ElasticSearch.Info()
	if err != nil {
		panic(fmt.Errorf("error getting response: %s", err))
	}

	defer response.Body.Close()
	log.Println(color.Ize(color.Green, response.String()))
}

func GetInfo() {
	log.Printf("Client: %s", elasticsearch.Version)
}
