package database

type Message struct {
	ID   int    `gorm:"not null;primaryKey;autoIncrement" json:"id"`
	Text string `gorm:"size:128;not null" json:"text"`
}
