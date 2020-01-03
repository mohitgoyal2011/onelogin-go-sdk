package models

// AuthBody is the request payload for authorization.
type AuthBody struct {
	GrantType string `json:"grant_type"`
}

// AuthResp is the authorization response payload.
type AuthResp struct {
	AccessToken  *string `json:"access_token,omitempty"`
	CreatedAt    *string `json:"created_at,omitempty"`
	ExpiresIn    *int32  `json:"expires_in,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	TokenType    *string `json:"token_type,omitempty"`
	AccountID    *int32  `json:"account_id,omitempty"`
}
