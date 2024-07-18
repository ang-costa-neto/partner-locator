package service

import (
	"partner-locator/internal/domain/model"
	"partner-locator/internal/infrastructure/repository"
)

type CompanyService struct {
	repo repository.CompanyRepository
}

func NewCompanyService(repo repository.CompanyRepository) CompanyService {
	return CompanyService{repo: repo}
}

func (s CompanyService) GetCompanies() ([]model.Company, error) {
	return s.repo.GetCompanies()
}

func (s CompanyService) CreateCompany(company model.Company) error {
	return s.repo.CreateCompany(company)
}

func (s CompanyService) UpdateCompany(company model.Company) error {
	return s.repo.UpdateCompany(company)
}

func (s CompanyService) DeleteCompany(id uint) error {
	return s.repo.DeleteCompany(id)
}
