package entity

type EggsbitNftCollection struct {
	Identifier  string   `bson:"identifier"`
	Name        string   `bson:"name"`
	Description string   `bson:"description"`
	Image       string   `bson:"image"`
	CoverImage  string   `bson:"cover_image"`
	SocialLinks []string `bson:"social_links"`
}
