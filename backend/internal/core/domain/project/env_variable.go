package project

type EnvVariable struct {
	ProjectID string `gorm:"primaryKey" json:"projectID"`
	Key       string `gorm:"primaryKey" json:"key"`
	Value     string `json:"value"`
}
