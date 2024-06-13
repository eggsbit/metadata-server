package eggsbitnftdata

import (
	"context"
	"fmt"
	"math/big"
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
	BuildStartEggByIndex(index *big.Int, ctx context.Context) (entity.NftItem, string)
}

type NftItemBuilder struct {
	attributeRuleRepository repository.EggsbitNftItemAttributeRulesDocRepositoryInterface
	logger                  log.LoggerInterface
	config                  *configs.Config
}

func (nib NftItemBuilder) BuildStartEggByIndex(index *big.Int, ctx context.Context) (entity.NftItem, string) {
	var attributesArray []entity.NftItemAttribute
	attributesArray = nib.getStartEggAttributes("root", attributesArray, ctx)
	nftImageUrl, imageUuid := nib.getNftImageUrl()

	return entity.NftItem{
		Index:                index.String(),
		CollectionIdentifier: "eggsbit_collection",
		Name:                 fmt.Sprintf("EggsBit #%s", index),
		Description:          "This special egg from the EggsBit collection will hatch into a unique pet in the future game. Adopt, care for, and watch your pet grow!",
		Image:                &nftImageUrl,
		Lottie:               nil,
		Attributes:           attributesArray,
	}, imageUuid
}

func (nib NftItemBuilder) getStartEggAttributes(identifier string, attributesArray []entity.NftItemAttribute, ctx context.Context) []entity.NftItemAttribute {
	keyNodes, _ := nib.attributeRuleRepository.GetRulesByParentIdentifier(identifier, "key", ctx)

	for _, keyNode := range keyNodes {
		valueNodes, _ := nib.attributeRuleRepository.GetRulesByParentIdentifier(keyNode.Identifier, "value", ctx)
		if len(valueNodes) == 0 {
			itemAttribute := entity.NftItemAttribute{
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

func (nib NftItemBuilder) getRandomAttributeByProbability(keyNode string, valueNodes []*entity.EggsbitNftItemAttributeRule) (entity.NftItemAttribute, string) {
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

	return entity.NftItemAttribute{
		TraitType: keyNode, Value: randomResult.Value,
	}, randomResult.Identifier
}

func (nib NftItemBuilder) getNftImageUrl() (string, string) {
	uuidId := uuid.NewString()
	baseUrl := nib.config.ApplicationConfig.NftItemImageBaseUrl
	return baseUrl + uuidId + ".png", uuidId
}
