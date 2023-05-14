package usecase

import (
	"context"
	"mnezerka/MySpots/server/domain"
	"time"
)

type spotsUsecase struct {
	spotsRepository domain.SpotsRepository
	contextTimeout  time.Duration
}

func NewSpotsUsecase(spotsRepository domain.SpotsRepository, timeout time.Duration) domain.SpotsUsecase {
	return &spotsUsecase{
		spotsRepository: spotsRepository,
		contextTimeout:  timeout,
	}
}

func (su *spotsUsecase) Create(c context.Context, spot *domain.Spot) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.spotsRepository.Create(ctx, spot)
}

func (su *spotsUsecase) Fetch(c context.Context) ([]domain.Spot, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.spotsRepository.Fetch(ctx)
}
