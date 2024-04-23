package database

import (
	"context"
	"encoding/json"
	"forest-run/common/runs"
	"forest-run/main-server/domain"
)

type RunsRepo interface {
	StoreRun(run domain.Run, ctx context.Context) (domain.Run, error)
	UpdateRun(run domain.Run, ctx context.Context) error
	GetUserCreatedRuns(userId string, ctx context.Context) ([]runs.Id, error)
	GetUserParticipatedRuns(userId string, ctx context.Context) ([]runs.Id, error)
	GetRun(runId runs.Id, ctx context.Context) (run domain.Run, err error)
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

	routeJson, err := json.Marshal(run.Route)
	if err != nil {
		return run, err
	}
	err = t.QueryRow(ctx, `INSERT INTO runs (
                  name,
                  creator,
                  description,
                  official_site,
                  avatar_url,
                  route,
                  start_time,
                  start_place,
                  start_place_latitude,
                  start_place_longitude,
                  max_participants,
                  participation_format,
				  registration_until,
                  permissions_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id`,
		run.Name,
		run.Creator,
		run.Description,
		run.OfficialSiteUrl,
		run.AvatarUrl,
		routeJson,
		run.StartTime,
		run.StartPlace.Address,
		run.StartPlace.Point.Latitude,
		run.StartPlace.Point.Longitude,
		run.MaxParticipants,
		run.ParticipationFormat,
		run.RegistrationUntil,
		run.RunPermissions).
		Scan(&run.Id)

	if err != nil {
		return run, err
	}

	return run, t.Commit(ctx)
}

func (p PgDbAdapter) UpdateRun(run domain.Run, ctx context.Context) error {
	t, err := p.dbPool.Begin(ctx)

	defer func() {
		if err != nil {
			_ = t.Rollback(ctx)
		}
	}()

	if err != nil {
		return err
	}

	routeJson, err := json.Marshal(run.Route)
	if err != nil {
		return err
	}

	err = t.QueryRow(ctx, `UPDATE runs SET  
                  name = $2,
                  description = $3,
                  official_site = $4,
                  avatar_url = $5,
                  route = $6,
                  start_time = $7,
                  start_place = $8,
                  start_place_latitude = $9,
                  start_place_longitude = $10,
                  max_participants = $11,
                  participation_format = $12,
				  status = $13,
				  registration_until = $14,
				  permissions_type = $15
            WHERE id = $1`,
		run.Id,
		run.Name,
		run.Description,
		run.OfficialSiteUrl,
		run.AvatarUrl,
		routeJson,
		run.StartTime,
		run.StartPlace.Address,
		run.StartPlace.Point.Latitude,
		run.StartPlace.Point.Longitude,
		run.MaxParticipants,
		run.ParticipationFormat,
		run.RunPermissions).
		Scan(&run.Id)

	if err != nil {
		return err
	}

	return t.Commit(ctx)
}

func (p PgDbAdapter) GetRun(runId runs.Id, ctx context.Context) (run domain.Run, err error) {
	t, err := p.dbPool.Begin(ctx)

	defer func() {
		if err != nil {
			_ = t.Rollback(ctx)
		}
	}()

	var routeJson string
	err = t.QueryRow(ctx, `SELECT 
								id,
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
								status,
								participation_format,
								registration_until,
								permissions_type,
								creator
						FROM runs WHERE id = $1`, runId).
		Scan(
			&run.Id,
			&run.Name,
			&run.Description,
			&run.OfficialSiteUrl,
			&run.AvatarUrl,
			&routeJson,
			&run.StartTime,
			&run.StartPlace.Address,
			&run.StartPlace.Point.Latitude,
			&run.StartPlace.Point.Longitude,
			&run.MaxParticipants,
			&run.Status,
			&run.ParticipationFormat,
			&run.RegistrationUntil,
			&run.RunPermissions,
			&run.Creator)

	if err != nil {
		return run, err
	}

	err = json.Unmarshal([]byte(routeJson), &run.Route)
	if err != nil {
		return run, err
	}

	return run, t.Commit(ctx)
}

func (p PgDbAdapter) GetUserCreatedRuns(userId string, ctx context.Context) ([]runs.Id, error) {
	//TODO implement me
	panic("implement me")
}

func (p PgDbAdapter) GetUserParticipatedRuns(userId string, ctx context.Context) ([]runs.Id, error) {
	//TODO implement me
	panic("implement me")
}
