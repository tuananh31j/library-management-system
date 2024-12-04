package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/controller"
	"github.com/tuananh31j/library-management-system/service"
)

func AuthorRoutes(apiGroup fiber.Router, u service.AuthorService) {
	authorController := controller.NewAuthorController(u)

	author := apiGroup.Group("/authors")

	author.Get("/all", authorController.GetAllAuthors)
	author.Get("/:id", authorController.GetAuthorByID)
	author.Post("/create", authorController.CreateNewAuthor)
	author.Put("/update/:id", authorController.UpdateAuthor)
	author.Delete("/delete/:id", authorController.DeleteAuthor)

}
