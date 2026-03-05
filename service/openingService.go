package service

import (
	"github.com/Aaron-GMM/govagas/config"
	"github.com/Aaron-GMM/govagas/repository"
	"github.com/Aaron-GMM/govagas/schemas"
)

type OpeningService interface {
	CreateOpening(opening *schemas.Opening) (*schemas.Opening, error)
	GetOpening(id string) (*schemas.Opening, error)
	UpdateOpening(opening *schemas.Opening) (*schemas.Opening, error)
	DeleteOpening(id string) error
	ListOpenings() ([]*schemas.Opening, error)
}

type openingService struct {
	repository repository.OpeningRepository
}

func NewOpeningService(repository repository.OpeningRepository) OpeningService {
	return &openingService{repository: repository}
}

var logger = config.GetLogger("OpeningService")

func (s *openingService) CreateOpening(openingRequest *schemas.Opening) (*schemas.Opening, error) {
	logger.InforF("Create opening %v", openingRequest.Name)
	opening, err := s.repository.CREATE(*openingRequest)
	if err != nil {
		logger.ErrorF("Create opening %v error: %v", opening.ID, err)
		return nil, err
	}
	return &opening, nil
}

func (s *openingService) GetOpening(id string) (*schemas.Opening, error) {
	logger.InforF("Get opening %v", id)
	opening, err := s.repository.FindById(id)
	if err != nil {
		logger.ErrorF("Get opening %v error: %v", id, err)
		return nil, err
	}

	return &opening, nil
}

func (s *openingService) UpdateOpening(openingRequest *schemas.Opening) (*schemas.Opening, error) {
	logger.InforF("Update opening %v", openingRequest.Name)
	newOpening, err := s.repository.UPDATE(*openingRequest)
	if err != nil {
		logger.ErrorF("Update opening %v error: %v", openingRequest.Name, err)
		return nil, err
	}
	return &newOpening, nil
}
func (s *openingService) DeleteOpening(id string) error {
	logger.InforF("Delete opening %v", id)
	err := s.repository.DELETE(id)
	if err != nil {
		logger.ErrorF("Delete opening %v error: %v", id, err)
		return err
	}
	return nil
}

func (s *openingService) ListOpenings() ([]*schemas.Opening, error) {
	logger.InforF("List openings")
	openings, err := s.repository.LIST()
	if err != nil {
		logger.ErrorF("List openings error: %v", err)
		return nil, err
	}
	return openings, nil
}
