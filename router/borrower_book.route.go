package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/controller"
	"github.com/tuananh31j/library-management-system/service"
)

func BorrowerBookRoutes(apiGroup fiber.Router, u service.BorrowerBookService) {
	borrowerBookController := controller.NewBorrowerBookController(u)

	user := apiGroup.Group("/borrower_books")

	user.Get("/all", borrowerBookController.GetAllBorrowerBooks)
	user.Get("/:id", borrowerBookController.GetBorrowerBookByID)
	user.Post("/create", borrowerBookController.CreateNewBorrowerBook)
	user.Put("/update/:id", borrowerBookController.UpdateBorrowerBook)
	user.Delete("/delete/:id", borrowerBookController.DeleteBorrowerBook)

}
