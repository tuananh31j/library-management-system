package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/controller"
	"github.com/tuananh31j/library-management-system/service"
)

func BorrowerRoutes(apiGroup fiber.Router, u service.BorrowerService) {
	borrowerController := controller.NewBorrowerController(u)

	user := apiGroup.Group("/borrowers")

	user.Get("/all", borrowerController.GetAllBorrowers)
	user.Get("/:id", borrowerController.GetBorrowerByID)
	user.Post("/create", borrowerController.CreateNewBorrower)
	user.Put("/update/:id", borrowerController.UpdateBorrower)
	user.Delete("/delete/:id", borrowerController.DeleteBorrower)

}
