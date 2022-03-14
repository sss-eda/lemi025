package text

import (
	"context"
	"time"
)

// Service TODO
type Service interface {
	GetFileByTime(context.Context, time.Time) (*File, error)
}

// ReadRepository TODO
type ReadRepository interface {
	GetFile(context.Context, string) (*File, error)
}

// WriteRepository TODO
type WriteRepository interface {
	
}

type service struct {
	repo Repository
}

// NewService TODO
func NewReadService(
	repository Repository,
) (Service, error) {
	return &service{
		repo: repository,
	}, nil
}

// GetFile TODO
func (svc *service) GetFileByTime(
	ctx context.Context,
	time time.Time,
) (*File, error) {
	return &File{}, nil
}
