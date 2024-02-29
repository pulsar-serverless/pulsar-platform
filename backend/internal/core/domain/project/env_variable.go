package project

type EnvVariable struct {
	ProjectID string `gorm:"primaryKey"`
	Key       string `gorm:"primaryKey"`
	Value     string ``
}
