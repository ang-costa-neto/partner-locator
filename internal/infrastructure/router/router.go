package router

import (
	"partner-locator/internal/app"
	"partner-locator/internal/infrastructure/http/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, app *app.Application) {
	companyHandler := handler.NewCompanyHandler(app.CompanyService)
	userHandler := handler.NewUserHandler(app.UserService)

	// Companies URL
	router.GET("/companies", companyHandler.GetCompanies)
	router.POST("/companies", companyHandler.CreateCompany)
	router.PUT("/companies/:id", companyHandler.UpdateCompany)
	router.DELETE("/companies/:id", companyHandler.DeleteCompany)

	// User URL
	router.GET("/users", userHandler.GetUsers)
	router.POST("/users", userHandler.CreateUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
}
