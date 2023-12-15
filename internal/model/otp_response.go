package model

// NewOtpResponse godoc
// @Summary New OTP response structure
// @Description Detailed description for new OTP response
// @Model NewOtpResponse
type NewOtpResponse struct {
	ReferenceCode string `json:"referenceCode"`
}

// ValidateOtpResponse godoc
// @Summary Validate OTP response structure
// @Description Detailed description for Validate OTP response
// @Model ValidateOtpResponse
type ValidateOtpResponse struct {
	IsValidate bool `json:"isValidate"`
}
