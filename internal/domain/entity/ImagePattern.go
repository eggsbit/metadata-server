package entity

type ImagePattern struct {
	Identifier string `bson:"identifier"`
	Path       string `bson:"path"`
}
