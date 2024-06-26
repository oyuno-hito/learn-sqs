package database

type Event struct {
	ID   int    `gorm:"not null;primaryKey"`
	Text string `gorm:"size:128;not null"`
}
