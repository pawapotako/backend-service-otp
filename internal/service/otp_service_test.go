package service

import (
	"encoding/hex"
	"errors"
	"m2ex-otp-service/internal/model"
	mocks "m2ex-otp-service/mocks/repository"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOtp(t *testing.T) {

	location, _ := time.LoadLocation("Asia/Bangkok")
	var nowTime = time.Now().In(location)
	var expiredAt = time.Now().Add(time.Minute * 2)

	bytes := make([]byte, 6)
	refCode := hex.EncodeToString(bytes)
	rand.Seed(time.Now().UnixNano())
	otp := ""
	for i := 0; i < 6; i++ {
		otp += strconv.Itoa(rand.Intn(10))
	}

	t.Run("CreateOtp Applied mobile number", func(t *testing.T) {

		// arrange
		otpRepo := mocks.NewOtpRepositoryMock()
		otpRepo.On("Create", mock.AnythingOfType("model.OtpModel")).Return(&model.OtpModel{
			Id:            1,
			ReferenceCode: refCode,
			Otp:           otp,
			ExpiredAt:     expiredAt,
			MobileNumber:  "0987654321",
			CreatedAt:     nowTime,
		}, nil)

		pluginRepo := mocks.NewPluginRepositoryMock()
		pluginRepo.On("EventProducer", mock.AnythingOfType("string"), mock.AnythingOfType("SmsRequest")).Return(nil)

		service := NewOtpService(otpRepo, pluginRepo)

		// act
		response, err := service.CreateOtp(model.NewOtpRequest{MobileNumber: "0987654321"})
		expected := &model.NewOtpResponse{ReferenceCode: refCode}

		// assert
		assert.Equal(t, expected, response)
		assert.Nil(t, err)
	})

	t.Run("CreateOtp Repository error", func(t *testing.T) {

		// arrange
		otpRepo := mocks.NewOtpRepositoryMock()
		otpRepo.On("Create", mock.AnythingOfType("model.OtpModel")).Return(&model.OtpModel{}, errors.New(""))

		pluginRepo := mocks.NewPluginRepositoryMock()
		pluginRepo.On("EventProducer", mock.AnythingOfType("string"), mock.AnythingOfType("SmsRequest")).Return(nil)

		service := NewOtpService(otpRepo, pluginRepo)

		// act
		response, err := service.CreateOtp(model.NewOtpRequest{MobileNumber: "0987654321"})

		// assert
		assert.Nil(t, response)
		assert.NotNil(t, err)
		otpRepo.AssertNotCalled(t, "EventProducer")
	})

	t.Run("CreateOtp Applied email", func(t *testing.T) {

		// arrange
		otpRepo := mocks.NewOtpRepositoryMock()
		otpRepo.On("Create", mock.AnythingOfType("model.OtpModel")).Return(&model.OtpModel{}, nil)

		pluginRepo := mocks.NewPluginRepositoryMock()
		pluginRepo.On("EventProducer", mock.AnythingOfType("string"), mock.AnythingOfType("SmsRequest")).Return(nil)

		service := NewOtpService(otpRepo, pluginRepo)

		// act
		response, err := service.CreateOtp(model.NewOtpRequest{MobileNumber: "pawapotako.p@gmail.com"})

		// assert
		assert.Nil(t, response)
		assert.EqualError(t, err, "it is not a mobile number format")
		otpRepo.AssertNotCalled(t, "Create")
		otpRepo.AssertNotCalled(t, "EventProducer")
	})
}

func TestValidateOtp(t *testing.T) {

	location, _ := time.LoadLocation("Asia/Bangkok")
	var nowTime = time.Now().In(location)

	t.Run("ValidateOtp Applied match otp", func(t *testing.T) {

		// arrange
		otpRepo := mocks.NewOtpRepositoryMock()
		otpRepo.On("GetOtpByReferenceCode", "a1b2c3d4e5f6", "123456").Return(&model.OtpModel{
			Id:            1,
			ReferenceCode: "a1b2c3d4e5f6",
			Otp:           "123456",
			ExpiredAt:     time.Now().Add(time.Minute * 2),
			MobileNumber:  "0987654321",
			CreatedAt:     nowTime,
		}, nil)

		pluginRepo := mocks.NewPluginRepositoryMock()

		service := NewOtpService(otpRepo, pluginRepo)

		// act
		response, err := service.ValidateOtp(model.ValidateOtpRequest{ReferenceCode: "a1b2c3d4e5f6", Otp: "123456"})
		expected := &model.ValidateOtpResponse{IsValidate: true}

		// assert
		assert.Equal(t, expected, response)
		assert.Nil(t, err)
	})

	t.Run("ValidateOtp Applied not match otp", func(t *testing.T) {

		// arrange
		otpRepo := mocks.NewOtpRepositoryMock()
		otpRepo.On("GetOtpByReferenceCode", "a1b2c3d4e5f6", "654321").Return(&model.OtpModel{}, errors.New(""))

		pluginRepo := mocks.NewPluginRepositoryMock()

		service := NewOtpService(otpRepo, pluginRepo)

		// act
		response, err := service.ValidateOtp(model.ValidateOtpRequest{ReferenceCode: "a1b2c3d4e5f6", Otp: "654321"})

		// assert
		assert.Nil(t, response)
		assert.NotNil(t, err)
	})

	t.Run("ValidateOtp Applied otp is expired", func(t *testing.T) {

		// arrange
		otpRepo := mocks.NewOtpRepositoryMock()
		otpRepo.On("GetOtpByReferenceCode", "a1b2c3d4e5f6", "123456").Return(&model.OtpModel{
			Id:            1,
			ReferenceCode: "a1b2c3d4e5f6",
			Otp:           "123456",
			ExpiredAt:     time.Now().Add(time.Minute * -2),
			MobileNumber:  "0987654321",
			CreatedAt:     nowTime,
		}, nil)

		pluginRepo := mocks.NewPluginRepositoryMock()

		service := NewOtpService(otpRepo, pluginRepo)

		// act
		response, err := service.ValidateOtp(model.ValidateOtpRequest{ReferenceCode: "a1b2c3d4e5f6", Otp: "123456"})

		// assert
		assert.Nil(t, response)
		assert.EqualError(t, err, "otp is expired")
	})
}
