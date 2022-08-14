package msgDao

import (
	"four-seasons-of-Kunming/app/models"
	"four-seasons-of-Kunming/config/database"
)

func GetMsg() ([]models.Msg, error) {
	var messages []models.Msg

	result := database.DB.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}

func CreateMsg(msg models.Msg) error {
	result := database.DB.Create(&msg)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
