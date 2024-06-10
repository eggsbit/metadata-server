package entity

type EggsbitNftItemAttributeRule struct {
	ParentIdentifier string   `bson:"parent_identifier"`
	Identifier       string   `bson:"identifier"`
	Key              *string  `bson:"key"`
	Value            *string  `bson:"value"`
	Probability      *float64 `bson:"probability"`
	Type             string   `bson:"type"`
}
