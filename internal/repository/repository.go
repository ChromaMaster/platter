package repository

import (
	"fmt"
	"platter/internal/model"
)

var (
	ErrAlreadyExists = fmt.Errorf("already exists")
	ErrNotExists     = fmt.Errorf("not exists")
)

type Repository[T model.Model] interface {
	Init() error

	Create(model *T) error
	GetAll() ([]*T, error)
	Remove(ID int) error
}
