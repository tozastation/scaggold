package repositories

import (
	"context"
	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in generics_repository.go  -out user_repository.go gen "KeyType=User"

// KeyType is ...
type KeyType generic.Type

// ValueType is ...
//ype ValueType generic.Type

// KeyTypeRepository is ...
type KeyTypeRepository interface {
	Insert(ctx context.Context) error
	SelectByID(ctx context.Context) (*KeyType, error)
	SelectAll(ctx context.Context) ([]*KeyType, error)
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}

// KeyTypeRepositoryImpl is ...
type KeyTypeRepositoryImpl struct{}

// NewKeyTypeRepositoryImpl is ...
func NewKeyTypeRepositoryImpl() KeyTypeRepository {
	return &KeyTypeRepositoryImpl{}
}

// Insert is ...
func (repo *KeyTypeRepositoryImpl) Insert(ctx context.Context) error {
	return nil
}

// SelectByID is ...
func (repo *KeyTypeRepositoryImpl) SelectByID(ctx context.Context) (*KeyType, error) {
	return nil, nil
}

// SelectAll is ...
func (repo *KeyTypeRepositoryImpl) SelectAll(ctx context.Context) ([]*KeyType, error) {
	return nil, nil
}

// Update is ...
func (repo *KeyTypeRepositoryImpl) Update(ctx context.Context) error {
	return nil
}

// Delete is ...
func (repo *KeyTypeRepositoryImpl) Delete(ctx context.Context) error {
	return nil
}
