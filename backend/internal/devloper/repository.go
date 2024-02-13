package devloper

import (
	"gorm.io/gorm"
	"task-manager/internal/model/entity"
)

type repository struct {
	db *gorm.DB
}

func NewDeveloperRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAll() ([]entity.Developer, error) {
	var developers []entity.Developer
	err := r.db.Find(&developers).Error
	return developers, err
}

func (r *repository) FindAllAndTask() ([]entity.Developer, error) {
	var developers []entity.Developer
	err := r.db.Preload("Tasks").Find(&developers).Error
	return developers, err
}
