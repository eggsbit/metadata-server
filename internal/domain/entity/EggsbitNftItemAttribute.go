package entity

type EggsbitNftItemAttribute struct {
	TraitType string  `bson:"trait_type"`
	Value     *string `bson:"value"`
}
