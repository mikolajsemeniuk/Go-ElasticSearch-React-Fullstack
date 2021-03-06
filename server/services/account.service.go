package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/entities"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/extensions"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/inputs"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/payloads"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/repositories"
)

var AccountService IAccountService = &accountService{}

type IAccountService interface {
	FindAccounts() ([]payloads.Account, error)
	AddAccount(input inputs.Account) ([]payloads.Account, error)
	FindAccount(id uuid.UUID) ([]payloads.Account, error)
	RemoveAccount(id uuid.UUID) ([]payloads.Account, error)
	UpdateAccount(id uuid.UUID, input inputs.Account) ([]payloads.Account, error)
}

type accountService struct{}

func (*accountService) FindAccounts() ([]payloads.Account, error) {
	payloads := []payloads.Account{}
	buffer := bytes.Buffer{}
	query := map[string]interface{}{}

	if err := json.NewEncoder(&buffer).Encode(query); err != nil {
		message := fmt.Errorf("error while encoding from json to buffer, %s", err.Error())
		extensions.Error(message.Error())
		return nil, message
	}

	entites, err := repositories.AccountRepository.FindAccounts(buffer)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&payloads, &entites)
	if err != nil {
		err = fmt.Errorf("error while copying from entity to payload, %s", err.Error())
		extensions.Error(err.Error())
		return nil, err
	}

	extensions.Info("done")
	return payloads, err
}

func (*accountService) AddAccount(input inputs.Account) ([]payloads.Account, error) {
	payload := payloads.Account{}
	entity := entities.NewAccount()

	err := copier.Copy(&entity, &input)
	if err != nil {
		err = fmt.Errorf("error while copying from input to entity, %s", err.Error())
		extensions.Error(err.Error())
		return nil, err
	}

	body, err := json.Marshal(entity)
	if err != nil {
		err = fmt.Errorf("error while marshal post to json, %s", err.Error())
		extensions.Error(err.Error())
		return nil, err
	}

	err = repositories.AccountRepository.AddAccount(entity.Id, body)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&payload, &entity)
	if err != nil {
		err = fmt.Errorf("error while copying from entity to payload, %s", err.Error())
		extensions.Error(err.Error())
		return nil, err
	}

	extensions.Info("done")
	return []payloads.Account{payload}, nil
}

func (*accountService) FindAccount(id uuid.UUID) ([]payloads.Account, error) {
	payload := payloads.Account{}

	entity, err := repositories.AccountRepository.FindAccount(id)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		err = fmt.Errorf("account not found")
		return nil, err
	}

	err = copier.Copy(&payload, &entity)
	if err != nil {
		err = fmt.Errorf("error while copying to payload from entity, %s", err.Error())
		extensions.Error(err.Error())
		return nil, err
	}

	extensions.Info("done")
	return []payloads.Account{payload}, nil
}

func (*accountService) RemoveAccount(id uuid.UUID) ([]payloads.Account, error) {
	payload := payloads.Account{}

	entity, err := repositories.AccountRepository.FindAccount(id)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		err = fmt.Errorf("account not found")
		return nil, err
	}

	err = repositories.AccountRepository.RemoveAccount(id)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&payload, &entity)
	if err != nil {
		err = fmt.Errorf("error while copying to payload from entity, %s", err.Error())
		extensions.Error(err.Error())
		return nil, err
	}

	extensions.Info("done")
	return []payloads.Account{payload}, nil
}

func (*accountService) UpdateAccount(id uuid.UUID, input inputs.Account) ([]payloads.Account, error) {
	payload := payloads.Account{}
	updated := time.Now()

	entity, err := repositories.AccountRepository.FindAccount(id)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		err = fmt.Errorf("account not found")
		return nil, err
	}
	entity.Updated = &updated

	err = copier.Copy(&entity, &input)
	if err != nil {
		err = fmt.Errorf("error while copying to entity from input, %s", err.Error())
		extensions.Error(err.Error())
		return nil, err
	}

	body, err := json.Marshal(entity)
	if err != nil {
		err = fmt.Errorf("error while marshal post to json, %s", err.Error())
		extensions.Error(err.Error())
		return nil, err
	}

	err = repositories.AccountRepository.UpdateAccount(id, body)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&payload, &entity)
	if err != nil {
		err = fmt.Errorf("error while copying to payload from entity, %s", err.Error())
		extensions.Error(err.Error())
		return nil, err
	}

	return []payloads.Account{payload}, nil
}
