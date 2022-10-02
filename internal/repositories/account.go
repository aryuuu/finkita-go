package repositories

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/aryuuu/finkita/domain"
)

type accountRepo struct {
	dbCon *sql.DB
}

func NewAccountRepo(dbCon *sql.DB) domain.IAccountRepository {
	return &accountRepo{
		dbCon: dbCon,
	}
}

func (ar *accountRepo) AddAccount(ctx context.Context, account *domain.Account) error {
	query := sq.Insert("accounts").
		Columns(
			"id",
			"email",
			"bank",
			"user_id",
			"account_number",
			"password",
			"created_at",
			"updated_at",
		).
		Values(
			sq.Expr("UUID_GENERATE_V4()"),
			account.Email,
			account.Bank,
			account.UserID,
			account.AccountNumber,
			account.Password,
			sq.Expr("NOW()"),
			sq.Expr("NOW()"),
		).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar)

	err := query.RunWith(ar.dbCon).QueryRowContext(ctx).Scan(&account.ID)
	if err != nil {
		log.Printf("error creating new account: %v", err)
		return err
	}

	account.Password = ""

	return err
}

func (ar *accountRepo) GetAccounts(ctx context.Context, filter map[string]interface{}) ([]domain.Account, error) {
	query := sq.Select(
		"id",
		"email",
		"bank",
		"user_id",
		"account_number",
	).
		From("accounts").
		Where(sq.Eq{"deleted_at": nil})

	if email, ok := filter["email"]; ok {
		query = query.Where(sq.Eq{"email": email})
	}

	query = query.PlaceholderFormat(sq.Dollar)

	queryString, _, _ := query.ToSql()
	log.Printf("query: %s", queryString)

	rows, err := query.RunWith(ar.dbCon).QueryContext(ctx)
	if err != nil {
		log.Printf("error running query on get all accounts: %v", err)
		return []domain.Account{}, err
	}

	accounts := []domain.Account{}
	account := domain.Account{}
	for rows.Next() {
		err = rows.Scan(
			&account.ID,
			&account.Email,
			&account.Bank,
			&account.UserID,
			&account.AccountNumber,
		)
		if err != nil {
			log.Printf("error scanning account row: %v", err)
			// return accounts, err
			continue
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (ar *accountRepo) GetAccountsWithPassword(ctx context.Context) ([]domain.Account, error) {
	query := sq.Select(
		"id",
		"email",
		"bank",
		"user_id",
		"account_number",
		"password",
	).
		From("accounts").
		Where(sq.Eq{"deleted_at": nil}).
		PlaceholderFormat(sq.Dollar)

	queryString, _, _ := query.ToSql()
	log.Printf("query: %s", queryString)

	rows, err := query.RunWith(ar.dbCon).QueryContext(ctx)
	if err != nil {
		log.Printf("error running query on get all accounts: %v", err)
		return []domain.Account{}, err
	}

	accounts := []domain.Account{}
	account := domain.Account{}
	for rows.Next() {
		err = rows.Scan(
			&account.ID,
			&account.Email,
			&account.Bank,
			&account.UserID,
			&account.AccountNumber,
			&account.Password,
		)
		if err != nil {
			log.Printf("error scanning account row: %v", err)
			// return accounts, err
			continue
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (ar *accountRepo) GetAccountByID(ctx context.Context, id string) (domain.Account, error) {
	query := sq.Select(
		"id",
		"email",
		"bank",
		"user_id",
		"account_number",
		"created_at",
		"updated_at",
	).
		From("accounts").
		Where(sq.Eq{"id": id}).
		Where(sq.Eq{"deleted_at": nil}).
		PlaceholderFormat(sq.Dollar)

	account := domain.Account{}
	err := query.RunWith(ar.dbCon).
		QueryRowContext(ctx).
		Scan(
			&account.ID,
			&account.Email,
			&account.Bank,
			&account.UserID,
			&account.AccountNumber,
			&account.CreatedAt,
			&account.UpdatedAt,
		)

	return account, err
}

func (ar *accountRepo) UpdateAccountByID(ctx context.Context, id string, account *domain.Account) error {
	query := sq.Update("accounts")
	if account.Bank != "" {
		query = query.Set("bank", account.Bank)
	}
	if account.UserID != "" {
		query = query.Set("user_id", account.UserID)
	}
	if account.AccountNumber != "" {
		query = query.Set("account_number", account.AccountNumber)
	}
	if account.Password != "" {
		query = query.Set("password", account.Password)
	}

	query = query.Set("updated_at", sq.Expr("NOW()")).Where(sq.Eq{"id": id}).Where(sq.Eq{"deleted_at": nil}).PlaceholderFormat(sq.Dollar)
	err := query.RunWith(ar.dbCon).QueryRowContext(ctx).Scan(account)

	return err
}

func (ar *accountRepo) Delete(ctx context.Context, id string) error {
	query := sq.Update("accounts").
		Set("deleted_at", sq.Expr("NOW()")).
		Where(sq.Eq{"id": id}).
		Where(sq.Eq{"deleted_at": nil}).
		PlaceholderFormat(sq.Dollar)

	query.RunWith(ar.dbCon).ExecContext(ctx)
	return nil
}
