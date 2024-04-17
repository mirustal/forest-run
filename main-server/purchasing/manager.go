package purchasing

import "main-server/domain"

type Manager interface {
	ValidateRunPermissionsTransaction(id domain.TransactionId) (domain.RunPermissionsType, error)
}

type manager struct {
}

func NewManager() Manager {
	return &manager{}
}

func (m manager) ValidateRunPermissionsTransaction(id domain.TransactionId) (domain.RunPermissionsType, error) {
	return domain.RunPermissionsType(10), nil
}
