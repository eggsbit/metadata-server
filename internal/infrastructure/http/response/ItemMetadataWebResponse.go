package response

type ItemMetadataWebResponse struct {
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Image       *string                    `json:"image,omitempty"`
	Lottie      *string                    `json:"lottie,omitempty"`
	Attributes  []ItemAttributeWebResponse `json:"attributes,omitempty"`
}
