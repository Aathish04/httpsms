package requests

import (
	"time"

	"github.com/NdoleStudio/httpsms/pkg/entities"
	"github.com/NdoleStudio/httpsms/pkg/services"
)

// HeartbeatStore is the payload for fetching entities.Heartbeat of a phone number
type HeartbeatStore struct {
	request
	Charging bool   `json:"charging"`
	Owner    string `json:"owner"`
}

// Sanitize sets defaults to MessageOutstanding
func (input *HeartbeatStore) Sanitize() HeartbeatStore {
	input.Owner = input.sanitizeAddress(input.Owner)
	return *input
}

// ToStoreParams converts HeartbeatIndex to repositories.IndexParams
func (input *HeartbeatStore) ToStoreParams(user entities.AuthUser, source string, version string) services.HeartbeatStoreParams {
	return services.HeartbeatStoreParams{
		Owner:     input.Owner,
		Version:   version,
		Source:    source,
		Charging:  input.Charging,
		Timestamp: time.Now().UTC(),
		UserID:    user.ID,
	}
}
