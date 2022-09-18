package domain

import (
	"context"
)

type Account struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Bank          string `json:"bank"`
	AccountNumber string `json:"account_number"`
	Password      string `json:"password,omitempty"`
	CreatedAt     int    `json:"created_at,omitempty"`
	UpdatedAt     int    `json:"updated_at,omitempty"`
	DeletedAt     int    `json:"deleted_at,omitempty"`
}

type IAccountService interface {
	AddAccount(ctx context.Context, account *Account) error
	GetAccounts(ctx context.Context) ([]Account, error)
	GetAccountByID(ctx context.Context, id string) (Account, error)
	UpdateAccountByID(ctx context.Context, id string, account *Account) error
	DeleteAccount(ctx context.Context, id string) error
}

type IAccountRepository interface {
	AddAccount(ctx context.Context, account *Account) error
	GetAccounts(ctx context.Context) ([]Account, error)
	GetAccountByID(ctx context.Context, id string) (Account, error)
	UpdateAccountByID(ctx context.Context, id string, account *Account) error
	Delete(ctx context.Context, id string) error
}
