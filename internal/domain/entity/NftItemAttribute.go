package entity

type NftItemAttribute struct {
	TraitType string  `bson:"trait_type"`
	Value     *string `bson:"value"`
}
