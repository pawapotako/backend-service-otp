package repository

import (
	"m2ex-otp-service/internal/model"

	"gorm.io/gorm"
)

type OtpRepository interface {
	GetOtpByReferenceCode(referenceCode string, otp string) (*model.OtpModel, error)
	Create(model.OtpModel) (*model.OtpModel, error)
}

type otpRepository struct {
	db *gorm.DB
}

func NewOtpRepository(db *gorm.DB) OtpRepository {
	return otpRepository{db}
}

func (r otpRepository) GetOtpByReferenceCode(referenceCode string, otp string) (*model.OtpModel, error) {
	entity := model.OtpModel{}
	if tx := r.db.Where("reference_code = ? AND otp = ?", referenceCode, otp).First(&entity); tx.Error != nil {
		return nil, tx.Error
	}

	return &entity, nil
}

func (r otpRepository) Create(entity model.OtpModel) (*model.OtpModel, error) {
	if tx := r.db.Create(&entity); tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
