package repository

import (
	"github.com/Aaron-GMM/govagas/handler"
	"github.com/Aaron-GMM/govagas/schemas"
	"gorm.io/gorm"
)

type OpeningRepository interface {
	CREATE(opening schemas.Opening) (schemas.Opening, error)
	LIST() ([]schemas.Opening, error)
	FindById(id string) (schemas.Opening, error)
	DELETE(id string) error
	UPDATE(opening schemas.Opening) (schemas.Opening, error)
}

type openingRepository struct {
	db *gorm.DB
}

func NewOpeningRepository(db *gorm.DB) OpeningRepository {
	return &openingRepository{db: db}
}
func (r *openingRepository) CREATE(opening schemas.Opening) (schemas.Opening, error) {
	err := r.db.Create(&opening).Error
	if err != nil {
		handler.Logger.ErrorF("Create opening %v", err.Error())
		return opening, err
	}
	return opening, nil
}

func (r *openingRepository) LIST() ([]schemas.Opening, error) {
	var openings []schemas.Opening
	err := r.db.Find(&openings).Error
	if err != nil {
		handler.Logger.ErrorF("Find opening %v", err.Error())
		return openings, err
	}
	return openings, nil
}
func (r *openingRepository) FindById(id string) (schemas.Opening, error) {
	var opening schemas.Opening
	err := r.db.First(&opening, id).Error
	if err != nil {
		handler.Logger.ErrorF("Find opening %v", err.Error())
		return opening, err
	}
	return opening, nil
}
func (r *openingRepository) DELETE(id string) error {
	err := r.db.Delete(&schemas.Opening{}, id).Error
	if err != nil {
		handler.Logger.ErrorF("Delete opening %v", err.Error())
		return err
	}
	return nil
}
func (r *openingRepository) UPDATE(opening schemas.Opening) (schemas.Opening, error) {
	err := r.db.Save(&opening).Error
	if err != nil {
		handler.Logger.ErrorF("Update opening %v", err.Error())
		return opening, err
	}
	return opening, nil
}
