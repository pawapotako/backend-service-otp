package handler

import (
	"m2ex-otp-service/internal/model"
	"m2ex-otp-service/internal/repository"
	"m2ex-otp-service/internal/service"
	"m2ex-otp-service/internal/util"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type otpHandler struct {
	service  service.OtpService
	validate *validator.Validate
}

func NewOtpHandler(db *gorm.DB, e *echo.Echo, validator *validator.Validate) {
	otpRepo := repository.NewOtpRepository(db)
	pluginRepo := repository.NewPluginRepository()
	service := service.NewOtpService(otpRepo, pluginRepo)
	handler := otpHandler{service, validator}

	v1 := e.Group("/v1")
	v1.POST("/otps", handler.createOtp)
	v1.POST("/otps/validate-otp", handler.validateOtp)
}

// Create Otp godoc
// @Summary Create Otp with a POST request
// @Description Sends a POST request to create Otp
// @Tags Otp
// @Accept  json
// @Produce  json
// @Param otpRequest body model.DefaultPayload[model.NewOtpRequest] true "OTP Request"
// @Success 200 {object} model.DefaultPayload[model.NewOtpResponse]
// @Failure 400  {object}  util.AppErrors
// @Failure 422  {object}  util.AppErrors
// @Failure 500  {object}  util.AppErrors
// @Router /v1/otps [post]
func (h otpHandler) createOtp(ctx echo.Context) error {

	request := model.DefaultPayload[model.NewOtpRequest]{}
	if err := ctx.Bind(&request); err != nil {
		util.Logger.Error("bind request", zap.Error(err))

		errs := util.AppErrors{}
		errs.Add(util.NewBadRequestError())
		return errorHandler(ctx, errs)
	}

	if err := h.validate.Struct(request); err != nil {
		util.Logger.Error("validate struct", zap.Error(err))

		errs := util.AppErrors{}
		errs.Add(util.NewValidationError(err.Error()))
		return errorHandler(ctx, errs)
	}

	response, err := h.service.CreateOtp(request.Data)
	if err != nil {
		util.Logger.Error("server error", zap.Error(err))

		errs := util.AppErrors{}
		errs.Add(util.NewCustomError(err.Error()))
		return errorHandler(ctx, errs)
	}
	return ctx.JSON(http.StatusOK, model.DefaultPayload[model.NewOtpResponse]{Data: *response})
}

// Validate Otp godoc
// @Summary Validate Otp with a POST request
// @Description Sends a POST request to Validate Otp
// @Tags Otp
// @Accept  json
// @Produce  json
// @Param signInRequest body model.DefaultPayload[model.ValidateOtpRequest] true "Validate Otp Request"
// @Success 200 {object} model.DefaultPayload[model.ValidateOtpResponse]
// @Failure 400  {object}  util.AppErrors
// @Failure 401  {object}  util.AppErrors
// @Failure 422  {object}  util.AppErrors
// @Failure 500  {object}  util.AppErrors
// @Router /v1/otps/validate-otp [post]
func (h otpHandler) validateOtp(ctx echo.Context) error {

	request := model.DefaultPayload[model.ValidateOtpRequest]{}
	if err := ctx.Bind(&request); err != nil {
		util.Logger.Error("bind request", zap.Error(err))

		errs := util.AppErrors{}
		errs.Add(util.NewBadRequestError())
		return errorHandler(ctx, errs)
	}

	if err := h.validate.Struct(request); err != nil {
		util.Logger.Error("validate struct", zap.Error(err))

		errs := util.AppErrors{}
		errs.Add(util.NewValidationError(err.Error()))
		return errorHandler(ctx, errs)
	}

	response, err := h.service.ValidateOtp(request.Data)
	if err != nil {
		util.Logger.Error("server error", zap.Error(err))

		errs := util.AppErrors{}

		if err.Error() == "record not found" || err.Error() == "otp is expired" {
			errs.Add(util.NewCustomBadRequestError(err.Error()))
			return errorHandler(ctx, errs)
		} else {
			errs.Add(util.NewCustomError(err.Error()))
			return errorHandler(ctx, errs)
		}
	}

	return ctx.JSON(http.StatusOK, model.DefaultPayload[model.ValidateOtpResponse]{Data: *response})
}
