package entity

type Task struct {
	ID                uint `gorm:"primarykey"`
	Name              string
	Difficulty        int
	EstimatedDuration int
	Weight            int `gorm:"-"`
	DeveloperID       uint
}
