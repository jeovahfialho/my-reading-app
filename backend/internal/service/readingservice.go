package service

import (
	"my-reading-app/internal/domain"
	"my-reading-app/internal/repository"
)

type ReadingService interface {
	GetReadingByDay(day int) (domain.Reading, error)
}

type readingService struct {
	repo repository.ReadingRepository
}

func NewReadingService(repo repository.ReadingRepository) ReadingService {
	return &readingService{repo: repo}
}

func (s *readingService) GetReadingByDay(day int) (domain.Reading, error) {
	return s.repo.GetReadingByDay(day)
}
