package types

type OAuthProvider struct {
	Provider string `json:"provider"`
	Key      string `json:"key"`
	Secret   string `json:"secret"`
	Callback string `json:"callback"`
}
