package model

// NewOtpRequest godoc
// @Summary New OTP request structure
// @Description Detailed description for new OTP request
// @Model NewOtpRequest
type NewOtpRequest struct {
	MobileNumber string `json:"mobileNumber" validate:"required"`
}

// ValidateOtpRequest godoc
// @Summary Validate OTP request structure
// @Description Detailed description for Validate OTP request
// @Model ValidateOtpRequest
type ValidateOtpRequest struct {
	ReferenceCode string `json:"referenceCode" validate:"required"`
	Otp           string `json:"otp" validate:"required"`
}
