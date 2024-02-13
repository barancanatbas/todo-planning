package entity

type Week struct {
	ID            uint `gorm:"primarykey"`
	NumberOfWeeks int
}
