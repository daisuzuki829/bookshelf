package controllers

import (
	"github.com/daisuzuki829/bookshelf/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// GetAllRentalHistories ...
func (h *Handler) GetAllRentalHistories(c *gin.Context) {
	r := models.NewRentalHistoryRepository()
	rentalHistories := r.GetAll()



	c.HTML(http.StatusOK, "rental_histories.html", gin.H{
		"rental_histories": rentalHistories,
	})
}

// AddRentalHistories ...
func (h *Handler) AddRentalHistory(c *gin.Context) {
	r := models.NewRentalHistoryRepository()
	userId, _ := c.GetPostForm("user_id")
	userIdFmt, _ := strconv.Atoi(userId)
	bookId, _ := c.GetPostForm("book_id")
	bookIdFmt, _ := strconv.Atoi(bookId)
	lendDate, _ := c.GetPostForm("lend_date")
	lendDateFmt, _ := time.Parse("2006/01/02", lendDate)
	returnDeadline, _ := c.GetPostForm("return_deadline")
	returnDeadlineFmt, _ := time.Parse("2006/01/02", returnDeadline)
	r.Add(&models.RentalHistory{
		UserID:         userIdFmt,
		BookID:         bookIdFmt,
		LendDate:       lendDateFmt,
		ReturnDeadline: returnDeadlineFmt,
	})
	c.Redirect(http.StatusMovedPermanently, "/rental_histories")
}

// GetRentalHistories ...
func (h *Handler) GetRentalHistory(c *gin.Context) {
	r := models.NewRentalHistoryRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	rentalHistory := r.GetOne(id)
	c.HTML(http.StatusOK, "rentalHistory_edit.html", gin.H{
		"rentalHistory": rentalHistory,
	})
}

// EditRentalHistories ...
func (h *Handler) EditRentalHistory(c *gin.Context) {
	r := models.NewRentalHistoryRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	rentalHistory := r.GetOne(id)

	userId, _ := c.GetPostForm("user_id")
	rentalHistory.UserID, _ = strconv.Atoi(userId)
	bookId, _ := c.GetPostForm("book_id")
	rentalHistory.BookID, _ = strconv.Atoi(bookId)
	lendDate, _ := c.GetPostForm("lend_date")
	rentalHistory.LendDate, _ = time.Parse("2006/01/02", lendDate)
	returnDeadline, _ := c.GetPostForm("return_deadline")
	rentalHistory.ReturnDeadline, _ = time.Parse("2006/01/02", returnDeadline)

	r.Edit(rentalHistory)
	c.Redirect(http.StatusMovedPermanently, "/rental_histories")
}

// DeleteRentalHistories ...
func (h *Handler) DeleteRentalHistory(c *gin.Context) {
	r := models.NewRentalHistoryRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/rental_histories")
}
