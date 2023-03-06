package models

type Bbm struct {
	Id        int64 `gorm:"primaryKey" json:"id"`
	Premium   int   `gorm:"type:int" json:"premium"`
	Pertalite int   `gorm:"type:int" json:"pertalite"`
}