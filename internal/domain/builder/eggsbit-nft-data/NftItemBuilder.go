package eggsbitnftdata

import (
	"fmt"

	"github.com/eggsbit/metadata-server/internal/domain/constant"
	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

func NewNftItemBuilder() NftItemBuilderInterface {
	return &NftItemBuilder{}
}

type NftItemBuilderInterface interface {
	BuildByIndexAndImage(index string, imagePath string) entity.EggsbitNftItem
	BuildByIndex(index string) entity.EggsbitNftItem
}

type NftItemBuilder struct {
}

func (nib NftItemBuilder) BuildByIndexAndImage(index string, imagePath string) entity.EggsbitNftItem {
	return entity.EggsbitNftItem{}
}

func (nib NftItemBuilder) BuildByIndex(index string) entity.EggsbitNftItem {
	statusValue := constant.StatusValueEgg
	patternValue := constant.PatternValueTwo
	colorSchemaValue := constant.ColorSchemaValueOne

	attributes := []entity.EggsbitNftItemAttribute{
		{TraitType: constant.KeyAttributeFather, Value: nil},
		{TraitType: constant.KeyAttributeMother, Value: nil},
		{TraitType: constant.KeyAttributeStatus, Value: &statusValue},
		{TraitType: constant.KeyAttributePattern, Value: &patternValue},
		{TraitType: constant.KeyAttributeColorSchema, Value: &colorSchemaValue},
	}

	return entity.EggsbitNftItem{
		Index:                index,
		CollectionIdentifier: "eggsbit_collection",
		Name:                 fmt.Sprintf("EggsBit #%s", index),
		Description:          "This special egg from the EggsBit collection will hatch into a unique pet in the future game. Adopt, care for, and watch your pet grow!",
		Image:                nil,
		Lottie:               nil,
		Attributes:           attributes,
	}
}
