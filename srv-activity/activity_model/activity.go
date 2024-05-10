package activity_model

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	ActivityName string
	Description  string
	StartAt      string
	EndAt        string
	Products     []*Product `gorm:"foreignKey:ActivityID"`
}
