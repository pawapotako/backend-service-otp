package mocks

import (
	"m2ex-otp-service/internal/model"

	"github.com/stretchr/testify/mock"
)

type otpRepositoryMock struct {
	mock.Mock
}

func NewOtpRepositoryMock() *otpRepositoryMock {
	return &otpRepositoryMock{}
}

func (m *otpRepositoryMock) GetOtpByReferenceCode(referenceCode string, otp string) (*model.OtpModel, error) {
	args := m.Called(referenceCode, otp)
	return args.Get(0).(*model.OtpModel), args.Error(1)
}

func (m *otpRepositoryMock) Create(entity model.OtpModel) (*model.OtpModel, error) {
	args := m.Called(entity)
	return args.Get(0).(*model.OtpModel), args.Error(1)
}
