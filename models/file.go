package models

type File struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Type     string
	ParentID *uint
	UserID   uint
}
