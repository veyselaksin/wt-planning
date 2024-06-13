package models

type Abstract struct {
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
	IsActive  bool  `gorm:"default:true"`
}
