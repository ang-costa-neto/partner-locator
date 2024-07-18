package router

import (
	"partner-locator/internal/app"
	"partner-locator/internal/infrastructure/http/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, app *app.Application) {
	companyHandler := handler.NewCompanyHandler(app.CompanyService)

	router.GET("/companies", companyHandler.GetCompanies)
	router.POST("/companies", companyHandler.CreateCompany)
	router.PUT("/companies/:id", companyHandler.UpdateCompany)
	router.DELETE("/companies/:id", companyHandler.DeleteCompany)
}
