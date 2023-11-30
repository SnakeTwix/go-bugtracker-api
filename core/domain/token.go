package domain

type Token struct {
	Jwt          string `json:"jwt"`
	RefreshToken string `json:"refreshToken"`
}
