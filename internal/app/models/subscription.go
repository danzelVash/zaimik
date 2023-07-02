package models

import (
	"time"
)

const (
	Active          = "active"
	Inactive        = "inactive"
	FirstPayAmount  = 1
	SecondPayAmount = 999
)

type Subscription struct {
	Id                       int        `json:"-" db:"id"`
	UserId                   int        `json:"user_id" binding:"required" db:"user_id"`
	LoanId                   int        `json:"-" db:"loan_id"`
	RequestDate              time.Time  `json:"request_date" binding:"required" db:"request_date"`
	FirstPayTime             *time.Time `json:"first_pay_time" binding:"required" db:"first_pay_time"`
	FirstPaySuccess          *bool      `json:"first_pay_success" binding:"required" db:"first_pay_success"`
	SecondPayAppointmentDate *time.Time `json:"second_pay_appointment_date" binding:"required" db:"second_pay_appointment_date"`
	SecondPayTime            *time.Time `json:"second_pay_time" binding:"required" db:"second_pay_time"`
	SecondPaySuccess         *bool      `json:"second_pay_success" binding:"required" db:"second_pay_success"`
	ExpiredDate              *time.Time `json:"expired_date" binding:"required" db:"expired_date"`
}

type SubscriptionForAdmin struct {
	Id          int        `json:"id" binding:"required"`
	UserId      int        `json:"user_id"`
	RequestDate time.Time  `json:"request_date"`
	ExpiredDate *time.Time `json:"expired_date" binding:"required"`
	Status      bool       `json:"status"`
	Amount      int        `json:"amount"`
}

func (s Subscription) Active() bool {
	if s.ExpiredDate != nil && s.ExpiredDate.Before(time.Now()) {
		return false
	}
	if s.FirstPaySuccess != nil {
		if s.SecondPaySuccess != nil && s.ExpiredDate != nil {
			if *s.FirstPaySuccess && *s.SecondPaySuccess && time.Now().Before(*s.ExpiredDate) {
				return true
			}
		} else if s.SecondPayAppointmentDate != nil {
			if *s.FirstPaySuccess && time.Now().Before(*s.SecondPayAppointmentDate) {
				return true
			}
		}
	}
	return false
}

func (s Subscription) Amount() int {
	if s.FirstPaySuccess != nil {
		if s.SecondPaySuccess != nil {
			if *s.FirstPaySuccess && *s.SecondPaySuccess {
				return FirstPayAmount + SecondPayAmount
			}
		} else if *s.FirstPaySuccess {
			return FirstPayAmount
		}
	}

	return 0
}
