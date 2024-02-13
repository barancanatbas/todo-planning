package task

import (
	"gorm.io/gorm"
	"task-manager/internal/model/entity"
)

const (
	CreateTaskBatchSize = 100
)

type repository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(tasks []entity.Task) error {
	return r.db.CreateInBatches(tasks, CreateTaskBatchSize).Error
}

func (r *repository) SaveNumberOfWeek(numberOfWeek int) error {
	return r.db.Save(&entity.Week{NumberOfWeeks: numberOfWeek}).Error
}

func (r *repository) GetNumberOfWeek() (entity.Week, error) {
	var week entity.Week
	err := r.db.First(&week).Error

	return week, err
}
