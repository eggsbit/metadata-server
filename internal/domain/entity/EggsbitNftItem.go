package entity

type EggsbitNftItem struct {
	Index                string `bson:"index" json:"index"`
	CollectionIdentifier string `bson:"collection_identifier" json:"collection_identifier"`
	Name                 string `bson:"name" json:"name"`
	Description          string `bson:"description" json:"description"`
	Image                string `bson:"image" json:"image"`
	Lottie               string `bson:"lottie" json:"lottie"`
	// Attributes -> {"trait_type": "Material", "value": "Wool fabric"}
}
