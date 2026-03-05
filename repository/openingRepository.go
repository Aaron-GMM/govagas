package repository

import (
	"github.com/Aaron-GMM/govagas/config"
	"github.com/Aaron-GMM/govagas/schemas"
	"gorm.io/gorm"
)

type OpeningRepository interface {
	Create(opening *schemas.Opening) error
	List() ([]*schemas.Opening, error)
	FindById(id string) (schemas.Opening, error)
	Delete(id string) error
	Update(opening schemas.Opening) (schemas.Opening, error)
}

type openingRepository struct {
	db *gorm.DB
}

var Logger = config.GetLogger("repository")

func NewOpeningRepository(db *gorm.DB) OpeningRepository {
	return &openingRepository{db: db}
}
func (r *openingRepository) Create(opening *schemas.Opening) error {
	if err := r.db.Create(opening).Error; err != nil {
		Logger.ErrorF("Create opening %v", err.Error())
		return err
	}
	return nil
}

func (r *openingRepository) List() ([]*schemas.Opening, error) {
	var openings []*schemas.Opening
	err := r.db.Find(&openings).Error
	if err != nil {
		Logger.ErrorF("Find opening %v", err.Error())
		return nil, err
	}
	return openings, nil
}
func (r *openingRepository) FindById(id string) (schemas.Opening, error) {
	var opening schemas.Opening
	err := r.db.First(&opening, id).Error
	if err != nil {
		Logger.ErrorF("Find opening %v", err.Error())
		return opening, err
	}
	return opening, nil
}
func (r *openingRepository) Delete(id string) error {
	err := r.db.Delete(&schemas.Opening{}, id).Error
	if err != nil {
		Logger.ErrorF("Delete opening %v", err.Error())
		return err
	}
	return nil
}
func (r *openingRepository) Update(opening schemas.Opening) (schemas.Opening, error) {
	err := r.db.Save(&opening).Error
	if err != nil {
		Logger.ErrorF("Update opening %v", err.Error())
		return opening, err
	}
	return opening, nil
}
