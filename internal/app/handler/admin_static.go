package handler

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"html/template"
	"net/http"
	"strconv"
	"zaimik/internal/app/models"
	"zaimik/internal/pkg/template_parser"
)

func (h *Handler) authPage(ctx *gin.Context) {

	params := template_parser.TemplateParams{
		TemplateName: "auth.html",
		Vars: struct {
			Domain string
		}{
			Domain: viper.GetString("http.domain"),
		},
	}

	data, err := template_parser.TemplateParser(params)
	if err != nil {
		h.logger.Errorf("error while parsing template: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
	}

	ctx.Data(http.StatusOK, "text/html", data)
}

func (h *Handler) panelPage(ctx *gin.Context) {
	//if _, ok := ctx.Get(adminCtx); !ok {
	//	newRespErr(ctx, http.StatusForbidden, "have no session")
	//	return
	//}

	tmplNames := []string{"left_bar.html", "footer.html"}
	templates, err := template_parser.GetTemplates(tmplNames)
	if err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	users, err := h.services.Administration.SelectAllUsers()
	if err != nil {
		h.logger.Errorf("error while getting users: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	params := template_parser.TemplateParams{
		TemplateName: "panel.html",
		Vars: struct {
			LeftBar template.HTML
			Footer  template.HTML
			Users   []models.User
			Domain  string
		}{
			LeftBar: template.HTML(templates[0]),
			Footer:  template.HTML(templates[1]),
			Users:   users,
			Domain:  viper.GetString("http.protocol"),
		},
	}

	data, err := template_parser.TemplateParser(params)
	if err != nil {
		h.logger.Errorf("error while parsing template: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.Data(http.StatusOK, "text/html", data)
}

func (h *Handler) companiesPage(ctx *gin.Context) {
	tmplNames := []string{"left_bar.html", "footer.html"}
	templates, err := template_parser.GetTemplates(tmplNames)
	if err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	companies, err := h.services.Administration.GetAllCompanies()
	if err != nil {
		h.logger.Errorf("error getting companies in admin: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	params := template_parser.TemplateParams{
		TemplateName: "companies.html",
		Vars: struct {
			LeftBar   template.HTML
			Footer    template.HTML
			Companies []models.LoanCompanyAdmin
			Domain    string
		}{
			LeftBar:   template.HTML(templates[0]),
			Footer:    template.HTML(templates[1]),
			Companies: companies,
			Domain:    viper.GetString("http.protocol"),
		},
	}

	data, err := template_parser.TemplateParser(params)
	if err != nil {
		h.logger.Errorf("error while parsing template: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.Data(http.StatusOK, "text/html", data)
}

func (h *Handler) companyRefactorPage(ctx *gin.Context) {
	id := ctx.Param("id")

	companyId, err := strconv.Atoi(id)
	if err != nil {
		newRespErr(ctx, http.StatusNotFound, "bad url")
		return
	}

	tmplNames := []string{"left_bar.html", "footer.html"}
	templates, err := template_parser.GetTemplates(tmplNames)
	if err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	company, logo, err := h.services.Administration.GetCompanyById(companyId)
	if err != nil {
		h.logger.Errorf("error while getting company: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	logotype := base64.StdEncoding.EncodeToString(logo)
	params := template_parser.TemplateParams{
		TemplateName: "company_refactor.html",
		Vars: struct {
			LeftBar template.HTML
			Footer  template.HTML
			Company models.LoanCompanyAdmin
			Logo    string
			Domain  string
		}{
			LeftBar: template.HTML(templates[0]),
			Footer:  template.HTML(templates[1]),
			Company: company,
			Logo:    logotype,
			Domain:  viper.GetString("http.protocol"),
		},
	}

	data, err := template_parser.TemplateParser(params)
	if err != nil {
		h.logger.Errorf("error while parsing template: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.Data(http.StatusOK, "text/html", data)
}

func (h *Handler) companyUploadPage(ctx *gin.Context) {
	tmplNames := []string{"left_bar.html", "footer.html"}
	templates, err := template_parser.GetTemplates(tmplNames)
	if err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	params := template_parser.TemplateParams{
		TemplateName: "company_upload.html",
		Vars: struct {
			LeftBar template.HTML
			Footer  template.HTML
			Domain  string
		}{
			LeftBar: template.HTML(templates[0]),
			Footer:  template.HTML(templates[1]),
			Domain:  viper.GetString("http.protocol"),
		},
	}

	data, err := template_parser.TemplateParser(params)
	if err != nil {
		h.logger.Errorf("error while parsing template: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.Data(http.StatusOK, "text/html", data)
}

func (h *Handler) reviewsPage(ctx *gin.Context) {
	tmplNames := []string{"left_bar.html", "footer.html"}
	templates, err := template_parser.GetTemplates(tmplNames)
	if err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	reviews, err := h.services.Administration.SelectReviews()
	if err != nil {
		h.logger.Errorf("error while getting reviews: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	params := template_parser.TemplateParams{
		TemplateName: "reviews.html",
		Vars: struct {
			LeftBar template.HTML
			Footer  template.HTML
			Reviews []models.ReviewAdmin
			Domain  string
		}{
			LeftBar: template.HTML(templates[0]),
			Footer:  template.HTML(templates[1]),
			Reviews: reviews,
			Domain:  viper.GetString("http.protocol"),
		},
	}

	data, err := template_parser.TemplateParser(params)
	if err != nil {
		h.logger.Errorf("error while parsing template: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.Data(http.StatusOK, "text/html", data)
}

func (h *Handler) reviewRefactorPage(ctx *gin.Context) {
	id := ctx.Param("id")

	reviewId, err := strconv.Atoi(id)
	if err != nil {
		newRespErr(ctx, http.StatusNotFound, "bad url")
		return
	}

	tmplNames := []string{"left_bar.html", "footer.html"}
	templates, err := template_parser.GetTemplates(tmplNames)
	if err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	review, err := h.services.Administration.GetReviewById(reviewId)
	if err != nil {
		h.logger.Errorf("error while getting review: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	params := template_parser.TemplateParams{
		TemplateName: "review_refactor.html",
		Vars: struct {
			LeftBar template.HTML
			Footer  template.HTML
			Review  models.ReviewAdmin
			Domain  string
		}{
			LeftBar: template.HTML(templates[0]),
			Footer:  template.HTML(templates[1]),
			Review:  review,
			Domain:  viper.GetString("http.protocol"),
		},
	}

	data, err := template_parser.TemplateParser(params)
	if err != nil {
		h.logger.Errorf("error while parsing template: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.Data(http.StatusOK, "text/html", data)
}

func (h *Handler) reviewUploadPage(ctx *gin.Context) {
	tmplNames := []string{"left_bar.html", "footer.html"}
	templates, err := template_parser.GetTemplates(tmplNames)
	if err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	params := template_parser.TemplateParams{
		TemplateName: "review_upload.html",
		Vars: struct {
			LeftBar template.HTML
			Footer  template.HTML
			Domain  string
		}{
			LeftBar: template.HTML(templates[0]),
			Footer:  template.HTML(templates[1]),
			Domain:  viper.GetString("http.protocol"),
		},
	}

	data, err := template_parser.TemplateParser(params)
	if err != nil {
		h.logger.Errorf("error while parsing template: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.Data(http.StatusOK, "text/html", data)
}

func (h *Handler) subscriptionsPage(ctx *gin.Context) {
	tmplNames := []string{"left_bar.html", "footer.html"}
	templates, err := template_parser.GetTemplates(tmplNames)
	if err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	subs, err := h.services.LoanCompaniesManager.GetAllSubscriptions()
	if err != nil {
		h.logger.Errorf("error while getting subscriptions: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	params := template_parser.TemplateParams{
		TemplateName: "subscriptions.html",
		Vars: struct {
			LeftBar       template.HTML
			Footer        template.HTML
			Domain        string
			Subscriptions []models.SubscriptionForAdmin
		}{
			LeftBar:       template.HTML(templates[0]),
			Footer:        template.HTML(templates[1]),
			Domain:        viper.GetString("http.protocol"),
			Subscriptions: subs,
		},
	}

	data, err := template_parser.TemplateParser(params)
	if err != nil {
		h.logger.Errorf("error while parsing template: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.Data(http.StatusOK, "text/html", data)
}

func (h *Handler) subscriptionsRefactorPage(ctx *gin.Context) {
	id := ctx.Param("id")
	subId, err := strconv.Atoi(id)
	if err != nil {
		newRespErr(ctx, http.StatusNotFound, "invalid url")
		return
	}

	tmplNames := []string{"left_bar.html", "footer.html"}
	templates, err := template_parser.GetTemplates(tmplNames)
	if err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	sub, err := h.services.LoanCompaniesManager.GetSubscriptionById(subId)
	if err != nil {
		h.logger.Errorf("error while getting subscriptions: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	params := template_parser.TemplateParams{
		TemplateName: "subscription_refactor.html",
		Vars: struct {
			LeftBar      template.HTML
			Footer       template.HTML
			Domain       string
			Subscription models.SubscriptionForAdmin
		}{
			LeftBar:      template.HTML(templates[0]),
			Footer:       template.HTML(templates[1]),
			Domain:       viper.GetString("http.protocol"),
			Subscription: sub,
		},
	}

	data, err := template_parser.TemplateParser(params)
	if err != nil {
		h.logger.Errorf("error while parsing template: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.Data(http.StatusOK, "text/html", data)
}
