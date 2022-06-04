package responses

import "github.com/NdoleStudio/http-sms-manager/pkg/entities"

// MessageResponse is the payload containing an entities.Message
type MessageResponse struct {
	response
	Data entities.Message `json:"data"`
}
