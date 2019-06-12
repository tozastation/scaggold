package services

import (
	"context"
	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in generics_service.go  -out user_service.go gen "KeyType=User"

// KeyType is ...
type KeyType generic.Type

// ValueType is ...
//ype ValueType generic.Type

// KeyTypeService is ...
type KeyTypeService interface {
	Add(ctx context.Context) error
	GetByID(ctx context.Context) (*KeyType, error)
	GetAll(ctx context.Context) ([]*KeyType, error)
	Put(ctx context.Context) error
	Delete(ctx context.Context) error
}

// KeyTypeServiceImpl is ...
type KeyTypeServiceImpl struct{}

// NewKeyTypeServiceImpl is ...
func NewKeyTypeServiceImpl() KeyTypeService {
	return &KeyTypeServiceImpl{}
}

// Add is ...
func (repo *KeyTypeServiceImpl) Add(ctx context.Context) error {
	return nil
}

// GetByID is ...
func (repo *KeyTypeServiceImpl) GetByID(ctx context.Context) (*KeyType, error) {
	return nil, nil
}

// GetAll is ...
func (repo *KeyTypeServiceImpl) GetAll(ctx context.Context) ([]*KeyType, error) {
	return nil, nil
}

// Put is ...
func (repo *KeyTypeServiceImpl) Put(ctx context.Context) error {
	return nil
}

// Delete is ...
func (repo *KeyTypeServiceImpl) Delete(ctx context.Context) error {
	return nil
}
