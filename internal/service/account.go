package service

import (
	"context"

	"github.com/aryuuu/finkita/internal/domain"
)

// TODO: add repos
type AccountService struct {
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s AccountService) AddAccount(ctx context.Context, account *domain.Account) error {
	return nil
}

func (s AccountService) GetAccounts(ctx context.Context) ([]domain.Account, error) {
	return []domain.Account{}, nil
}

func (s AccountService) GetAccountByID(ctx context.Context, id string) (domain.Account, error) {
	return domain.Account{}, nil
}

func (s AccountService) UpdateAccount(ctx context.Context, id string, account domain.Account) (domain.Account, error) {
	return domain.Account{}, nil
}

func (s AccountService) DeleteAccount(ctx context.Context, id string) (domain.Account, error) {
	return domain.Account{}, nil
}
