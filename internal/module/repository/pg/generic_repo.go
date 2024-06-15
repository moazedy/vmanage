package pg

import (
	"context"
	"errors"
	"fmt"
	"vmanage/pkg/infra/tx"
	"vmanage/pkg/module/vmanage/model/entity"
	"vmanage/pkg/module/vmanage/repository"

	"gorm.io/gorm"
)

const (
	fieldNameID = "id"
)

type genericRepoFactory[E entity.Entity] struct{}

func NewGenericRepoFactory[E entity.Entity]() repository.GenericRepoFactory[E] {
	return genericRepoFactory[E]{}
}

func (grf genericRepoFactory[E]) NewGenericRepo(tx tx.TX) repository.GenericRepo[E] {
	return newGenericRepo[E](tx.GetConnection())
}

type genericRepo[E entity.Entity] struct {
	db *gorm.DB
}

func newGenericRepo[E entity.Entity](db *gorm.DB) repository.GenericRepo[E] {
	if db == nil {
		panic("no db connection provided for credentials")
	}

	return genericRepo[E]{
		db: db,
	}
}

func (gr genericRepo[E]) Create(ctx context.Context, entityData E) (*E, error) {
	err := gr.db.Create(&entityData).Error
	if err != nil {
		return nil, err
	}

	return &entityData, nil
}

func (gr genericRepo[E]) Update(ctx context.Context, entityData E) (*E, error) {
	err := gr.db.Save(&entityData).Error
	if err != nil {
		return nil, err
	}

	return &entityData, nil
}

func (gr genericRepo[E]) Delete(ctx context.Context, entityID string) error {
	return gr.db.Where(fmt.Sprintf("%s = ?", fieldNameID), entityID).Delete(new(E)).Error
}

func (gr genericRepo[E]) GetByStringField(ctx context.Context, fieldName, fieldValue string) (*E, error) {
	entityInstance := new(E)
	err := gr.db.Where(fmt.Sprintf("%s = ?", fieldName), fieldValue).First(entityInstance).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return entityInstance, nil
}

func (gr genericRepo[E]) List(context.Context) ([]E, error) {
	entityList := make([]E, 0)
	err := gr.db.Find(&entityList).Error
	if err != nil {
		return nil, err
	}

	return entityList, nil
}
