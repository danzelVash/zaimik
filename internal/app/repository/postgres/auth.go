package postgres

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"zaimik/internal/app/models"
	"zaimik/internal/pkg/logging"
)

type AuthPostgres struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func NewAuthPostgres(db *sqlx.DB, logger *logging.Logger) *AuthPostgres {
	return &AuthPostgres{
		db:     db,
		logger: logger,
	}
}

func (a *AuthPostgres) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id, last_name, first_name, email, phone_number, city, is_admin FROM %s WHERE email=$1;", usersTable)

	if err := a.db.Get(&user, query, email); err == sql.ErrNoRows {
		return user, ErrNoRows
	} else {
		return user, err
	}
}

func (a *AuthPostgres) CreateUser(email string) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (email) VALUES ($1) RETURNING id;", usersTable)
	row := a.db.QueryRow(query, email)

	var id int
	if err := row.Scan(&id); err != nil {
		a.logger.Errorf("unknown error occured while getting response from db: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (a *AuthPostgres) CreateSession(userId int, SID string, expiredDate time.Time) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, session, expired_date) VALUES ($1, $2, $3) RETURNING id;", sessionTable)
	//query := fmt.Sprintf("WITH inserted AS (INSERT INTO %s (email) SELECT $1 WHERE NOT EXISTS(SELECT 1 FROM %s WHERE email = $2) RETURNING id) INSERT INTO %s (user_id, session, expired_date) SELECT COALESCE(id, (SELECT id FROM users WHERE email = $1)), $3, $4 FROM inserted RETURNING id;",
	//	usersTable, usersTable, sessionTable)

	row := a.db.QueryRow(query, userId, SID, expiredDate)

	var id int
	if err := row.Scan(&id); err != nil {
		a.logger.Errorf("error while scanning id: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (a *AuthPostgres) CheckSession(sessionId string) (int, error) {
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE session=$1;", sessionTable)

	row := a.db.QueryRow(query, sessionId)

	var id int
	if err := row.Scan(&id); err != nil {
		a.logger.Error("error while scanning id: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (a *AuthPostgres) UpdateUser(user models.User) error {
	query := fmt.Sprintf("UPDATE %s SET (last_name, first_name, sur_name, phone_number, city) = ($1, $2, $3, $4, $5) WHERE id = $6;", usersTable)
	row := a.db.QueryRow(query, user.LastName, user.FirstName, user.SurName, user.PhoneNumber, user.City, user.Id)
	return row.Err()
}

func (a *AuthPostgres) DeleteSession(sid string, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE session = $1 AND user_id = $2;", sessionTable)
	row := a.db.QueryRow(query, sid, userId)
	err := row.Err()
	if err != nil {
		a.logger.Errorf("error deleting where session = %s: %s", sid, err.Error())
	}
	return err

}
