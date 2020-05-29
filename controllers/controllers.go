package controllers

import (
	"github.com/jinzhu/gorm"
)

type Handler struct {
	Db *gorm.DB
}