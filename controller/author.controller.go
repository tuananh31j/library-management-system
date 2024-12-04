package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tuananh31j/library-management-system/model"
	"github.com/tuananh31j/library-management-system/response"
	"github.com/tuananh31j/library-management-system/service"
	"github.com/tuananh31j/library-management-system/validation"
)

type AuthorController struct {
	AuthorService service.AuthorService
}

func NewAuthorController(u service.AuthorService) *AuthorController {
	return &AuthorController{
		AuthorService: u,
	}
}

// GET all authors
func (ctr *AuthorController) GetAllAuthors(c *fiber.Ctx) error {
	authors, err := ctr.AuthorService.GetAllAuthors(c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[[]model.Author]{
		Data:    authors,
		Code:    fiber.StatusOK,
		Message: "Get all authors successfully",
		Status:  "success",
	})
}

// GET author by ID
func (ctr *AuthorController) GetAuthorByID(c *fiber.Ctx) error {
	id, errParseUUID := uuid.Parse(c.Params("id"))
	if errParseUUID != nil {
		return fiber.NewError(fiber.StatusBadRequest, errParseUUID.Error())
	}
	author, err := ctr.AuthorService.GetAuthorByID(c, id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[model.Author]{
		Data:    author,
		Code:    fiber.StatusOK,
		Message: "Get author by ID successfully",
		Status:  "success",
	})

}

// POST create author
func (ctr *AuthorController) CreateNewAuthor(c *fiber.Ctx) error {
	body := new(validation.CreateAuthor)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	newAuthor := &model.Author{Name: body.Name}
	author, errCreate := ctr.AuthorService.CreateNewAuthor(c, newAuthor)
	if errCreate != nil {
		return fiber.NewError(fiber.StatusBadRequest, errCreate.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(response.CustomResponse[model.Author]{
		Data:    author,
		Code:    fiber.StatusCreated,
		Status:  "Success",
		Message: "Created author successfully!",
	})

}

// PUT update author
func (ctr *AuthorController) UpdateAuthor(c *fiber.Ctx) error {
	id, errParseUUID := uuid.Parse(c.Params("id"))
	if errParseUUID != nil {
		return fiber.NewError(fiber.StatusBadRequest, errParseUUID.Error())
	}
	body := new(model.Author)
	if errBody := c.BodyParser(body); errBody != nil {
		return fiber.NewError(fiber.StatusBadRequest, errBody.Error())
	}
	author, err := ctr.AuthorService.UpdateAuthor(c, id, *body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[model.Author]{
		Data:    author,
		Code:    fiber.StatusOK,
		Message: "Update author by ID successfully",
		Status:  "success",
	})
}

// DELETE author
func (ctr *AuthorController) DeleteAuthor(c *fiber.Ctx) error {
	id, errParseUUID := uuid.Parse(c.Params("id"))
	if errParseUUID != nil {
		return fiber.NewError(fiber.StatusBadRequest, errParseUUID.Error())
	}
	err := ctr.AuthorService.DeleteAuthor(c, id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[interface{}]{
		Data:    nil,
		Code:    fiber.StatusOK,
		Message: "Delete author by ID successfully",
		Status:  "success",
	})
}
