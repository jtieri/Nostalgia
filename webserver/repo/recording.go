package repo

import (
	"github.com/jtieri/Nostalgia/webserver/app"
	"github.com/jtieri/Nostalgia/webserver/models"
	"gorm.io/gorm/clause"
)

func CreateRecording(recording *models.Recording) error {
	result := app.WebApp.DB.Create(&recording)

	if result != nil {
		return result.Error
	} else {
		return nil
	}
}

func GetAllRecordings() ([]*models.Recording, error) {
	var recordings []*models.Recording

	result := app.WebApp.DB.Preload(clause.Associations).Find(&recordings)
	if result.Error != nil {
		return nil, result.Error
	}

	return recordings, nil
}

func GetRecordingsByYear(year uint) ([]*models.Recording, error) {
	var recordings []*models.Recording

	result := app.WebApp.DB.Preload(clause.Associations).Where("year = ?", year).Find(&recordings)
	if result.Error != nil {
		return nil, result.Error
	}

	return recordings, nil
}

func GetRecordingsByNetwork(network string) ([]*models.Recording, error) {
	var recordings []*models.Recording

	result := app.WebApp.DB.Preload(clause.Associations).Where("networks LIKE ?", "%"+network+"%").Find(&recordings)
	if result.Error != nil {
		return nil, result.Error
	}

	return recordings, nil
}
