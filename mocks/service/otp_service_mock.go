package mocks

import (
	"m2ex-otp-service/internal/model"

	"github.com/stretchr/testify/mock"
)

type otpServiceMock struct {
	mock.Mock
}

func NewOtpServiceMock() *otpServiceMock {
	return &otpServiceMock{}
}

func (m *otpServiceMock) CreateOtp(request model.NewOtpRequest) (*model.NewOtpResponse, error) {
	args := m.Called(request)
	return args.Get(0).(*model.NewOtpResponse), args.Error(1)
}

func (m *otpServiceMock) ValidateOtp(request model.ValidateOtpRequest) (*model.ValidateOtpResponse, error) {
	args := m.Called(request)
	return args.Get(0).(*model.ValidateOtpResponse), args.Error(1)
}
