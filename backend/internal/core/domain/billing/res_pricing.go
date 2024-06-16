package billing

type ResourcePricing struct {
	ID       string  `gorm:"primaryKey"`
	MemPrice float64 `gorm:"not null"`
	NetPrice float64 `gorm:"not null"`
	ReqPrice float64 `gorm:"not null"`
}
