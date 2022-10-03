package service

import (
	"context"

	"github.com/aryuuu/finkita/domain"
)

type mutationService struct {
	mutationRepo domain.IMutationRepository
}

func NewMutationService(mr domain.IMutationRepository) domain.IMutationService {
	return &mutationService{
		mutationRepo: mr,
	}
}

func (s mutationService) AddMutation(ctx context.Context, mutation *domain.Mutation) error {
	return nil
}

func (s mutationService) GetMutations(ctx context.Context) ([]domain.Mutation, error) {
	return s.mutationRepo.GetMutations(ctx, map[string]interface{}{})
}

func (s mutationService) GetMutationsByEmail(ctx context.Context, email string) ([]domain.Mutation, error) {
	filter := map[string]interface{}{
		"email": email,
	}

	return s.mutationRepo.GetMutations(ctx, filter)
}

func (s mutationService) GetMutationByID(ctx context.Context, id string) (domain.Mutation, error) {
	return domain.Mutation{}, nil
}

func (s mutationService) UpdateMutationByID(ctx context.Context, id string, mutation domain.Mutation) (domain.Mutation, error) {
	return domain.Mutation{}, nil
}

func (s mutationService) DeleteMutation(ctx context.Context, id string) (domain.Mutation, error) {
	return domain.Mutation{}, nil
}
