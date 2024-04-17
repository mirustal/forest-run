package database

import (
	"context"
	"main-server/domain"
)

type RunsRepo interface {
	StoreRun(run domain.Run, ctx context.Context) (domain.Run, error)
	UpdateRun(run domain.Run, ctx context.Context) error
	GetUserCreatedRuns(userId string, ctx context.Context) ([]domain.RunId, error)
	GetUserParticipatedRuns(userId string, ctx context.Context) ([]domain.RunId, error)
}

func (p PgDbAdapter) StoreRun(run domain.Run, ctx context.Context) (domain.Run, error) {
	//TODO implement me
	panic("implement me")
}

func (p PgDbAdapter) UpdateRun(run domain.Run, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p PgDbAdapter) GetUserCreatedRuns(userId string, ctx context.Context) ([]domain.RunId, error) {
	//TODO implement me
	panic("implement me")
}

func (p PgDbAdapter) GetUserParticipatedRuns(userId string, ctx context.Context) ([]domain.RunId, error) {
	//TODO implement me
	panic("implement me")
}
