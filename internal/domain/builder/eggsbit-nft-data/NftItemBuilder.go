package eggsbitnftdata

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"

	"github.com/eggsbit/metadata-server/configs"
	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/domain/repository"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"
)

func NewNftItemBuilder(
	attributeRuleRepository repository.EggsbitNftItemAttributeRulesDocRepositoryInterface,
	logger log.LoggerInterface,
	config *configs.Config,
) NftItemBuilderInterface {
	return &NftItemBuilder{
		attributeRuleRepository: attributeRuleRepository,
		logger:                  logger,
		config:                  config,
	}
}

type NftItemBuilderInterface interface {
	BuildStartEggByIndex(index string, ctx context.Context) entity.EggsbitNftItem
}

type NftItemBuilder struct {
	attributeRuleRepository repository.EggsbitNftItemAttributeRulesDocRepositoryInterface
	logger                  log.LoggerInterface
	config                  *configs.Config
}

func (nib NftItemBuilder) BuildStartEggByIndex(index string, ctx context.Context) entity.EggsbitNftItem {
	var attributesArray []entity.EggsbitNftItemAttribute
	attributesArray = nib.getStartEggAttributes("root", attributesArray, ctx)
	nftImageUrl := nib.getNftImageUrl()

	return entity.EggsbitNftItem{
		Index:                index,
		CollectionIdentifier: "eggsbit_collection",
		Name:                 fmt.Sprintf("EggsBit #%s", index),
		Description:          "This special egg from the EggsBit collection will hatch into a unique pet in the future game. Adopt, care for, and watch your pet grow!",
		Image:                &nftImageUrl,
		Lottie:               nil,
		Attributes:           attributesArray,
	}
}

func (nib NftItemBuilder) getStartEggAttributes(identifier string, attributesArray []entity.EggsbitNftItemAttribute, ctx context.Context) []entity.EggsbitNftItemAttribute {
	keyNodes, _ := nib.attributeRuleRepository.GetRulesByParentIdentifier(identifier, "key", ctx)

	for _, keyNode := range keyNodes {
		valueNodes, _ := nib.attributeRuleRepository.GetRulesByParentIdentifier(keyNode.Identifier, "value", ctx)
		if len(valueNodes) == 0 {
			itemAttribute := entity.EggsbitNftItemAttribute{
				TraitType: *keyNode.Key,
				Value:     nil,
			}
			attributesArray = append(attributesArray, itemAttribute)
			continue
		} else {
			newAttributeItem, newAttributeItemIdentifier := nib.getRandomAttributeByProbability(*keyNode.Key, valueNodes)
			attributesArray = append(attributesArray, newAttributeItem)
			attributesArray = nib.getStartEggAttributes(newAttributeItemIdentifier, attributesArray, ctx)
		}
	}

	return attributesArray
}

func (nib NftItemBuilder) getRandomAttributeByProbability(keyNode string, valueNodes []*entity.EggsbitNftItemAttributeRule) (entity.EggsbitNftItemAttribute, string) {
	source := rand.NewSource(time.Now().UnixNano() + int64(len(valueNodes)))
	rng := rand.New(source)
	r := rng.Float64()
	var cumulativeProbability float64
	var randomResult *entity.EggsbitNftItemAttributeRule

	for _, valueNode := range valueNodes {
		var currentProbability float64
		if valueNode.Probability == nil {
			currentProbability = 0
		} else {
			currentProbability = *valueNode.Probability
		}

		cumulativeProbability += currentProbability
		if r < cumulativeProbability {
			randomResult = valueNode
			break
		}
	}

	return entity.EggsbitNftItemAttribute{
		TraitType: keyNode, Value: randomResult.Value,
	}, randomResult.Identifier
}

func (nib NftItemBuilder) getNftImageUrl() string {
	uuidId := uuid.NewString()
	baseUrl := nib.config.ApplicationConfig.NftItemImageBaseUrl
	return baseUrl + uuidId + ".png"
}
