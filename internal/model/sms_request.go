package model

import "time"

type SmsRequest struct {
	To                []SmsTo    `json:"to"`
	Message           string     `json:"message" validate:"required"`
	ScheduledDelivery *time.Time `json:"scheduledDelivery"`
}

type SmsTo struct {
	MobileNumber string `json:"mobileNumber" validate:"required"`
}
