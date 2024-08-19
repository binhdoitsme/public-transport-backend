package createtokens

type SessionResult struct {
	Ok           bool   `json:"ok"`
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
