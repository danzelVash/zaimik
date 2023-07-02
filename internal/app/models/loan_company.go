package models

import (
	"io"
	"os"
)

type LogoType interface {
	[]byte | *os.File
}

type LoanCompany struct {
	Id                int    `json:"-" db:"id"`
	Name              string `json:"name" binding:"required" db:"name"`
	Logo              []byte `json:"logo" binding:"required" db:"-"`
	LogoNameOnS3      string `json:"_" binding:"required" db:"logo_name_on_s3"`
	LinkOnCompanySite string `json:"link_on_company_site" binding:"required" db:"link_on_company_site"`
	MaxLoanAmount     int    `json:"max_loan_amount" binding:"required" db:"max_loan_amount"`
	MaxLoanDuration   int    `json:"max_loan_duration" binding:"required" db:"max_loan_duration"`
	MinLoanPercent    int    `json:"min_loan_percent" binding:"required" db:"min_loan_percent"`
	Priority          int    `json:"priority" binding:"required" db:"priority"`
}

type LoanCompanyAdmin struct {
	Id                int       `json:"id" binding:"-" db:"id"`
	Name              string    `json:"name" form:"name" binding:"required" db:"name"`
	Logo              io.Reader `json:"logo" form:"-" binding:"-" db:"-"`
	LogoNameOnS3      string    `json:"logo_name_on_s3" form:"-" binding:"-" db:"logo_name_on_s3"`
	LinkOnCompanySite string    `json:"link_on_company_site" form:"link_on_company_site" binding:"required" db:"link_on_company_site"`
	MaxLoanAmount     int       `json:"max_loan_amount" form:"max_loan_amount" binding:"required" db:"max_loan_amount"`
	MaxLoanDuration   int       `json:"max_loan_duration" form:"max_loan_duration" binding:"required" db:"max_loan_duration"`
	MinLoanPercent    int       `json:"min_loan_percent" form:"min_loan_percent" binding:"required" db:"min_loan_percent"`
	Priority          int       `json:"priority" binding:"-" db:"priority"`
}

type LoanCompanyPriorityUpdate struct {
	Id       int `json:"id" binding:"required" db:"id"`
	Priority int `json:"priority" binding:"required" db:"priority"`
}

func (l LoanCompanyAdmin) Valid() bool {
	if l.Name == "" || l.Logo == nil || l.LinkOnCompanySite == "" || l.MinLoanPercent == 0 {
		return false
	}
	return true
}
