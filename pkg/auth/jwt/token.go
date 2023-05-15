package jwt

import "encoding/json"

type TokenInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   int64  `json:"expires_at"`
}

func (t *TokenInfo) GetAccessToken() string {
	return t.AccessToken
}

func (t *TokenInfo) GetExpiresAt() int64 {
	return t.ExpiresAt
}

func (t *TokenInfo) EncodeToJson() ([]byte, error) {
	return json.Marshal(t.AccessToken)
}

func (t *TokenInfo) GetTokenType() string {
	return t.TokenType
}
