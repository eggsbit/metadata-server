package entity

type EggsbitNftCollection struct {
	Identifier  string `bson:"identifier" json:"identifier"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Image       string `bson:"image" json:"image"`
	CoverImage  string `bson:"cover_image" json:"cover_image"`
	SocialLinks []string
}
