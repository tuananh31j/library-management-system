package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/controller"
	"github.com/tuananh31j/library-management-system/service"
)

func BookRoutes(apiGroup fiber.Router, u service.BookService) {
	bookController := controller.NewBookController(u)

	user := apiGroup.Group("/books")

	user.Get("/all", bookController.GetAllBooks)
	user.Get("/:id", bookController.GetBookByID)
	user.Post("/create", bookController.CreateNewBook)
	user.Put("/update/:id", bookController.UpdateBook)
	user.Delete("/delete/:id", bookController.DeleteBook)

}
