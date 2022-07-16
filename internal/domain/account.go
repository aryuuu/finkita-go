package domain

import (
	"context"
)

type Account struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	Bank          string `json:"bank"`
	AccountNumber string `json:"accountNumber"`
	Password      string `json:"password"`
	CreatedAt     int    `json:"createdAt"`
	UpdatedAt     int    `json:"updatedAt"`
	DeletedAt     int    `json:"deletedAt"`
}

type IAccountService interface {
	AddAccount(ctx context.Context, account *Account) error
	GetAccounts(ctx context.Context) ([]Account, error)
	GetAccountByID(ctx context.Context, id string) (Account, error)
	UpdateAccount(ctx context.Context, id string, account Account) (Account, error)
	DeleteAccount(ctx context.Context, id string) (Account, error)
}

type IAccountRepository interface {
}
