package service

import (
	"bytes"
	"context"
	"github.com/mozillazg/go-unidecode"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"
	"zaimik/internal/app/models"
	"zaimik/internal/app/repository"
	"zaimik/internal/pkg/logging"
	"zaimik/internal/pkg/yandex_cloud"
)

type AdminService struct {
	repos  *repository.Repository
	logger *logging.Logger
}

func NewAdminService(repos *repository.Repository, logger *logging.Logger) *AdminService {
	return &AdminService{
		repos:  repos,
		logger: logger,
	}
}

func (a *AdminService) SelectAllUsers() ([]models.User, error) {
	return a.repos.AdminRepository.GetAllUsers()
}

func (a *AdminService) CheckAdminSession(sid string) (int, error) {
	return a.repos.AdminRepository.GetAdminIdBySession(sid)
}

func (a *AdminService) AuthorizeAdmin(code string) (string, error) {
	email := os.Getenv("ADMIN_MAIL")
	if a.repos.Storage.Exist(email, code) {
		a.repos.Storage.DeleteCode(email, code)

		SID := generateRandomString(sessLen)
		expiredDate := time.Now().AddDate(0, 0, 15).Round(time.Hour)
		if _, err := a.repos.AdminRepository.CreateAdminSession(SID, expiredDate); err != nil {
			a.logger.Errorf("error while creating admin session: %s", err.Error())
			return "", err
		}

		return SID, nil
	}

	return "", IncorrectAuthCode
}

func (a *AdminService) LogOutAdmin(sid string) error {
	return a.repos.AdminRepository.DeleteAdminSession(sid)
}

func (a *AdminService) AddCompany(company models.LoanCompanyAdmin) (int, error) {
	if !company.Valid() {
		return 0, InvalidModel
	}

	company.LogoNameOnS3 = regenerateName(company.LogoNameOnS3)

	id, err := a.repos.AdminRepository.AddLoanCompany(company)
	if err != nil {
		return 0, err
	}

	if err := yandex_cloud.PutObjectInBucket(company.LogoNameOnS3, company.Logo); err != nil {
		// TODO сделать удаление из базы
		a.logger.Errorf("error while puting object: %s", err.Error())
		return 0, err
	}

	return id, err
}

func (a *AdminService) GetAllCompanies() ([]models.LoanCompanyAdmin, error) {
	return a.repos.AdminRepository.SelectAllCompanies()
}

func (a *AdminService) GetCompanyById(id int) (models.LoanCompanyAdmin, []byte, error) {
	company, err := a.repos.AdminRepository.SelectCompanyById(id)
	if err != nil {
		return company, nil, err
	}

	logo, err := yandex_cloud.GetObjectFromYandexCloud(company.LogoNameOnS3)
	if err != nil {
		return company, nil, err
	}

	company.Logo = bytes.NewBuffer(logo)

	return company, logo, nil
}

func (a *AdminService) UpdateCompaniesPriority(companies []models.LoanCompanyPriorityUpdate) error {
	return a.repos.AdminRepository.UpdateCompaniesPriority(companies)
}

func (a *AdminService) RefactorCompany(company models.LoanCompanyAdmin) error {
	return a.repos.AdminRepository.UpdateCompanyFields(company)
}

func (a *AdminService) DeleteCompanyById(c context.Context, id int) error {
	logoNameOnS3, err := a.repos.LoanRepository.GetCompanyLogoNameById(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(c)

	ch := make(chan error, 2)
	go a.repos.DeleteCompanyById(ctx, ch, id)
	go yandex_cloud.DeleteObjectFromCloud(ctx, ch, logoNameOnS3)

	for i := 0; i < 2; i++ {
		if err := <-ch; err != nil {
			cancel()
			a.logger.Errorf("error while deleting company")
			return err
		}
	}
	cancel()
	return nil
}

func (a *AdminService) SelectReviews() ([]models.ReviewAdmin, error) {
	return a.repos.AdminRepository.SelectAllReviews()
}

func (a *AdminService) GetReviewById(id int) (models.ReviewAdmin, error) {
	return a.repos.AdminRepository.GetReviewById(id)
}

func (a *AdminService) SetReviewModerated(review models.ReviewAdmin) (int, error) {
	if review.Valid() {
		return a.repos.AdminRepository.UpdateReview(review)
	}

	return 0, InvalidReview
}

func (a *AdminService) UploadReview(review models.ReviewAdmin) (int, error) {
	return a.repos.AdminRepository.InsertReview(review)
}

func (a *AdminService) DeleteReview(id int) error {
	return a.repos.AdminRepository.DeleteReview(id)
}

func regenerateName(s string) string {
	s = unidecode.Unidecode(s)
	ext := filepath.Ext(s)

	var newS string
	for i := 0; i < utf8.RuneCountInString(s)-utf8.RuneCountInString(ext); i++ {
		newS += string(s[i])
	}

	newS += generateRandomString(20)

	return newS + ext
}
