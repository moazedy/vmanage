package repository

import (
	"context"
	"vmanage/pkg/infra/tx"
	"vmanage/pkg/module/vmanage/model/entity"
)

type GenericRepo[E entity.Entity] interface {
	Create(context.Context, E) (*E, error)
	Update(context.Context, E) (*E, error)
	Delete(context.Context, string) error
	GetByStringField(ctx context.Context, fieldName, fieldValue string) (*E, error)
	List(context.Context) ([]E, error)
}

type GenericRepoFactory[E entity.Entity] interface {
	NewGenericRepo(tx tx.TX) GenericRepo[E]
}
