package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/data"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/entities"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/extensions"
)

const index = "accounts"

var AccountRepository IAccountRepository = &accountRepository{}

type IAccountRepository interface {
	FindAccounts(buffer bytes.Buffer) ([]entities.Account, error)
}

type accountRepository struct{}

func (*accountRepository) FindAccounts(buffer bytes.Buffer) ([]entities.Account, error) {
	type signal struct {
		Accounts []entities.Account
		Error    error
	}

	channel := make(chan signal)
	go func() {
		response, err := data.ElasticSearch.Search(
			data.ElasticSearch.Search.WithContext(context.Background()),
			data.ElasticSearch.Search.WithIndex(index),
			data.ElasticSearch.Search.WithBody(&buffer),
		)

		defer response.Body.Close()
		if err != nil {
			message := fmt.Errorf("error while fetching records from database, %s", err.Error())
			extensions.Error(message.Error())
			channel <- signal{nil, message}
			return
		}

		if response.IsError() {
			message := fmt.Errorf("error indexing documents status: %s", response.Status())
			extensions.Error(message.Error())
			channel <- signal{nil, message}
			return
		}

		body := make(map[string]interface{})
		if err = json.NewDecoder(response.Body).Decode(&body); err != nil {
			message := fmt.Errorf("error parsing the response body: %s", err.Error())
			extensions.Error(message.Error())
			channel <- signal{nil, message}
			return
		}

		accounts := []entities.Account{}
		for _, hit := range body["hits"].(map[string]interface{})["hits"].([]interface{}) {
			var account entities.Account

			err = extensions.Decode(hit.(map[string]interface{})["_source"].(map[string]interface{}), &account)
			if err != nil {
				message := fmt.Errorf("error mapping from _source to entity: %s", err.Error())
				extensions.Error(message.Error())
				channel <- signal{nil, message}
				return
			}

			accounts = append(accounts, account)
		}

		channel <- signal{accounts, nil}
	}()

	result := <-channel
	extensions.Info("done")
	return result.Accounts, result.Error
}
