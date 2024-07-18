package repository

import (
	"partner-locator/internal/domain/model"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return CompanyRepository{db: db}
}

func (r *CompanyRepository) GetCompanies() ([]model.Company, error) {
	var companies []model.Company
	if err := r.db.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (r *CompanyRepository) CreateCompany(company model.Company) error {
	if err := r.db.Create(&company).Error; err != nil {
		return err
	}
	return nil
}

func (r *CompanyRepository) UpdateCompany(company model.Company) error {
	if err := r.db.Save(&company).Error; err != nil {
		return err
	}
	return nil
}

func (r *CompanyRepository) DeleteCompany(id uint) error {
	if err := r.db.Delete(&model.Company{}, id).Error; err != nil {
		return err
	}
	return nil
}
