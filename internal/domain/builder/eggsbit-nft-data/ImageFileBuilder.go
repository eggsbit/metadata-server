package eggsbitnftdata

import (
	"errors"
)

func NewImageFileBuilder() ImageFileBuilderInterface {
	return &ImageFileBuilder{}
}

type ImageFileBuilderInterface interface {
	CreateRandomStartingEggImage() (*string, error)
}

type ImageFileBuilder struct {
}

func (ifb ImageFileBuilder) CreateRandomStartingEggImage() (*string, error) {
	return nil, errors.New("create random starting egg image error")
}
