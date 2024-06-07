package entity

type EggsbitNftItem struct {
	Index                string                    `bson:"index"`
	CollectionIdentifier string                    `bson:"collection_identifier"`
	Name                 string                    `bson:"name"`
	Description          string                    `bson:"description"`
	Image                *string                   `bson:"image"`
	Lottie               *string                   `bson:"lottie"`
	Attributes           []EggsbitNftItemAttribute `bson:"attributes"`
}
