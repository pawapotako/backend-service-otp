package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"m2ex-otp-service/internal/model"
	"m2ex-otp-service/internal/util"
	mocks "m2ex-otp-service/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateOtp(t *testing.T) {

	t.Run("CreateOtp Success", func(t *testing.T) {

		// arrange
		requestData := model.DefaultPayload[model.NewOtpRequest]{
			Data: model.NewOtpRequest{
				MobileNumber: "0987654321",
			},
		}
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			panic(err)
		}
		requestBody := bytes.NewBuffer(jsonData)

		service := mocks.NewOtpServiceMock()
		service.On("CreateOtp", mock.AnythingOfType("model.NewOtpRequest")).Return(&model.NewOtpResponse{ReferenceCode: "a1b2c3d4e5f6"}, nil)

		validator := validator.New()

		handler := otpHandler{service, validator}

		e := echo.New()
		v1 := e.Group("/v1")
		v1.POST("/otps", handler.createOtp)

		request := httptest.NewRequest("POST", "/v1/otps", requestBody)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept-Language", "en_US")
		recorder := httptest.NewRecorder()

		context := e.NewContext(request, recorder)

		// act
		handler.createOtp(context)

		// assert
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("CreateOtp Error server error", func(t *testing.T) {

		// arrange
		requestData := model.DefaultPayload[model.NewOtpRequest]{
			Data: model.NewOtpRequest{
				MobileNumber: "0987654321",
			},
		}
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			panic(err)
		}
		requestBody := bytes.NewBuffer(jsonData)

		service := mocks.NewOtpServiceMock()
		service.On("CreateOtp", mock.AnythingOfType("model.NewOtpRequest")).Return(&model.NewOtpResponse{}, errors.New(""))

		validator := validator.New()

		util.InitLogger()

		handler := otpHandler{service, validator}

		e := echo.New()
		v1 := e.Group("/v1")
		v1.POST("/otps", handler.createOtp)

		request := httptest.NewRequest("POST", "/v1/otps", requestBody)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept-Language", "en_US")
		recorder := httptest.NewRecorder()

		context := e.NewContext(request, recorder)

		// act
		handler.createOtp(context)

		// assert
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("CreateOtp Error bind request", func(t *testing.T) {

		// arrange
		requestData := ""
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			panic(err)
		}
		requestBody := bytes.NewBuffer(jsonData)

		service := mocks.NewOtpServiceMock()
		service.On("CreateOtp", mock.AnythingOfType("model.NewOtpRequest")).Return(&model.NewOtpResponse{ReferenceCode: "a1b2c3d4e5f6"}, nil)

		validator := validator.New()

		util.InitLogger()

		handler := otpHandler{service, validator}

		e := echo.New()
		v1 := e.Group("/v1")
		v1.POST("/otps", handler.createOtp)

		request := httptest.NewRequest("POST", "/v1/otps", requestBody)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept-Language", "en_US")
		recorder := httptest.NewRecorder()

		context := e.NewContext(request, recorder)

		// act
		handler.createOtp(context)

		// assert
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		service.AssertNotCalled(t, "CreateOtp")
	})

	t.Run("CreateOtp Error validate struct", func(t *testing.T) {

		// arrange
		requestData := model.DefaultPayload[model.NewOtpRequest]{
			Data: model.NewOtpRequest{},
		}
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			panic(err)
		}
		requestBody := bytes.NewBuffer(jsonData)

		service := mocks.NewOtpServiceMock()
		service.On("CreateOtp", mock.AnythingOfType("model.NewOtpRequest")).Return(&model.NewOtpResponse{ReferenceCode: "a1b2c3d4e5f6"}, nil)

		validator := validator.New()

		util.InitLogger()

		handler := otpHandler{service, validator}

		e := echo.New()
		v1 := e.Group("/v1")
		v1.POST("/otps", handler.createOtp)

		request := httptest.NewRequest("POST", "/v1/otps", requestBody)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept-Language", "en_US")
		recorder := httptest.NewRecorder()

		context := e.NewContext(request, recorder)

		// act
		handler.createOtp(context)

		// assert
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		service.AssertNotCalled(t, "CreateOtp")
	})
}

func TestValidateOtp(t *testing.T) {

	t.Run("ValidateOtp Success", func(t *testing.T) {

		// arrange
		requestData := model.DefaultPayload[model.ValidateOtpRequest]{
			Data: model.ValidateOtpRequest{
				ReferenceCode: "a1b2c3d4e5f6",
				Otp:           "123456",
			},
		}
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			panic(err)
		}
		requestBody := bytes.NewBuffer(jsonData)

		service := mocks.NewOtpServiceMock()
		service.On("ValidateOtp", mock.AnythingOfType("model.ValidateOtpRequest")).Return(&model.ValidateOtpResponse{IsValidate: true}, nil)

		validator := validator.New()

		handler := otpHandler{service, validator}

		e := echo.New()
		v1 := e.Group("/v1")
		v1.POST("/otps/validate-otp", handler.validateOtp)

		request := httptest.NewRequest("POST", "/v1/otps/validate-otp", requestBody)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept-Language", "en_US")
		recorder := httptest.NewRecorder()

		context := e.NewContext(request, recorder)

		// act
		handler.validateOtp(context)

		// assert
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("ValidateOtp Error server error", func(t *testing.T) {

		t.Run("ValidateOtp Error record not found", func(t *testing.T) {

			// arrange
			requestData := model.DefaultPayload[model.ValidateOtpRequest]{
				Data: model.ValidateOtpRequest{
					ReferenceCode: "a1b2c3d4e5f6",
					Otp:           "654321",
				},
			}
			jsonData, err := json.Marshal(requestData)
			if err != nil {
				panic(err)
			}
			requestBody := bytes.NewBuffer(jsonData)

			service := mocks.NewOtpServiceMock()
			service.On("ValidateOtp", mock.AnythingOfType("model.ValidateOtpRequest")).Return(&model.ValidateOtpResponse{}, errors.New("record not found"))

			validator := validator.New()

			handler := otpHandler{service, validator}

			e := echo.New()
			v1 := e.Group("/v1")
			v1.POST("/otps/validate-otp", handler.validateOtp)

			request := httptest.NewRequest("POST", "/v1/otps/validate-otp", requestBody)
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("Accept-Language", "en_US")
			recorder := httptest.NewRecorder()

			context := e.NewContext(request, recorder)

			// act
			handler.validateOtp(context)

			// assert
			assert.Equal(t, http.StatusBadRequest, recorder.Code)
		})

		t.Run("ValidateOtp Error otp is expired", func(t *testing.T) {

			// arrange
			requestData := model.DefaultPayload[model.ValidateOtpRequest]{
				Data: model.ValidateOtpRequest{
					ReferenceCode: "a1b2c3d4e5f6",
					Otp:           "654321",
				},
			}
			jsonData, err := json.Marshal(requestData)
			if err != nil {
				panic(err)
			}
			requestBody := bytes.NewBuffer(jsonData)

			service := mocks.NewOtpServiceMock()
			service.On("ValidateOtp", mock.AnythingOfType("model.ValidateOtpRequest")).Return(&model.ValidateOtpResponse{}, errors.New("otp is expired"))

			validator := validator.New()

			handler := otpHandler{service, validator}

			e := echo.New()
			v1 := e.Group("/v1")
			v1.POST("/otps/validate-otp", handler.validateOtp)

			request := httptest.NewRequest("POST", "/v1/otps/validate-otp", requestBody)
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("Accept-Language", "en_US")
			recorder := httptest.NewRecorder()

			context := e.NewContext(request, recorder)

			// act
			handler.validateOtp(context)

			// assert
			assert.Equal(t, http.StatusBadRequest, recorder.Code)
		})

		t.Run("ValidateOtp Error anything else", func(t *testing.T) {

			// arrange
			requestData := model.DefaultPayload[model.ValidateOtpRequest]{
				Data: model.ValidateOtpRequest{
					ReferenceCode: "a1b2c3d4e5f6",
					Otp:           "654321",
				},
			}
			jsonData, err := json.Marshal(requestData)
			if err != nil {
				panic(err)
			}
			requestBody := bytes.NewBuffer(jsonData)

			service := mocks.NewOtpServiceMock()
			service.On("ValidateOtp", mock.AnythingOfType("model.ValidateOtpRequest")).Return(&model.ValidateOtpResponse{}, errors.New(""))

			validator := validator.New()

			handler := otpHandler{service, validator}

			e := echo.New()
			v1 := e.Group("/v1")
			v1.POST("/otps/validate-otp", handler.validateOtp)

			request := httptest.NewRequest("POST", "/v1/otps/validate-otp", requestBody)
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("Accept-Language", "en_US")
			recorder := httptest.NewRecorder()

			context := e.NewContext(request, recorder)

			// act
			handler.validateOtp(context)

			// assert
			assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		})

	})

	t.Run("ValidateOtp Error bind request", func(t *testing.T) {

		// arrange
		requestData := ""
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			panic(err)
		}
		requestBody := bytes.NewBuffer(jsonData)

		service := mocks.NewOtpServiceMock()
		service.On("ValidateOtp", mock.AnythingOfType("model.ValidateOtpRequest")).Return(&model.ValidateOtpResponse{IsValidate: true}, nil)

		validator := validator.New()

		handler := otpHandler{service, validator}

		e := echo.New()
		v1 := e.Group("/v1")
		v1.POST("/otps/validate-otp", handler.validateOtp)

		request := httptest.NewRequest("POST", "/v1/otps/validate-otp", requestBody)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept-Language", "en_US")
		recorder := httptest.NewRecorder()

		context := e.NewContext(request, recorder)

		// act
		handler.validateOtp(context)

		// assert
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		service.AssertNotCalled(t, "ValidateOtp")
	})

	t.Run("ValidateOtp Error validate struct", func(t *testing.T) {

		// arrange
		requestData := model.DefaultPayload[model.ValidateOtpRequest]{
			Data: model.ValidateOtpRequest{},
		}
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			panic(err)
		}
		requestBody := bytes.NewBuffer(jsonData)

		service := mocks.NewOtpServiceMock()
		service.On("ValidateOtp", mock.AnythingOfType("model.ValidateOtpRequest")).Return(&model.ValidateOtpResponse{IsValidate: true}, nil)

		validator := validator.New()

		handler := otpHandler{service, validator}

		e := echo.New()
		v1 := e.Group("/v1")
		v1.POST("/otps/validate-otp", handler.validateOtp)

		request := httptest.NewRequest("POST", "/v1/otps/validate-otp", requestBody)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept-Language", "en_US")
		recorder := httptest.NewRecorder()

		context := e.NewContext(request, recorder)

		// act
		handler.validateOtp(context)

		// assert
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		service.AssertNotCalled(t, "ValidateOtp")
	})

}
