package repositories

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/aryuuu/finkita/domain"
)

type mutationRepo struct {
	dbCon *sql.DB
}

func NewMutationRepo(dbCon *sql.DB) domain.IMutationRepository {
	return &mutationRepo{
		dbCon: dbCon,
	}
}

func (mr *mutationRepo) BulkAddMutation(ctx context.Context, mutations []domain.Mutation) error {
	query := sq.Insert("mutations").
		Columns(
			"id",
			"account_id",
			"email",
			"date",
			"description",
			"type",
			"amount",
			"balance",
			"currency",
			"created_at",
			"updated_at",
		)

	for _, mutation := range mutations {
		query = query.Values(
			sq.Expr("UUID_GENERATE_V4()"),
			mutation.AccountID,
			mutation.Email,
			mutation.Date,
			mutation.Description,
			mutation.Type,
			mutation.Amount,
			mutation.Balance,
			mutation.Currency,
			sq.Expr("NOW()"),
			sq.Expr("NOW()"),
		)
	}

	query.PlaceholderFormat(sq.Dollar)
	_, err := query.RunWith(mr.dbCon).ExecContext(ctx)
	if err != nil {
		log.Printf("error bulk inserting mutations: %v", err)
		return err
	}

	return nil
}

func (mr *mutationRepo) GetMutations(ctx context.Context, filter map[string]interface{}) ([]domain.Mutation, error) {
	query := sq.Select(
		"id",
		"account_id",
		"email",
		"date",
		"description",
		"type",
		"amount",
		"balance",
		"currency",
		"created_at",
		"updated_at",
	).
		From("mutations").
		Where(sq.Eq{"deleted_at": nil})

	rows, err := query.RunWith(mr.dbCon).QueryContext(ctx)
	if err != nil {
		log.Printf("error running query on get all mutations: %v", err)
		return []domain.Mutation{}, err
	}

	mutations := []domain.Mutation{}
	mutation := domain.Mutation{}
	for rows.Next() {
		err = rows.Scan(
			&mutation.ID,
			&mutation.AccountID,
			&mutation.Email,
			&mutation.Date,
			&mutation.Description,
			&mutation.Type,
			&mutation.Amount,
			&mutation.Balance,
			&mutation.Currency,
			&mutation.CreatedAt,
			&mutation.UpdatedAt,
		)
		if err != nil {
			log.Printf("error scanning mutation row: %v", err)
			continue
		}
		mutations = append(mutations, mutation)
	}

	return mutations, nil
}

func (mr *mutationRepo) GetMutationByID(ctx context.Context, id string) (domain.Mutation, error) {
	query := sq.Select(
		"id",
		"account_id",
		"email",
		"date",
		"description",
		"type",
		"amount",
		"balance",
		"currency",
		"created_at",
		"updated_at",
	).From("mutations").Where(sq.Eq{"id": id}).Where(sq.Eq{"deleted_at": nil})

	mutation := domain.Mutation{}
	err := query.RunWith(mr.dbCon).
		QueryRowContext(ctx).
		Scan(
			&mutation.ID,
			&mutation.AccountID,
			&mutation.Email,
			&mutation.Date,
			&mutation.Description,
			&mutation.Type,
			&mutation.Amount,
			&mutation.Balance,
			&mutation.Currency,
			&mutation.CreatedAt,
			&mutation.UpdatedAt,
		)

	return mutation, err
}
