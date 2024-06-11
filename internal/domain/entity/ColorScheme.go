package entity

type ColorScheme struct {
	Identifier string            `bson:"identifier"`
	Type       string            `bson:"type"`
	Colors     map[string]string `bson:"colors"`
}
