package service

import (
	"time"
	"zaimik/internal/app/models"
	"zaimik/internal/app/repository"
	"zaimik/internal/app/repository/postgres"
	"zaimik/internal/pkg/logging"
)

type CompanyManager struct {
	repos  *repository.Repository
	logger *logging.Logger
}

func (m *CompanyManager) RefactorSubscriptionExpiredDate(id int, expiredDate *time.Time) error {
	return m.repos.LoanRepository.UpdateSubscriptionExpiredDate(id, expiredDate)
}

func NewCompanyManager(repos *repository.Repository, logger *logging.Logger) *CompanyManager {
	return &CompanyManager{
		repos:  repos,
		logger: logger,
	}
}

func (m *CompanyManager) AddLoanRequest(loan models.Loan) (int, error) {
	if loan.Valid() {
		return m.repos.LoanRepository.AddLoanRequest(loan)
	}
	return 0, InvalidModel
}

func (m *CompanyManager) InitSubscription(userId, loanId int) (models.Subscription, error) {
	subscription := models.Subscription{
		UserId:      userId,
		LoanId:      loanId,
		RequestDate: time.Now(),
	}
	id, err := m.repos.InitSubscription(subscription)

	if err != nil {
		return models.Subscription{}, err
	}

	subscription.Id = id

	return subscription, nil
}

func (m *CompanyManager) ActivateSubscription(subscription models.Subscription) error {
	if subscription.UserId <= 0 || subscription.Id <= 0 {
		return InvalidModel
	}
	return m.repos.LoanRepository.UpdateSubscription(subscription)
}

func (m *CompanyManager) GetAllSubscriptions() ([]models.SubscriptionForAdmin, error) {
	subscriptions, err := m.repos.LoanRepository.GetAllSubscriptions()
	if err != nil {
		return nil, err
	}

	adminSubscriptions := make([]models.SubscriptionForAdmin, 0, len(subscriptions))
	for _, subscription := range subscriptions {
		sub := models.SubscriptionForAdmin{
			Id:          subscription.Id,
			UserId:      subscription.UserId,
			RequestDate: subscription.RequestDate,
			ExpiredDate: subscription.ExpiredDate,
			Status:      subscription.Active(),
			Amount:      subscription.Amount(),
		}
		adminSubscriptions = append(adminSubscriptions, sub)
	}

	return adminSubscriptions, nil
}

func (m *CompanyManager) GetSubscriptionById(id int) (models.SubscriptionForAdmin, error) {
	sub, err := m.repos.GetSubscriptionById(id)
	if err != nil {
		return models.SubscriptionForAdmin{}, err
	}
	subForAdmin := models.SubscriptionForAdmin{
		Id:          id,
		UserId:      sub.UserId,
		RequestDate: sub.RequestDate,
		ExpiredDate: sub.ExpiredDate,
		Status:      sub.Active(),
		Amount:      sub.Amount(),
	}
	return subForAdmin, err
}

func (m *CompanyManager) CheckSubscriptionByUserId(id int) (active bool, expiredDate time.Time, err error) {
	subscription, err := m.repos.GetSubscriptionByUserId(id)
	if err == postgres.ErrNoRows {
		return false, time.Time{}, HaveNoSubscription
	} else if err != nil {
		return false, time.Time{}, err
	}

	if subscription.Active() {
		if subscription.ExpiredDate != nil {
			return true, *subscription.ExpiredDate, nil
		} else {
			return false, time.Time{}, InvalidModel
		}
	}

	return false, *subscription.ExpiredDate, nil
}

func (m *CompanyManager) GetSortedSuitableCatalogByUserId(id int) ([]models.LoanCompany, error) {
	// TODO сделать проверку на активную подписку
	loan, err := m.repos.LoanRepository.GetLoanCharacteristicsByUserId(id)
	if err != nil {
		return nil, err
	}

	companies, err := m.repos.LoanRepository.GetSuitableLoanCompanies(loan)
	switch err {
	case nil:
		return companies, err
	case postgres.ErrNoRows:
		return nil, HaveNoSuitableCompanies
	default:
		m.logger.Errorf("unknown error while getting suitable companies from db: %s", err.Error())
		return nil, err
	}
}

func (m *CompanyManager) GetReviews() ([]models.Review, error) {
	return m.repos.LoanRepository.GetReviews()
}

func (m *CompanyManager) AddReview(review models.Review) error {
	_, err := m.repos.LoanRepository.AddReview(review)
	return err
}
