package models

type PendingInstanceError struct {
	InstanceID string `json:"instance_id,omitempty" url:"instance_id,omitempty"`
	Message    string `json:"message,omitempty" url:"message,omitempty"`
}
