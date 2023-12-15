package model

import (
	"time"
)

type OtpModel struct {
	Id            uint      `gorm:"column:id;primaryKey;not null:true"`
	ReferenceCode string    `gorm:"column:reference_code;type:varchar(300);not null:true"`
	Otp           string    `gorm:"column:otp;type:varchar(300);not null:true"`
	ExpiredAt     time.Time `gorm:"column:expired_at;not null:true"`
	MobileNumber  string    `gorm:"column:mobile_number;type:varchar(300);not null:true"`
	CreatedAt     time.Time `gorm:"column:created_at;not null:true"`
}
