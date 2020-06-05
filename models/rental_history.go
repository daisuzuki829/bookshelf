package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RentalHistory struct {
	gorm.Model
	UserID         int        `gorm:"foreignkey:UserRefer"`
	BookID         int        `gorm:"foreignkey:RentalHistoryRefer"`
	LendDate       time.Time  `sql:"not null;type:date"`
	ReturnDeadline time.Time  `sql:"not null;type:date"`
	ReturnDate     time.Time  `sql:"type:date"`
	IgnoreMe       string     `gorm:"-"`
}

// RentalHistoryRepository is
type RentalHistoryRepository struct {
}

// NewRentalHistoryRepository ...
func NewRentalHistoryRepository() RentalHistoryRepository {
	return RentalHistoryRepository{}
}


//DB追加
func (r *RentalHistoryRepository) Add(rentalHistory *RentalHistory) {
	db, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic(err)
	}
	db.Create(rentalHistory)
	defer db.Close()
}

//DB更新
func (r *RentalHistoryRepository) Edit(rentalHistory RentalHistory) {
	db, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic(err)
	}
	db.Save(rentalHistory)
	db.Close()
}

//DB全取得
func (r *RentalHistoryRepository) GetAll() []RentalHistory {
	db, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic(err)
	}
	var rentalHistories []RentalHistory
	db.Find(&rentalHistories)

	//fmt.Printf("&rentalHistories : ")
	//spew.Dump(&rentalHistories)
	//fmt.Printf("\n")

	db.Close()
	return rentalHistories
}

//DB一つ取得
func (r *RentalHistoryRepository) GetOne(id int) RentalHistory {
	db, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic(err)
	}
	var rentalHistory RentalHistory
	db.First(&rentalHistory, id)
	db.Close()
	return rentalHistory
}

//DB削除
func (r *RentalHistoryRepository) Delete(id int) {
	db, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic(err)
	}
	var rentalHistory RentalHistory
	db.First(&rentalHistory, id)
	db.Delete(&rentalHistory)
	db.Close()
}

