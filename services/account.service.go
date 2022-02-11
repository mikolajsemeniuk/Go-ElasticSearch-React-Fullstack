package services

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/extensions"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/payloads"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/repositories"
)

var AccountService IAccountService = &accountService{}

type IAccountService interface {
	FindAccounts() ([]payloads.Account, error)
}

type accountService struct{}

func (*accountService) FindAccounts() ([]payloads.Account, error) {
	payloads := []payloads.Account{}
	buffer := bytes.Buffer{}
	query := map[string]interface{}{}

	if err := json.NewEncoder(&buffer).Encode(query); err != nil {
		message := fmt.Errorf("error while encoding to buffer from json, %s", err.Error())
		extensions.Error(message.Error())
		return nil, message
	}

	entites, err := repositories.AccountRepository.FindAccounts(buffer)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&payloads, &entites)
	if err != nil {
		message := fmt.Errorf("error while copying from entity to payload, %s", err.Error())
		extensions.Error(message.Error())
		return nil, message
	}

	extensions.Info("done")
	return payloads, err
}
