package entity

type MiniApp struct {
	Identifier       string  `bson:"identifier"`
	Url              string  `bson:"url"`
	Name             string  `bson:"name"`
	IconUrl          string  `bson:"icon_url"`
	TermsOfUseUrl    *string `bson:"terms_of_use_url"`
	PrivacyPolicyUrl *string `bson:"privacy_policy_url"`
}
