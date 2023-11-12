package repository

import (
	"fmt"
	"platter/internal/model"
)

var ErrAlreadyExists = fmt.Errorf("already exists")

type Repository[T model.Model] interface {
	Create(model *T) error
}
