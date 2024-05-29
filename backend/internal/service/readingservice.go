package service

import (
	"my-reading-app/internal/domain"
	"my-reading-app/internal/repository"
	"strconv"
)

type ReadingService interface {
	GetReading(day string) (*domain.Reading, error)
}

type readingService struct {
	repo repository.ReadingRepository
}

func NewReadingService(repo repository.ReadingRepository) ReadingService {
	return &readingService{repo: repo}
}

func (s *readingService) GetReading(day string) (*domain.Reading, error) {
	dayInt, err := strconv.Atoi(day)
	if err != nil {
		return nil, err
	}
	return s.repo.GetReadingByDay(dayInt)
}
