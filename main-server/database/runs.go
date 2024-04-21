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
	t, err := p.dbPool.Begin(ctx)

	defer func() {
		if err != nil {
			_ = t.Rollback(ctx)
		}
	}()

	if err != nil {
		return run, err
	}

	err = t.QueryRow(ctx, `INSERT INTO runs (
                  name,
                  description,
                  official_site,
                  avatar_url,
                  route,
                  start_time,
                  start_place,
                  start_place_latitude,
                  start_place_longitude,
                  max_participants,
                  participation_format) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`,
		run.Name,
		run.Description,
		run.OfficialSiteUrl,
		run.AvatarUrl,
		run.Route,
		run.StartTime,
		run.StartPlace.Address,
		run.StartPlace.Point.Latitude,
		run.StartPlace.Point.Longitude,
		run.MaxParticipants,
		run.ParticipationFormat).
		Scan(&run.Id)

	if err != nil {
		return run, err
	}

	return run, t.Commit(ctx)
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
