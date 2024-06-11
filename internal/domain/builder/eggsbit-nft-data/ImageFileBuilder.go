package eggsbitnftdata

import (
	"errors"
)

func NewImageFileBuilder() ImageFileBuilderInterface {
	return &ImageFileBuilder{}
}

type ImageFileBuilderInterface interface {
	CreateStartingEggImage(imageUuid string, eggPattern string, eggColorScheme string) error
}

type ImageFileBuilder struct {
}

func (ifb ImageFileBuilder) CreateStartingEggImage(imageUuid string, eggPattern string, eggColorScheme string) error {
	// repository with pattern
	// repository with color

	// get file
	// replace some lines - https://stackoverflow.com/questions/26152901/replace-a-line-in-text-file-golang
	// export to the folder from config path - tmp
	// export to real folder by inskape

	// * rename color in the files to keys from the db
	return errors.New("create random starting egg image error")
}
