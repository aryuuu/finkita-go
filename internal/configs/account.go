package configs

import "os"

type account struct {
	EncKey string
}

func setupAccount() *account {
	return &account{
		EncKey: os.Getenv("ACCOUNT_ENC_KEY"),
	}
}
