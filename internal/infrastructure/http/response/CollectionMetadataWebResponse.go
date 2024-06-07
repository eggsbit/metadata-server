package response

type CollectionMetadataWebResponse struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	CoverImage  string   `json:"cover_image"`
	SocialLinks []string `json:"social_links"`
}
