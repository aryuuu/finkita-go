package service

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/aryuuu/finkita/domain"
	"github.com/aryuuu/finkita/internal/configs"
)

// TODO: add repos
type accountService struct {
	accountRepo domain.IAccountRepository
}

func NewAccountService(ar domain.IAccountRepository) domain.IAccountService {
	return &accountService{
		accountRepo: ar,
	}
}

func (s *accountService) AddAccount(ctx context.Context, account *domain.Account) error {
	var err error
	account.Password, err = encryptPassword(account.Password)
	if err != nil {
		return fmt.Errorf("accountService.AddAccount: failed to encrypt account password: %v", err)
	}

	return s.accountRepo.AddAccount(ctx, account)
}

func (s *accountService) GetAccounts(ctx context.Context) ([]domain.Account, error) {
	accounts, err := s.accountRepo.GetAccounts(ctx, map[string]interface{}{})
	if err != nil {
		return accounts, err
	}

	return accounts, nil
}

func (s *accountService) GetAccountsByEmail(ctx context.Context, email string) ([]domain.Account, error) {
	filter := map[string]interface{}{
		"email": email,
	}

	accounts, err := s.accountRepo.GetAccounts(ctx, filter)
	if err != nil {
		return accounts, err
	}

	return accounts, nil
}

func (s *accountService) GetAccountsWithPassword(ctx context.Context) ([]domain.Account, error) {
	accounts, err := s.accountRepo.GetAccountsWithPassword(ctx)
	if err != nil {
		return accounts, err
	}

	for i := range accounts {
		decryptedPassword, err := decryptPassword(accounts[i].Password)
		if err != nil {
			log.Printf("failed to decrypt password for account: %+v, error: %v", accounts[i], err)
		}
		accounts[i].Password = decryptedPassword
	}

	return accounts, nil
}

func (s *accountService) GetAccountByID(ctx context.Context, id string) (domain.Account, error) {
	return s.accountRepo.GetAccountByID(ctx, id)
}

func (s *accountService) UpdateAccountByID(ctx context.Context, id string, account *domain.Account) error {
	return s.accountRepo.UpdateAccountByID(ctx, id, account)
}

func (s *accountService) DeleteAccount(ctx context.Context, id string) error {
	return s.accountRepo.Delete(ctx, id)
}

func encryptPassword(password string) (string, error) {
	passwordByte := []byte(password)
	key := []byte(configs.Account.EncKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("error init aesgcm: %v", err)
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	rand.Read(nonce)

	ciphertext := aesgcm.Seal(nonce, nonce, passwordByte, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func decryptPassword(encPassword string) (string, error) {
	if len(encPassword) == 0 {
		return encPassword, nil
	}

	encPasswordBytes, err := hex.DecodeString(encPassword)
	if err != nil {
		log.Printf("failed to decode hex string of encrypted password, error: %v", err)
		return "", err
	}
	block, err := aes.NewCipher([]byte(configs.Account.EncKey))
	if err != nil {
		log.Printf("failed to create new cipher: %v", err)
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("failed to create new aesgcm cipher: %v", err)
		return "", err
	}

	nonceSize := aesgcm.NonceSize()

	_nonce, _cipher := encPasswordBytes[:nonceSize], encPasswordBytes[nonceSize:]

	plaintext, err := aesgcm.Open(nil, _nonce, _cipher, nil)
	if err != nil {
		log.Printf("failed to open aesgcm cipher: %v", err)
		return "", err
	}

	return string(plaintext), nil
}
