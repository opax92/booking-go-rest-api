package model

type Booking struct {
	Id        uint64 `gorm:"primary_key"`
	Event Event `gorm:"ForeignKey:EventId;AssociationForeignKey:Id"`
	EventId uint64
	BookedBy string
}
