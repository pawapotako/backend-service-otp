package model

type Tabler interface {
	TableName() string
}

func (OtpModel) TableName() string {
	return "otp"
}
