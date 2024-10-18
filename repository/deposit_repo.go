package repository

import (
	"github.com/HIUNCY/rest-api-go/model"
	"gorm.io/gorm"
)

type DepositRepository interface {
	CreateDeposit(depo *model.Deposit) error
	GetDepositByNik(nik string) (*model.Deposit, error)
	UpdateDeposit(depo *model.Deposit) error
	DeleteDeposit(nik string) error
	GetDepositList() (*[]model.Deposit, error)
}

type depositRepository struct {
	db *gorm.DB
}

func (r *depositRepository) CreateDeposit(depo *model.Deposit) error {
	return r.db.Create(depo).Error
}

func (r *depositRepository) DeleteDeposit(nik string) error {
	return r.db.Where("nik = ?", nik).Delete(&model.Deposit{}).Error
}

func (r *depositRepository) GetDepositByNik(nik string) (*model.Deposit, error) {
	var depo model.Deposit
	err := r.db.Where("nik = ?", nik).First(&depo).Error
	return &depo, err
}

func (r *depositRepository) UpdateDeposit(depo *model.Deposit) error {
	return r.db.Save(depo).Error
}

func NewDepositRepository(db *gorm.DB) DepositRepository {
	return &depositRepository{db}
}

func (r *depositRepository) GetDepositList() (*[]model.Deposit, error) {
	var deposits []model.Deposit
	err := r.db.Find(&deposits).Error
	return &deposits, err
}
