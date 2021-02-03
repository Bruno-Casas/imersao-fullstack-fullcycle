package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	// TransactionPending represents the pending status
	TransactionPending string = "pending"
	// TransactionCompleted represents the completed status
	TransactionCompleted string = "completed"
	// TransactionError represents the error status
	TransactionError string = "error"
	// TransactionConfirmed represents the confirmed status for
	TransactionConfirmed string = "confirmed"
)

// TransactionRepositoryInterface is the contract to PixKey repository
type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

// Transactions represents a Transaction list
type Transactions struct {
	Transaction []Transaction
}

// Transaction represent the pix transaction in codePix project
type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"colunn:account_from_id;type:uuid;not null" valid:"notnull"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyToID        string   `gorm:"colunn:pix_key_id;type:uuid;not null" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

func (transaction *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)

	if transaction.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	if transaction.Status != TransactionPending && transaction.Status != TransactionCompleted && transaction.Status != TransactionError {
		return errors.New("invalid status for the tran saction")
	}

	if transaction.PixKeyTo.AccountID == transaction.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the sane")
	}

	if err != nil {
		return err
	}

	return nil
}

// NewTransaction create a Transaction instance
func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom:   accountFrom,
		AccountFromID: accountFrom.ID,
		Amount:        amount,
		PixKeyTo:      pixKeyTo,
		PixKeyToID:    pixKeyTo.ID,
		Status:        TransactionPending,
		Description:   description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// Complete change status of transaction to completed
func (transaction *Transaction) Complete() error {
	transaction.Status = TransactionCompleted
	transaction.UpdatedAt = time.Now()
	err := transaction.isValid()
	return err
}

// Confirm change status of transaction to comfirmed
func (transaction *Transaction) Confirm() error {
	transaction.Status = TransactionConfirmed
	transaction.UpdatedAt = time.Now()
	err := transaction.isValid()
	return err
}

// Cancel change status of transaction to error
func (transaction *Transaction) Cancel(description string) error {
	transaction.Status = TransactionError
	transaction.UpdatedAt = time.Now()
	transaction.CancelDescription = description
	err := transaction.isValid()
	return err
}
