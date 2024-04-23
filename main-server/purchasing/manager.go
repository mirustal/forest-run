package purchasing

import (
	"context"
	"forest-run/common/runs"
	"forest-run/main-server/domain"
)

type Manager interface {
	ValidateRunPermissionsTransaction(id domain.TransactionId, permissions runs.PermissionsType, ctx context.Context) error
}

type manager struct {
}

func NewManager() Manager {
	return &manager{}
}

func (m manager) ValidateRunPermissionsTransaction(id domain.TransactionId, permissions runs.PermissionsType, ctx context.Context) error {
	return nil
}
