package refreshtokens

type SessionResult struct {
	Ok           bool   `json:"ok"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
