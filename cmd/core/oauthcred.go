package core

type OAuthCred struct {

	// The bot account's OAuth password.
	Password string `json:"password,omitempty"`

	// The developer application client ID. Used for API calls to Twitch.
	ClientID string `json:"client_id,omitempty"`
}