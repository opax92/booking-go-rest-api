package model

type Event struct{
	Id        uint64 `gorm:"primary_key"`
	EventName string
}
