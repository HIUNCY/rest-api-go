package service

import (
	"github.com/HIUNCY/rest-api-go/model"
	"github.com/HIUNCY/rest-api-go/repository"
)

type TransactionService interface {
	CreateTransaction(transaction *model.Transaction) error
	HistoryTransaction(nik string) ([]model.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepo}
}

func (s *transactionService) CreateTransaction(transaction *model.Transaction) error {
	return s.transactionRepo.CreateTransaction(transaction)
}

func (s *transactionService) HistoryTransaction(nik string) ([]model.Transaction, error) {
	return s.transactionRepo.HistoryTransaction(nik)
}
