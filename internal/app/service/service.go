package service

import (
	"context"
	"time"
	"zaimik/internal/app/models"
	"zaimik/internal/app/repository"
	"zaimik/internal/pkg/logging"
)

type Authorization interface {
	CheckEmailAndSendAuthCode(email, emailTemplate, subject string) (code string, err error)
	AuthorizeUser(email, code string) (SID string, err error)
	CheckSession(sessionId string) (int, error)
	UpdateUser(user models.User) error
	DeleteSession(sid string, userId int) error
}

type LoanCompaniesManager interface {
	AddLoanRequest(loan models.Loan) (int, error)
	InitSubscription(userId, loanId int) (models.Subscription, error)
	ActivateSubscription(subscription models.Subscription) error
	GetAllSubscriptions() ([]models.SubscriptionForAdmin, error)
	GetSubscriptionById(id int) (models.SubscriptionForAdmin, error)
	RefactorSubscriptionExpiredDate(id int, expiredDate *time.Time) error

	CheckSubscriptionByUserId(id int) (active bool, expiredDate time.Time, err error)
	GetReviews() ([]models.Review, error)
	GetSortedSuitableCatalogByUserId(id int) ([]models.LoanCompany, error)
	AddReview(review models.Review) error
}

type Administration interface {
	CheckAdminSession(sid string) (int, error)
	AuthorizeAdmin(code string) (string, error)
	LogOutAdmin(sid string) error

	SelectAllUsers() ([]models.User, error)

	AddCompany(company models.LoanCompanyAdmin) (int, error)
	GetCompanyById(id int) (models.LoanCompanyAdmin, []byte, error)
	GetAllCompanies() ([]models.LoanCompanyAdmin, error)
	UpdateCompaniesPriority([]models.LoanCompanyPriorityUpdate) error
	RefactorCompany(company models.LoanCompanyAdmin) error
	DeleteCompanyById(ctx context.Context, id int) error

	SelectReviews() ([]models.ReviewAdmin, error)
	GetReviewById(id int) (models.ReviewAdmin, error)
	SetReviewModerated(review models.ReviewAdmin) (int, error)
	UploadReview(review models.ReviewAdmin) (int, error)
	DeleteReview(id int) error
}

type Service struct {
	Authorization
	Administration
	LoanCompaniesManager
}

func NewService(repos *repository.Repository, logger *logging.Logger) *Service {
	return &Service{
		Authorization:        NewAuthService(repos, logger),
		Administration:       NewAdminService(repos, logger),
		LoanCompaniesManager: NewCompanyManager(repos, logger),
	}
}
