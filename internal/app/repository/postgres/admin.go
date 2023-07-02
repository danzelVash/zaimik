package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"zaimik/internal/app/models"
	"zaimik/internal/pkg/logging"
)

type AdminPostgres struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func NewAdminPostgres(db *sqlx.DB, logger *logging.Logger) *AdminPostgres {
	return &AdminPostgres{
		db:     db,
		logger: logger,
	}
}

func (a *AdminPostgres) GetAllUsers() ([]models.User, error) {
	query := fmt.Sprintf("SELECT id, last_name, first_name, sur_name, email, phone_number, city, is_admin FROM %s;", usersTable)
	rows, err := a.db.Query(query)

	if err == ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			a.logger.Errorf("memory leak in admin_pstgrs.SelectAllUsers error closing *sql.Rows: %s", err.Error())
		}
	}(rows)

	users := make([]models.User, 0, 10)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.LastName, &user.FirstName, &user.SurName, &user.Email, &user.PhoneNumber, &user.City, &user.Admin); err != nil {
			a.logger.Errorf("error while scanning user in admin_pstgrs.SelectAllUsers: %s", err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		a.logger.Errorf("error while rows.Err(): %s", err.Error())
	}
	return users, err
}

func (a *AdminPostgres) GetAdminIdBySession(sid string) (int, error) {
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE session=$1 AND user_id = (SELECT id FROM %s WHERE is_admin=true);", sessionTable, usersTable)

	var id int
	row := a.db.QueryRow(query, sid)
	err := row.Scan(&id)
	if err != sql.ErrNoRows && err != nil {
		a.logger.Errorf("error getting admin id: %s", err.Error())
	} else if err == sql.ErrNoRows {
		a.logger.Errorf("admin id not found in db: %s")
	}
	return id, err
}

func (a *AdminPostgres) CreateAdminSession(sid string, expiredDate time.Time) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, session, expired_date) VALUES ((SELECT id FROM %s WHERE is_admin=true), $1, $2) RETURNING id;", sessionTable, usersTable)
	row := a.db.QueryRow(query, sid, expiredDate)

	var id int
	if err := row.Scan(&id); err != nil {
		a.logger.Errorf("error while scanning id: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (a *AdminPostgres) DeleteAdminSession(sid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE session = $1 AND user_id = (SELECT id FROM %s WHERE is_admin = TRUE);", sessionTable, usersTable)
	row := a.db.QueryRow(query, sid)
	return row.Err()
}

func (a *AdminPostgres) AddLoanCompany(company models.LoanCompanyAdmin) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, logo_name_on_s3, link_on_company_site, max_loan_amount, max_loan_duration, min_loan_percent, priority) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;", loanCompaniesTable)

	row := a.db.QueryRow(query, company.Name, company.LogoNameOnS3, company.LinkOnCompanySite, company.MaxLoanAmount, company.MaxLoanDuration, company.MinLoanPercent, company.Priority)

	var id int
	if err := row.Scan(&id); err != nil {
		a.logger.Errorf("error insert loan company into db: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (a *AdminPostgres) SelectAllCompanies() ([]models.LoanCompanyAdmin, error) {
	query := fmt.Sprintf("SELECT id, name, logo_name_on_s3, link_on_company_site, max_loan_amount, max_loan_duration, min_loan_percent, priority FROM %s ORDER BY priority DESC;", loanCompaniesTable)
	rows, err := a.db.Query(query)

	companies := make([]models.LoanCompanyAdmin, 0, 10)
	if err == ErrNoRows {
		return companies, nil
	} else if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			a.logger.Errorf("memory leak in admin_pstgrs.SelectAllUsers error closing *sql.Rows: %s", err.Error())
		}
	}(rows)

	for rows.Next() {
		var company models.LoanCompanyAdmin
		if err := rows.Scan(&company.Id, &company.Name, &company.LogoNameOnS3, &company.LinkOnCompanySite, &company.MaxLoanAmount, &company.MaxLoanDuration, &company.MinLoanPercent, &company.Priority); err != nil {
			a.logger.Errorf("error while scanning user in admin_pstgrs.SelectAllUsers: %s", err.Error())
			return nil, err
		}
		companies = append(companies, company)
	}
	if err := rows.Err(); err != nil {
		a.logger.Errorf("error while rows.Err(): %s", err.Error())
	}
	return companies, err
}

func (a *AdminPostgres) UpdateCompaniesPriority(companies []models.LoanCompanyPriorityUpdate) error {
	query := fmt.Sprintf("UPDATE %s SET priority = updated.priority FROM (VALUES ", loanCompaniesTable)

	values := make([]interface{}, 0, len(companies)*2)

	i := 1
	for ind, company := range companies {
		if ind != len(companies)-1 {
			query += fmt.Sprintf("($%d::integer, $%d::integer), ", i, i+1)
			i += 2
			values = append(values, company.Id, company.Priority)
		} else {
			query += fmt.Sprintf("($%d::integer, $%d::integer)", i, i+1)
			values = append(values, company.Id, company.Priority)
		}

	}

	query += fmt.Sprintf(") AS updated (id, priority) WHERE %s.id = updated.id;", loanCompaniesTable)

	res, err := a.db.Exec(query, values...)
	if err != nil {
		return err
	}

	numAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if numAffected != int64(len(companies)) {
		a.logger.Infof("while update companies priority was affected %d rows from %d", numAffected, len(companies))
	}

	return nil
}

func (a *AdminPostgres) SelectCompanyById(id int) (models.LoanCompanyAdmin, error) {
	query := fmt.Sprintf("SELECT name, logo_name_on_s3, link_on_company_site, max_loan_amount, max_loan_duration, min_loan_percent, priority FROM companies WHERE id = $1;")
	row := a.db.QueryRow(query, id)

	var review models.LoanCompanyAdmin
	review.Id = id
	err := row.Scan(&review.Name, &review.LogoNameOnS3, &review.LinkOnCompanySite, &review.MaxLoanAmount, &review.MaxLoanDuration, &review.MinLoanPercent, &review.Priority)
	return review, err
}

func (a *AdminPostgres) UpdateCompanyFields(company models.LoanCompanyAdmin) error {
	query := fmt.Sprintf("UPDATE companies SET name = $1, link_on_company_site = $2, max_loan_amount = $3, max_loan_duration = $4, min_loan_percent = $5 WHERE id = $6 RETURNING id;")

	row := a.db.QueryRow(query, company.Name, company.LinkOnCompanySite, company.MaxLoanAmount, company.MaxLoanDuration, company.MinLoanPercent, company.Id)

	var id int
	err := row.Scan(&id)
	return err
}

func (a *AdminPostgres) DeleteCompanyById(ctx context.Context, ch chan<- error, id int) {
	query := fmt.Sprintf("DELETE FROM companies WHERE id = $1 RETURNING id;")
	var returned int
	row := a.db.QueryRowContext(ctx, query, id)
	ch <- row.Scan(&returned)
}

func (a *AdminPostgres) SelectAllReviews() ([]models.ReviewAdmin, error) {
	query := fmt.Sprintf("SELECT id, reviewer_name, reviewer_phone, review, moderated FROM %s ORDER BY id, moderated;", reviewsTable)

	rows, err := a.db.Query(query)

	reviews := make([]models.ReviewAdmin, 0, 10)
	if err == ErrNoRows {
		return reviews, nil
	} else if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			a.logger.Errorf("memory leak in admin_pstgrs.SelectAllUsers error closing *sql.Rows: %s", err.Error())
		}
	}(rows)

	for rows.Next() {
		var review models.ReviewAdmin
		if err := rows.Scan(&review.Id, &review.ReviewerName, &review.ReviewerPhone, &review.Review, &review.Moderated); err != nil {
			a.logger.Errorf("error while scanning review: %s", err.Error())
			return nil, err
		}
		reviews = append(reviews, review)
	}
	if err := rows.Err(); err != nil {
		a.logger.Errorf("error while rows.Err(): %s", err.Error())
	}
	return reviews, err
}

func (a *AdminPostgres) GetReviewById(id int) (models.ReviewAdmin, error) {
	query := fmt.Sprintf("SELECT id, reviewer_name, reviewer_phone, review, moderated FROM %s WHERE id = $1;", reviewsTable)
	row := a.db.QueryRow(query, id)

	var review models.ReviewAdmin
	err := row.Scan(&review.Id, &review.ReviewerName, &review.ReviewerPhone, &review.Review, &review.Moderated)
	return review, err
}

func (a *AdminPostgres) UpdateReview(review models.ReviewAdmin) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET reviewer_name = $1, reviewer_phone = $2, review = $3, moderated = $4 WHERE id = $5 RETURNING id;", reviewsTable)

	row := a.db.QueryRow(query, review.ReviewerName, review.ReviewerPhone, review.Review, review.Moderated, review.Id)

	var id int
	if err := row.Scan(&id); err != nil {
		a.logger.Errorf("error update review: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (a *AdminPostgres) InsertReview(review models.ReviewAdmin) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (reviewer_name, reviewer_phone, review, moderated) VALUES ($1, $2, $3, $4) RETURNING id;", reviewsTable)
	row := a.db.QueryRow(query, review.ReviewerName, review.ReviewerPhone, review.Review, review.Moderated)

	var id int
	if err := row.Scan(&id); err != nil {
		a.logger.Errorf("error while scanning id: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (a *AdminPostgres) DeleteReview(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id;", reviewsTable)
	row := a.db.QueryRow(query, id)

	var deleted int
	if err := row.Scan(&deleted); err != nil {
		a.logger.Errorf("error while deleting review with id = %d: %s", id, err.Error())
		return err
	}

	return nil
}
