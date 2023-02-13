package models

//Status   int    `gorm:"type:tinyint unsigned;not null"`
type EdgeStreaming struct {
	StreamingID uint
	Streaming   Streaming
	EdgeID      uint
	Edge        Edge
	Status      int `gorm:"type:tinyint unsigned;default:0;not null"`
}
