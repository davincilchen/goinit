package models

type EdgeStreaming struct {
	Streaming
	ServerIp string `gorm:"type:char(32);not null"`
	Status   int    `gorm:"type:tinyint unsigned;not null"`
}
