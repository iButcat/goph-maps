package service

import (
	"context"

	"github.com/iButcat/repository"
)

type Service interface {
	Save(ctx context.Context)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Save(ctx context.Context) {
	var test struct {
		Test string
	}
	test.Test = "hello"
	s.repository.Create(ctx, &test)
}
