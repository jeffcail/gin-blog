package models

import (
	core "github.com/jeffcail/gin-blog/db"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"column:name;default:0"`
}

func (m *Category) Category() string {
	return "category"
}

func init() {
	db := core.GetDB()

	db.AutoMigrate(&Category{})
}
