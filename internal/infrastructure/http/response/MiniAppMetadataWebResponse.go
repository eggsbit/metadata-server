package response

type MiniAppMetadataWebResponse struct {
	Url              string  `json:"url"`
	Name             string  `json:"name"`
	IconUrl          string  `json:"iconUrl"`
	TermsOfUseUrl    *string `json:"termsOfUseUrl,omitempty"`
	PrivacyPolicyUrl *string `json:"privacyPolicyUrl,omitempty"`
}
