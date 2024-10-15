package service

import (
	"github.com/HIUNCY/rest-api-go/model"
	"github.com/HIUNCY/rest-api-go/repository"
)

type DepositService interface {
	CreateDeposit(depo *model.Deposit) error
	GetDepositByNik(nik string) (*model.Deposit, error)
	UpdateDeposit(depo *model.Deposit) error
	DeleteDeposit(nik string) error
}

type depositService struct {
	depoRepo repository.DepositRepository
}

// CreateDeposit implements DepositService.
func (s *depositService) CreateDeposit(depo *model.Deposit) error {
	return s.depoRepo.CreateDeposit(depo)
}

func (s *depositService) DeleteDeposit(nik string) error {
	return s.depoRepo.DeleteDeposit(nik)
}

func (s *depositService) GetDepositByNik(nik string) (*model.Deposit, error) {
	depo, err := s.depoRepo.GetDepositByNik(nik)
	if err != nil {
		return nil, err
	}
	return depo, err
}

func (s *depositService) UpdateDeposit(depo *model.Deposit) error {
	return s.depoRepo.UpdateDeposit(depo)
}

func NewDepositService(depoRepo repository.DepositRepository) DepositService {
	return &depositService{depoRepo: depoRepo}
}
