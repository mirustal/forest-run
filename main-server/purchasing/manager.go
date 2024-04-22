package purchasing

import (
	"context"
	"main-server/domain"
)

type Manager interface {
	ValidateRunPermissionsTransaction(id domain.TransactionId, permissions domain.RunPermissionsType, ctx context.Context) error
}

type manager struct {
}

func NewManager() Manager {
	return &manager{}
}

func (m manager) ValidateRunPermissionsTransaction(id domain.TransactionId, permissions domain.RunPermissionsType, ctx context.Context) error {
	return nil
}
