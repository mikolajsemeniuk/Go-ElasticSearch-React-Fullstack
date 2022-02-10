package services

import (
	"github.com/jinzhu/copier"
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
	entites, err := repositories.AccountRepository.FindAccounts()

	copier.Copy(&payloads, &entites)

	return payloads, err
}
