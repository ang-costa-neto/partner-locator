package app

import (
	"partner-locator/internal/domain/service"
	"partner-locator/internal/infrastructure/db"
	"partner-locator/internal/infrastructure/repository"

	"gorm.io/gorm"
)

type Application struct {
	CompanyService service.CompanyService
	DB             *gorm.DB
}

func NewApplication() *Application {
	database := db.NewDB()

	companyRepo := repository.NewCompanyRepository(database)
	companyService := service.NewCompanyService(companyRepo)

	return &Application{
		CompanyService: companyService,
		DB:             database,
	}
}
