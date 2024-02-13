package entity

type Developer struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Capacity int
	Tasks    []Task
}
