package service

import (
	"fmt"
	"math/rand"
	"time"
	"zaimik/internal/app/models"
	"zaimik/internal/app/repository"
	"zaimik/internal/app/repository/postgres"
	"zaimik/internal/pkg/logging"
	"zaimik/internal/pkg/smtp"
)

const (
	symbolsForAuthCode = "0123456789"
	authCodeLength     = 4
	alphabet           = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890!#*()_+=-<>?"
	sessLen            = 40
)

type AuthService struct {
	repos  *repository.Repository
	logger *logging.Logger
}

func NewAuthService(repos *repository.Repository, logger *logging.Logger) *AuthService {
	return &AuthService{
		repos:  repos,
		logger: logger,
	}
}

func (s *AuthService) CheckEmailAndSendAuthCode(email, emailTemplate, subject string) (string, error) {
	if email == "" {
		return "", InvalidEmailAddress

	}

	code := generateAuthCode()

	emailParams := smtp.EmailParams{
		TemplateName: fmt.Sprintf("templates/email/%s", emailTemplate),
		TemplateVars: struct {
			Code string
		}{
			Code: code,
		},
		Destination: email,
		Subject:     subject,
	}

	if err := smtp.SendEmail(emailParams); err != smtp.BadEmail && err != nil {
		s.logger.Errorf("unknown error occured while sending mail: %s", err.Error())
		return "", err
	} else if err == smtp.BadEmail {
		s.logger.Errorf("invalid email address: %s", err.Error())
		return "", err
	}

	s.repos.Storage.AddCode(email, code)

	return code, nil
}

func (s *AuthService) AuthorizeUser(email, code string) (string, error) {
	if s.repos.Storage.Exist(email, code) {

		s.repos.Storage.DeleteCode(email, code)

		user, err := s.repos.AuthRepository.GetUserByEmail(email)
		switch err {
		case nil:
		//	pass
		case postgres.ErrNoRows:
			id, err := s.repos.AuthRepository.CreateUser(email)
			if err != nil {
				s.logger.Errorf("error creating user: %s", err.Error())
				return "", err
			}
			user.Id = id
		default:
			s.logger.Errorf("unknown error occured while GetUserByEmail: %s", err.Error())
			return "", err
		}

		SID := generateRandomString(sessLen)
		expiredDate := time.Now().AddDate(0, 0, 15).Round(time.Hour)
		if _, err := s.repos.AuthRepository.CreateSession(user.Id, SID, expiredDate); err != nil {
			s.logger.Errorf("error while creating session: %s", err.Error())
			return "", err
		}

		return SID, nil
	}

	return "", IncorrectAuthCode
}

func (s *AuthService) CheckSession(sessionId string) (int, error) {
	return s.repos.AuthRepository.CheckSession(sessionId)
}

func (s *AuthService) UpdateUser(user models.User) error {
	return s.repos.AuthRepository.UpdateUser(user)
}

func (s *AuthService) DeleteSession(sid string, userId int) error {
	return s.repos.AuthRepository.DeleteSession(sid, userId)
}

func generateAuthCode() string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < authCodeLength; i++ {
		code += string(symbolsForAuthCode[rand.Intn(len(symbolsForAuthCode))])

	}

	return code
}

// TODO сделать с помощью generic проверку любых структур на то, что все поля ненулевые
func validateStruct() error {
	return nil
}

func generateRandomString(strLen int) string {
	rand.Seed(time.Now().UnixNano())
	str := ""
	for i := 0; i < strLen; i++ {
		str += string(alphabet[rand.Intn(len(alphabet))])
	}
	return str
}
