package utilsTypes

import "github.com/develharsh/golang-crud-mongo/models"

type ResponseOfUserCRU struct {
	Success  bool         `json:"success"`
	Message  string       `json:"message"`
	UserData *models.User `json:"userData"`
}

type ResponseOfUserD struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
