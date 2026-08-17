package application

import (
	"context"

	"GO-Crud/internal/modules/user/domain"
	"GO-Crud/internal/modules/user/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(
	ctx context.Context,
	user *domain.User,
) error {
	return s.repo.Create(ctx, user)
}

func (s *UserService) GetByID(
	ctx context.Context,
	id int64,
) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) GetAll(
	ctx context.Context,
) ([]domain.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) Update(
	ctx context.Context,
	user *domain.User,
) error {
	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(
	ctx context.Context,
	id int64,
) error {
	return s.repo.Delete(ctx, id)
}
