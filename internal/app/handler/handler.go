package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "zaimik/docs"
	"zaimik/internal/app/service"
	"zaimik/internal/pkg/logging"
)

type Handler struct {
	services *service.Service
	logger   *logging.Logger
}

func NewHandler(services *service.Service, logger *logging.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func CorsSettings() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedOrigins:     []string{viper.GetString("http.protocol")},
		AllowCredentials:   true,
		AllowedHeaders:     []string{},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{},
		Debug:              true,
	})

	return c
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("templates/public/*.html")
	router.Static("/static/js", "templates/js")

	router.Use(gin.Recovery())
	router.Use(gin.ErrorLogger())

	// TODO фейк эндпоинт (надо будет удалить)
	router.GET("/", h.indexPage)

	router.POST("/get-code/", h.getAuthCode)
	router.POST("/", h.signIn)
	router.GET("/reviews/", h.getReviews)

	api := router.Group("/api", h.checkSession)
	{
		// TODO сделать проверку статуса подписки
		api.GET("/check-subscription/", h.checkSubscription)

		// TODO сделать тут отдачу ссылки и добавить после подписки отдачу компаний
		api.PUT("/catalog/", h.updateUserGetLinkOnPayment)
		api.POST("/catalog/", h.getSortedSuitableCatalog)
		api.POST("/sign-out/", h.signOut)

		api.POST("/send-review/", h.sendReview)

		// TODO фейковый эндпоинт (надо будет удалить)
		api.GET("/catalog/", h.chatPage)

	}

	admin := router.Group("/admin")
	{
		admin.GET("/", h.authPage)
		admin.POST("/get-code/", h.checkAdmin)
		admin.POST("/", h.authAdmin)

		adminApi := admin.Group("/panel", h.checkAdminSession)
		{
			adminApi.GET("/", h.panelPage)
			adminApi.POST("/sign-out/", h.signOutAdmin)

			companies := adminApi.Group("/companies")
			{
				companies.GET("/", h.companiesPage)
				companies.POST("/update/", h.updateCompanies)

				companies.GET("/:id/", h.companyRefactorPage)
				companies.PUT("/:id/", h.refactorCompany)
				companies.DELETE("/:id/", h.deleteCompany)
				companies.GET("/upload/", h.companyUploadPage)
				companies.POST("/upload/", h.uploadCompany)

			}

			reviews := adminApi.Group("/reviews")
			{
				reviews.GET("/", h.reviewsPage)
				reviews.GET("/:id/", h.reviewRefactorPage)
				reviews.PUT("/update/", h.updateReview)
				reviews.GET("/upload/", h.reviewUploadPage)
				reviews.POST("/upload/", h.uploadReview)
				reviews.DELETE("/delete/", h.deleteReview)
			}

			subscriptions := adminApi.Group("/subscriptions")
			{
				subscriptions.GET("/", h.subscriptionsPage)
				subscriptions.GET("/:id/", h.subscriptionsRefactorPage)
				subscriptions.PUT("/:id/", h.refactorSubscription)
			}
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
