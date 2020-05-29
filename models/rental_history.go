package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RentalHistory struct {
	gorm.Model
	UserId         int        `gorm:"foreignkey:UserRefer"`
	BookId         int        `gorm:"foreignkey:BookRefer"`
	LendDate       time.Time  `sql:"not null;type:date"`
	ReturnDeadline time.Time  `sql:"not null;type:date"`
	ReturnDate     time.Time  `sql:"type:date"`
	IgnoreMe       string     `gorm:"-"`
}

