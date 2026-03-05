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
	err := s.repository.Create(openingRequest)
	if err != nil {
		logger.ErrorF("Create opening  error: %v", err)
		return nil, err
	}
	return openingRequest, nil
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
	newOpening, err := s.repository.Update(*openingRequest)
	if err != nil {
		logger.ErrorF("Update opening %v error: %v", openingRequest.Name, err)
		return nil, err
	}
	return &newOpening, nil
}
func (s *openingService) DeleteOpening(id string) error {
	logger.InforF("Delete opening %v", id)
	_, err := s.GetOpening(id)
	if err != nil {
		logger.ErrorF("Erro: Vaga %v não encontrada para exclusão", id)
		return err
	}
	err = s.repository.Delete(id)
	if err != nil {
		logger.ErrorF("Delete opening %v error: %v", id, err)
		return err
	}
	return nil
}

func (s *openingService) ListOpenings() ([]*schemas.Opening, error) {
	logger.InforF("List openings")
	openings, err := s.repository.List()
	if err != nil {
		logger.ErrorF("List openings error: %v", err)
		return nil, err
	}
	return openings, nil
}
