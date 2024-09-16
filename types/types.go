package types

import "time"

type RancherAPIResponse struct {
	ExpiresAt time.Time `json:"expiresAt"`
	Token     string    `json:"token"`
}
