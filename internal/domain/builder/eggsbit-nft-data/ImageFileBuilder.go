package eggsbitnftdata

import (
	"context"
	"os"
	"os/exec"
	"regexp"
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

func (ifb ImageFileBuilder) CreateStartingEggImage(
	imageUuid string,
	eggPatternIdentifier string,
	eggColorSchemeIdentifier string,
	ctx context.Context,
) error {
	imagePattern, _ := ifb.imagePatternRepository.GetImagePatternByIdentifier(eggPatternIdentifier, ctx)
	colorScheme, _ := ifb.colorSchemeRepository.GetColorSchemeByIdentifier(eggColorSchemeIdentifier, ctx)

	inputPatternFile, errReadFile := os.ReadFile(imagePattern.Path)
	if errReadFile != nil {
		return errReadFile
	}

	lines := strings.Split(string(inputPatternFile), "\n")
	findSubstringRegexp := regexp.MustCompile(`\[\[(color_\d+)\]\]`)
	for lineIndex, line := range lines {
		match := findSubstringRegexp.FindStringSubmatch(line)

		if len(match) == 0 {
			continue
		}

		colorKey := match[1]
		replaceSubstringRegexp := regexp.MustCompile(`(\[\[` + colorKey + `\]\])`)
		lines[lineIndex] = replaceSubstringRegexp.ReplaceAllString(line, colorScheme.Colors[colorKey])
	}
	output := strings.Join(lines, "\n")

	mkdirExportCmd := exec.Command("mkdir", "-p", ifb.config.ApplicationConfig.ExportFolderPath)
	if mkdirExportErr := mkdirExportCmd.Run(); mkdirExportErr != nil {
		ifb.logger.Error(log.LogCategorySystem, mkdirExportErr.Error())
		return mkdirExportErr
	}

	filePathSvg := ifb.config.ApplicationConfig.ExportFolderPath + imageUuid + ".svg"
	errWriteFile := os.WriteFile(filePathSvg, []byte(output), 0644)

	if errWriteFile != nil {
		return errWriteFile
	}

	filePathPng := ifb.config.ApplicationConfig.ExportFolderPath + imageUuid + ".png"
	inkscapeCmd := exec.Command("inkscape", "--export-type=png", "--export-dpi=300", "--export-filename="+filePathPng, filePathSvg)
	if inkscapeErr := inkscapeCmd.Run(); inkscapeErr != nil {
		return inkscapeErr
	}

	rmSvgCmd := exec.Command("rm", "-f", filePathSvg)
	if rmSvgErr := rmSvgCmd.Run(); rmSvgErr != nil {
		ifb.logger.Error(log.LogCategorySystem, rmSvgErr.Error())
	}

	return nil
}
