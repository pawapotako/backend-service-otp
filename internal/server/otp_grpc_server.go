package server

import (
	"context"
	"fmt"
	"m2ex-otp-service/internal/model"
	"m2ex-otp-service/internal/repository"
	"m2ex-otp-service/internal/service"
	"m2ex-otp-service/internal/util"
	"net"

	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
	"gorm.io/gorm"
)

type otpServer struct {
	service service.OtpService
}

func NewOtpGrpcServer(db *gorm.DB, config util.Config) {

	opts := make([]grpc.ServerOption, 0)
	s := grpc.NewServer(opts...)

	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "", config.Grpc.Port))
	if err != nil {
		util.Logger.Error("cannot create listener", zap.Error(err))
	}

	otpRepo := repository.NewOtpRepository(db)
	pluginRepo := repository.NewPluginRepository()
	service := service.NewOtpService(otpRepo, pluginRepo)

	RegisterOtpServer(s, otpServer{service})

	fmt.Printf("gRPC server started on port [::]:%v\n", config.Grpc.Port)
	err = s.Serve(listener)
	if err != nil {
		util.Logger.Error("cannot serve server", zap.Error(err))
	}

	defer s.Stop()
}

func (otpServer) mustEmbedUnimplementedOtpServer() {}

func (h otpServer) ValidateOtp(ctx context.Context, request *ValidateOtpRequest) (*ValidateOtpResponse, error) {

	fmt.Printf("request : %v\n", request)
	var response *ValidateOtpResponse

	validateOtpRequest := model.ValidateOtpRequest{
		ReferenceCode: request.ReferenceCode,
		Otp:           request.Otp,
	}

	_, err := h.service.ValidateOtp(validateOtpRequest)
	if err != nil {
		util.Logger.Error("validate otp fail", zap.Error(err))
		response = &ValidateOtpResponse{
			IsValidate: false,
		}
		return response, err
	}

	response = &ValidateOtpResponse{
		IsValidate: true,
	}
	return response, nil
}
