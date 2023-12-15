package service

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"m2ex-otp-service/internal/model"
	"m2ex-otp-service/internal/repository"
	"m2ex-otp-service/internal/util"
	"math/rand"
	"strconv"
	"time"
)

type OtpService interface {
	CreateOtp(request model.NewOtpRequest) (*model.NewOtpResponse, error)
	ValidateOtp(request model.ValidateOtpRequest) (*model.ValidateOtpResponse, error)
}

type otpService struct {
	otpRepo    repository.OtpRepository
	pluginRepo repository.PluginRepository
}

func NewOtpService(otpRepo repository.OtpRepository, pluginRepo repository.PluginRepository) OtpService {
	return otpService{
		otpRepo:    otpRepo,
		pluginRepo: pluginRepo,
	}
}

func (s otpService) CreateOtp(request model.NewOtpRequest) (*model.NewOtpResponse, error) {

	config, _ := util.LoadConfig()
	location, _ := time.LoadLocation("Asia/Bangkok")
	var nowTime = time.Now().In(location)

	isMobileNumber := util.IsMobileNumber(request.MobileNumber)
	if !isMobileNumber {
		return nil, errors.New("it is not a mobile number format")
	}

	bytes := make([]byte, 6)
	rand.Read(bytes)
	refCode := hex.EncodeToString(bytes)
	rand.Seed(time.Now().UnixNano())
	otp := ""
	for i := 0; i < 6; i++ {
		otp += strconv.Itoa(rand.Intn(10))
	}

	expiredAt := time.Now().Add(time.Minute * 2)

	otpCreate := model.OtpModel{
		ReferenceCode: refCode,
		Otp:           otp,
		ExpiredAt:     expiredAt,
		MobileNumber:  request.MobileNumber,
		CreatedAt:     nowTime,
	}

	entity, err := s.otpRepo.Create(otpCreate)
	if err != nil {
		return nil, err
	}

	smsTo := []model.SmsTo{}
	smsTo = append(smsTo, model.SmsTo{
		MobileNumber: request.MobileNumber,
	})

	message := fmt.Sprintf("To validate your mobile number, please use the following One Time Password (OTP): %s", entity.Otp)
	message += fmt.Sprintf("(Ref: %s) Do not share this OTP with anyone.", entity.ReferenceCode)

	smsRequestData := model.SmsRequest{
		To:      smsTo,
		Message: base64.StdEncoding.EncodeToString([]byte(message)),
	}

	// send kafka to sms service
	go s.pluginRepo.EventProducer(config.Kafka.Topic, smsRequestData)

	response := model.NewOtpResponse{
		ReferenceCode: entity.ReferenceCode,
	}

	return &response, nil
}

func (s otpService) ValidateOtp(request model.ValidateOtpRequest) (*model.ValidateOtpResponse, error) {

	otpEntity, err := s.otpRepo.GetOtpByReferenceCode(request.ReferenceCode, request.Otp)
	if err != nil {
		return nil, err
	}

	if otpEntity.ExpiredAt.Unix() < time.Now().Unix() {
		return nil, fmt.Errorf("otp is expired")
	}

	response := model.ValidateOtpResponse{
		IsValidate: true,
	}

	return &response, nil
}
