package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tuananh31j/library-management-system/model"
	"github.com/tuananh31j/library-management-system/response"
	"github.com/tuananh31j/library-management-system/service"
	"github.com/tuananh31j/library-management-system/validation"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(u service.BookService) *BookController {
	return &BookController{
		BookService: u,
	}
}

// GET all books
func (ctr *BookController) GetAllBooks(c *fiber.Ctx) error {
	books, err := ctr.BookService.GetAllBooks(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to get all books",
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[[]model.Book]{
		Data:    books,
		Code:    fiber.StatusOK,
		Message: "Get all books successfully",
		Status:  "success",
	})
}

// GET book by ID
func (ctr *BookController) GetBookByID(c *fiber.Ctx) error {
	id, errParseUUID := uuid.Parse(c.Params("id"))
	if errParseUUID != nil {
		return fiber.NewError(fiber.StatusBadRequest, errParseUUID.Error())
	}
	book, err := ctr.BookService.GetBookByID(c, id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[model.Book]{
		Data:    book,
		Code:    fiber.StatusOK,
		Message: "Get book by ID successfully",
		Status:  "success",
	})
}

// POST create book
func (ctr *BookController) CreateNewBook(c *fiber.Ctx) error {
	body := new(validation.CreateBook)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusBadRequest,
			Message: "Failed to parse request body",
			Status:  "error",
		})
	}
	author_id, err := uuid.Parse(body.AuthorID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Author id is not valid!")
	}
	bookBody := &model.Book{
		AuthorID: author_id,
		Name:     body.Name,
		ID:       uuid.New(),
	}
	book, err := ctr.BookService.CreateNewBook(c, bookBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to create new book",
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(response.CustomResponse[model.Book]{
		Data:    book,
		Code:    fiber.StatusCreated,
		Message: "Create new book successfully",
		Status:  "success",
	})
}

// PUT update book
func (ctr *BookController) UpdateBook(c *fiber.Ctx) error {
	id, errParseUUID := uuid.Parse(c.Params("id"))
	if errParseUUID != nil {
		return fiber.NewError(fiber.StatusBadRequest, errParseUUID.Error())
	}
	var body model.Book
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusBadRequest,
			Message: "Failed to parse request body",
			Status:  "error",
		})
	}
	book, err := ctr.BookService.UpdateBook(c, id, body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[model.Book]{
		Data:    book,
		Code:    fiber.StatusOK,
		Message: "Update book by ID successfully",
		Status:  "success",
	})
}

// DELETE book
func (ctr *BookController) DeleteBook(c *fiber.Ctx) error {
	id, errParseUUID := uuid.Parse(c.Params("id"))
	if errParseUUID != nil {
		return fiber.NewError(fiber.StatusBadRequest, errParseUUID.Error())
	}
	err := ctr.BookService.DeleteBook(c, id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[interface{}]{
		Data:    nil,
		Code:    fiber.StatusOK,
		Message: "Delete book by ID successfully",
		Status:  "success",
	})
}
