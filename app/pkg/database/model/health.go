package database

type Health struct {
	ID      int    `gorm:"not null;primaryKey"`
	Message string `gorm:"size:128;not null"`
}
