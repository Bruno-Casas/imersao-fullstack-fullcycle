package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// Account represents a account in the bank
type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" gorm:"column:owner_name;type:varchar(255)" valid:"notnull"`
	BankID    string    `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Bank      *Bank     `valid:"-"`
	Number    string    `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	PixKeys   []*PixKey `gorm:"ForeignKey: AccountID" valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

// NewAccount create a Account instance
func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	account := Account{
		OwnerName: ownerName,
		BankID:    bank.ID,
		Bank:      bank,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}
