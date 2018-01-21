package model

type Booking struct {
	Id        uint `gorm:"primary_key"`
	Event Event `gorm:"ForeignKey:EventId;AssociationForeignKey:Id"`
	EventId uint
	BookedBy string
}
