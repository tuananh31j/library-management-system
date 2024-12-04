package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/helper"
	"github.com/tuananh31j/library-management-system/model"
	"github.com/tuananh31j/library-management-system/service"
	"gorm.io/gorm"
)

func InitRouter(a *fiber.App, db *gorm.DB) {
	apiGroup := a.Group("/api")

	AuthorRoutes(apiGroup, service.NewAuthorService(db))
	BookRoutes(apiGroup, service.NewBookService(db))
	BorrowerRoutes(apiGroup, service.NewBorrowerService(db))
	BorrowerBookRoutes(apiGroup, service.NewBorrowerBookService(db))
	apiGroup.Get("/import-data", func(c *fiber.Ctx) error {
		db.Exec("DELETE FROM borrower_books")
		db.Exec("DELETE FROM authors")
		db.Exec("DELETE FROM borrowers")
		db.Exec("DELETE FROM books")

		helper.InsertData(db, "./mock/author.csv", model.Author{})
		helper.InsertData(db, "./mock/book.csv", model.Book{})
		helper.InsertData(db, "./mock/borrower.csv", model.Borrower{})
		helper.InsertData(db, "./mock/borrower_book.csv", model.BorrowerBooks{})

		return c.SendStatus(fiber.StatusOK)
	})

}
