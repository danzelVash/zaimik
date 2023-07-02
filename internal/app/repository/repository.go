package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
	"zaimik/internal/app/models"
	"zaimik/internal/app/repository/postgres"
	"zaimik/internal/pkg/inmemory_storage"
	"zaimik/internal/pkg/logging"
)

type AuthRepository interface {
	GetUserByEmail(email string) (models.User, error)
	CreateUser(email string) (int, error)
	CreateSession(userId int, SID string, expiredDate time.Time) (int, error)
	CheckSession(sessionId string) (int, error)
	UpdateUser(user models.User) error
	DeleteSession(sid string, userId int) error
}

type LoanRepository interface {
	AddLoanRequest(loan models.Loan) (int, error)
	InitSubscription(subscription models.Subscription) (int, error)
	UpdateSubscription(subscription models.Subscription) error
	UpdateSubscriptionExpiredDate(id int, expiredDate *time.Time) error
	GetSubscriptionByUserId(userId int) (models.Subscription, error)
	GetSubscriptionById(id int) (models.Subscription, error)
	GetAllSubscriptions() ([]models.Subscription, error)

	GetLoanCharacteristicsByUserId(id int) (models.Loan, error)
	GetSuitableLoanCompanies(loan models.Loan) ([]models.LoanCompany, error)
	GetReviews() ([]models.Review, error)
	AddReview(review models.Review) (int, error)
	GetCompanyLogoNameById(id int) (string, error)
}

type Storage interface {
	AddCode(key string, val string)
	Exist(key, val string) bool
	DeleteCode(key, val string)
}

type AdminRepository interface {
	CreateAdminSession(sid string, expiredDate time.Time) (int, error)
	DeleteAdminSession(sid string) error

	GetAllUsers() ([]models.User, error)
	GetAdminIdBySession(sid string) (int, error)

	SelectAllCompanies() ([]models.LoanCompanyAdmin, error)
	AddLoanCompany(company models.LoanCompanyAdmin) (int, error)
	UpdateCompaniesPriority(companies []models.LoanCompanyPriorityUpdate) error
	SelectCompanyById(id int) (models.LoanCompanyAdmin, error)
	UpdateCompanyFields(company models.LoanCompanyAdmin) error
	DeleteCompanyById(ctx context.Context, ch chan<- error, id int)

	SelectAllReviews() ([]models.ReviewAdmin, error)
	GetReviewById(id int) (models.ReviewAdmin, error)
	UpdateReview(review models.ReviewAdmin) (int, error)
	InsertReview(review models.ReviewAdmin) (int, error)
	DeleteReview(id int) error
}

type Repository struct {
	AuthRepository
	AdminRepository
	LoanRepository
	Storage
}

func NewRepository(db *sqlx.DB, storage *inmemory_storage.DataStorage, logger *logging.Logger) *Repository {
	return &Repository{
		AuthRepository:  postgres.NewAuthPostgres(db, logger),
		AdminRepository: postgres.NewAdminPostgres(db, logger),
		LoanRepository:  postgres.NewLoanRepos(db, logger),
		Storage:         storage,
	}
}
