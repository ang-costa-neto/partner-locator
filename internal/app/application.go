package app

import (
	"partner-locator/internal/domain/service"
	"partner-locator/internal/infrastructure/db"
	"partner-locator/internal/infrastructure/repository"

	"gorm.io/gorm"
)

type Application struct {
	CompanyService service.CompanyService
	UserService    service.UserService
	DB             *gorm.DB
}

func NewApplication() *Application {
	database := db.NewDB()

	companyRepo := repository.NewCompanyRepository(database)
	companyService := service.NewCompanyService(companyRepo)

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)

	return &Application{
		CompanyService: companyService,
		UserService:    userService,
		DB:             database,
	}
}
