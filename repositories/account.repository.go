package repositories

import "github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/entities"

var AccountRepository IAccountRepository = &accountRepository{}

type IAccountRepository interface {
	FindAccounts() ([]entities.Account, error)
}

type accountRepository struct{}

func (*accountRepository) FindAccounts() ([]entities.Account, error) {
	return nil, nil
}
