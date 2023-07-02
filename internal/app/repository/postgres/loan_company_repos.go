package postgres

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"zaimik/internal/app/models"
	"zaimik/internal/pkg/logging"
)

type LoanRepos struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func (r *LoanRepos) UpdateSubscriptionExpiredDate(id int, expiredDate *time.Time) error {
	query := fmt.Sprintf("UPDATE subscriptions SET expired_date = $1 WHERE id = $2;")
	//fmt.Println(query)
	_, err := r.db.Exec(query, expiredDate, id)
	return err
}

func NewLoanRepos(db *sqlx.DB, logger *logging.Logger) *LoanRepos {
	return &LoanRepos{
		db:     db,
		logger: logger,
	}
}

func (r *LoanRepos) AddLoanRequest(loan models.Loan) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, amount, duration) VALUES ($1, $2, $3) RETURNING id;", loanRequestsTable)
	row := r.db.QueryRow(query, loan.UserId, loan.LoanAmount, loan.LoanDurationInDays)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *LoanRepos) InitSubscription(subscription models.Subscription) (int, error) {
	query := fmt.Sprintf("INSERT INTO subscriptions (user_id, loan_id, request_date) VALUES  ($1, $2, $3) RETURNING id;")
	row := r.db.QueryRow(query, subscription.UserId, subscription.LoanId, subscription.RequestDate)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *LoanRepos) UpdateSubscription(subscription models.Subscription) error {
	query := fmt.Sprintf("UPDATE subscriptions SET first_pay_time = $1, first_pay_success = $2, second_pay_appointment_date = $3, second_pay_time = $4, second_pay_success = $5, expired_date = $6 WHERE id = $7;")
	row := r.db.QueryRow(query, subscription.FirstPayTime, subscription.FirstPaySuccess, subscription.SecondPayAppointmentDate, subscription.SecondPayTime, subscription.SecondPaySuccess, subscription.ExpiredDate, subscription.Id)
	return row.Err()
}

func (r *LoanRepos) GetSubscriptionByUserId(userId int) (models.Subscription, error) {
	query := fmt.Sprintf("SELECT id, user_id, loan_id, request_date, first_pay_time, first_pay_success, second_pay_appointment_date, second_pay_time, second_pay_success, expired_date FROM %s WHERE user_id = $1;", subscriptionsTable)
	row := r.db.QueryRow(query, userId)

	var subscription models.Subscription
	err := row.Scan(&subscription.Id, &subscription.UserId, &subscription.LoanId, &subscription.RequestDate, &subscription.FirstPayTime, &subscription.FirstPaySuccess, &subscription.SecondPayAppointmentDate, &subscription.SecondPayTime, &subscription.SecondPaySuccess, &subscription.ExpiredDate)
	if err == sql.ErrNoRows {
		return models.Subscription{}, ErrNoRows
	}
	return subscription, err

}

func (r *LoanRepos) GetSubscriptionById(id int) (models.Subscription, error) {
	query := fmt.Sprintf("SELECT id, user_id, loan_id, request_date, first_pay_time, first_pay_success, second_pay_appointment_date, second_pay_time, second_pay_success, expired_date FROM subscriptions WHERE id = $1;")
	row := r.db.QueryRow(query, id)
	var subscription models.Subscription
	err := row.Scan(&subscription.Id, &subscription.UserId, &subscription.LoanId, &subscription.RequestDate, &subscription.FirstPayTime, &subscription.FirstPaySuccess, &subscription.SecondPayAppointmentDate, &subscription.SecondPayTime, &subscription.SecondPaySuccess, &subscription.ExpiredDate)
	if err == sql.ErrNoRows {
		return models.Subscription{}, ErrNoRows
	}
	return subscription, err
}

func (r *LoanRepos) GetAllSubscriptions() ([]models.Subscription, error) {
	query := fmt.Sprintf("SELECT id, user_id, loan_id, request_date, first_pay_time, first_pay_success, second_pay_appointment_date, second_pay_time, second_pay_success, expired_date FROM %s;", subscriptionsTable)
	rows, err := r.db.Query(query)

	if err == ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			r.logger.Errorf("memory leak in admin_pstgrs.SelectAllUsers error closing *sql.Rows: %s", err.Error())
		}
	}(rows)

	subscriptions := make([]models.Subscription, 0, 10)
	for rows.Next() {
		var user models.Subscription
		if err := rows.Scan(&user.Id, &user.UserId, &user.LoanId, &user.RequestDate, &user.FirstPayTime, &user.FirstPaySuccess, &user.SecondPayAppointmentDate, &user.SecondPayTime, &user.SecondPaySuccess, &user.ExpiredDate); err != nil {
			r.logger.Errorf("error while scanning user in admin_pstgrs.SelectAllUsers: %s", err.Error())
			return nil, err
		}
		subscriptions = append(subscriptions, user)
	}
	if err := rows.Err(); err != nil {
		r.logger.Errorf("error while rows.Err(): %s", err.Error())
	}
	return subscriptions, err
}

func (r *LoanRepos) GetLoanCharacteristicsByUserId(id int) (models.Loan, error) {
	query := fmt.Sprintf("SELECT user_id, amount, duration FROM %s WHERE id = $1;", loanRequestsTable)
	row := r.db.QueryRow(query, id)

	var loan models.Loan
	err := row.Scan(&loan.UserId, &loan.LoanAmount, &loan.LoanDurationInDays)
	return loan, err
}

func (r *LoanRepos) GetSuitableLoanCompanies(loan models.Loan) ([]models.LoanCompany, error) {
	query := fmt.Sprintf("SELECT id, name, link_on_company_site, min_loan_percent, priority FROM %s WHERE max_loan_amount > $1 AND max_loan_duration > $2 ORDER BY priority DESC;", loanCompaniesTable)
	rows, err := r.db.Query(query, loan.LoanAmount, loan.LoanDurationInDays)

	if err == ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			r.logger.Errorf("memory leak in admin_pstgrs.SelectAllUsers error closing *sql.Rows: %s", err.Error())
		}
	}(rows)

	companies := make([]models.LoanCompany, 0, 10)
	for rows.Next() {
		var company models.LoanCompany
		if err := rows.Scan(&company.Id, &company.Name, &company.LinkOnCompanySite, &company.MinLoanPercent, &company.Priority); err != nil {
			r.logger.Errorf("error while scanning user in admin_pstgrs.SelectAllUsers: %s", err.Error())
			return nil, err
		}
		companies = append(companies, company)
	}
	if err := rows.Err(); err != nil {
		r.logger.Errorf("error while rows.Err(): %s", err.Error())
	}
	return companies, err
}

func (r *LoanRepos) GetReviews() ([]models.Review, error) {
	query := fmt.Sprintf("SELECT id, reviewer_name, reviewer_phone, review FROM %s WHERE moderated = true;", reviewsTable)
	rows, err := r.db.Query(query)

	if err == ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			r.logger.Errorf("memory leak in admin_pstgrs.SelectAllUsers error closing *sql.Rows: %s", err.Error())
		}
	}(rows)

	reviews := make([]models.Review, 0, 10)
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.Id, &review.ReviewerName, &review.ReviwerPhone, &review.Review); err != nil {
			r.logger.Errorf("error while scanning user in admin_pstgrs.SelectAllUsers: %s", err.Error())
			return nil, err
		}
		reviews = append(reviews, review)
	}
	if err := rows.Err(); err != nil {
		r.logger.Errorf("error while rows.Err(): %s", err.Error())
	}
	return reviews, err
}

func (r *LoanRepos) AddReview(review models.Review) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (reviewer_name, reviewer_phone, review) VALUES ($1, $2, $3) RETURNING id;", reviewsTable)
	row := r.db.QueryRow(query, review.ReviewerName, review.ReviwerPhone, review.Review)

	var id int
	if err := row.Scan(&id); err != nil {
		r.logger.Errorf("unknown error occured while getting response from db: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (r *LoanRepos) GetCompanyLogoNameById(id int) (string, error) {
	query := fmt.Sprintf("SELECT logo_name_on_s3 FROM companies WHERE id = $1;")
	row := r.db.QueryRow(query, id)

	var logoName string
	if err := row.Scan(&logoName); err != nil {
		return "", err
	}
	return logoName, nil
}
