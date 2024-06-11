package entity

type EggImagePattern struct {
	Identifier string `bson:"identifier"`
	Path       string `bson:"path"`
}
