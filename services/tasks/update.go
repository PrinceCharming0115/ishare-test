package tasks

import (
	"ishare-test/models"
)

func (service *Service) UpdateTask(task *models.Task) error {
	return service.DB.
		Save(task).
		Error
}
