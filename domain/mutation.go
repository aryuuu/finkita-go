package domain

import (
	"context"
)

type Mutation struct {
	ID          string `json:"id"`
	AccountID   string `json:"account_id"`
	Email       string `json:"email"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Amount      int    `json:"amount"`
	Balance     int    `json:"balance"`
	Currency    string `json:"currency"`
	CreatedAt   int    `json:"created_at,omitempty"`
	UpdatedAt   int    `json:"updated_at,omitempty"`
	DeletedAt   int    `json:"deleted_at,omitempty"`
}

type IMutationService interface {
	AddMutation(ctx context.Context, mutation *Mutation) error
	GetMutations(ctx context.Context) ([]Mutation, error)
	GetMutationByID(ctx context.Context, id string) (Mutation, error)
}

type IMutationRepository interface {
	BulkAddMutation(ctx context.Context, mutations []Mutation) error
	GetMutations(ctx context.Context) ([]Mutation, error)
	GetMutationByID(ctx context.Context, id string) (Mutation, error)
}
