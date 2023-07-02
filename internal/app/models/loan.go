package models

type Loan struct {
	Id                 int `json:"-" binding:"-" db:"id"`
	UserId             int `json:"-" binding:"required" db:"user_id"`
	LoanAmount         int `json:"loan_amount" binding:"required" db:"amount"`
	LoanDurationInDays int `json:"loan_duration" binding:"required" db:"duration"`
}

func (l *Loan) Valid() bool {
	if l.UserId <= 0 {
		return false
	}
	return true
}
