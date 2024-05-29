package service

import (
	"my-reading-app/internal/domain"
	"my-reading-app/internal/repository"
)

type ReadingStatusService struct {
	repo repository.ReadingStatusRepository
}

func NewReadingStatusService(repo repository.ReadingStatusRepository) *ReadingStatusService {
	return &ReadingStatusService{repo: repo}
}

func (s *ReadingStatusService) GetStatus(userId string) ([]domain.ReadingStatus, error) {
	return s.repo.GetStatus(userId)
}

func (s *ReadingStatusService) UpdateStatus(userId string, day int, status string) error {
	return s.repo.UpdateStatus(userId, day, status)
}
