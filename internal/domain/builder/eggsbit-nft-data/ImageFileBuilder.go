package eggsbitnftdata

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/eggsbit/metadata-server/configs"
	"github.com/eggsbit/metadata-server/internal/domain/repository"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"
)

func NewImageFileBuilder(
	imagePatternRepository repository.ImagePatternDocRepositoryInterface,
	colorSchemeRepository repository.ColorSchemeDocRepositoryInterface,
	logger log.LoggerInterface,
	config *configs.Config,
) ImageFileBuilderInterface {
	return &ImageFileBuilder{
		imagePatternRepository: imagePatternRepository,
		colorSchemeRepository:  colorSchemeRepository,
		logger:                 logger,
		config:                 config,
	}
}

type ImageFileBuilderInterface interface {
	CreateStartingEggImage(imageUuid string, eggPattern string, eggColorScheme string, ctx context.Context) error
}

type ImageFileBuilder struct {
	imagePatternRepository repository.ImagePatternDocRepositoryInterface
	colorSchemeRepository  repository.ColorSchemeDocRepositoryInterface
	logger                 log.LoggerInterface
	config                 *configs.Config
}

func (ifb ImageFileBuilder) CreateStartingEggImage(imageUuid string, eggPattern string, eggColorScheme string, ctx context.Context) error {
	imagePattern, _ := ifb.imagePatternRepository.GetImagePatternByIdentifier(eggPattern, ctx)
	colorScheme, _ := ifb.colorSchemeRepository.GetColorSchemeByIdentifier(eggColorScheme, ctx)

	inputPatternFile, err := os.ReadFile(imagePattern.Path)
	if err != nil {
		ifb.logger.Error(log.LogCategorySystem, err.Error())
	}

	lines := strings.Split(string(inputPatternFile), "\n")
	for i, line := range lines {
		//берем строку и ищем в ней паттерн \[\[(color_\d+)\]\]
		//если нашелся то вычленяем его и используем в качестве ключа для того чтобы взять цвет из colorScheme.colors
		//подставляем вместо найденого ранее значения цвет и идем дальше
		// if strings.Contains(line, "]") {
		// 		lines[i] = "LOL"
		// }
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile("myfile", []byte(output), 0644)

	if err != nil {
		ifb.logger.Error(log.LogCategorySystem, err.Error())
	}

	// get file
	// replace some lines - https://stackoverflow.com/questions/26152901/replace-a-line-in-text-file-golang
	// export to the folder from config path - tmp
	// export to real folder by inskape

	return errors.New("create random starting egg image error")
}
