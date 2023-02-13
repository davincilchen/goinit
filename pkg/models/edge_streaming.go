package models

//Status   int    `gorm:"type:tinyint unsigned;not null"`
type EdgeStreaming struct {
	GormModel
	StreamingID uint `gorm:"not null"`
	Streaming   Streaming
	EdgeID      uint `gorm:"not null"`
	Edge        Edge
	Status      int `gorm:"type:tinyint unsigned;default:0;not null"`
}
