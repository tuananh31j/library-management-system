package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/model"
	"github.com/tuananh31j/library-management-system/response"
	"github.com/tuananh31j/library-management-system/service"
)

type BorrowerBookController struct {
	BorrowerBookService service.BorrowerBookService
}

func NewBorrowerBookController(u service.BorrowerBookService) *BorrowerBookController {
	return &BorrowerBookController{
		BorrowerBookService: u,
	}
}

// GET all borrower_books
func (ctr *BorrowerBookController) GetAllBorrowerBooks(c *fiber.Ctx) error {
	borrowerBooks, err := ctr.BorrowerBookService.GetAllBorrowerBooks(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to get all borrower_books",
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[interface{}]{
		Data:    borrowerBooks,
		Code:    fiber.StatusOK,
		Message: "Get all borrower_books successfully",
		Status:  "success",
	})
}

// GET borrower_book by ID
func (ctr *BorrowerBookController) GetBorrowerBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	borrowerBook, err := ctr.BorrowerBookService.GetBorrowerBookByID(c, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusBadRequest,
			Message: "Failed to get borrower_book by this ID" + id,
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[interface{}]{
		Data:    borrowerBook,
		Code:    fiber.StatusOK,
		Message: "Get borrower_book by ID successfully",
		Status:  "success",
	})
}

// POST create borrower_book
func (ctr *BorrowerBookController) CreateNewBorrowerBook(c *fiber.Ctx) error {
	var body model.BorrowerBooks
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusBadRequest,
			Message: "Failed to parse request body",
			Status:  "error",
		})
	}
	borrowerBook, err := ctr.BorrowerBookService.CreateNewBorrowerBook(c, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to create new borrower_book",
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(response.CustomResponse[interface{}]{
		Data:    borrowerBook,
		Code:    fiber.StatusCreated,
		Message: "Create new borrower_book successfully",
		Status:  "success",
	})
}

// PUT update borrower_book
func (ctr *BorrowerBookController) UpdateBorrowerBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var body model.BorrowerBooks
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusBadRequest,
			Message: "Failed to parse request body",
			Status:  "error",
		})
	}
	borrowerBook, err := ctr.BorrowerBookService.UpdateBorrowerBook(c, id, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to update borrower_book",
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[interface{}]{
		Data:    borrowerBook,
		Code:    fiber.StatusOK,
		Message: "Update borrower_book successfully",
		Status:  "success",
	})
}

// DELETE borrower_book
func (ctr *BorrowerBookController) DeleteBorrowerBook(c *fiber.Ctx) error {
	id := c.Params("id")
	err := ctr.BorrowerBookService.DeleteBorrowerBook(c, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to delete borrower_book",
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[interface{}]{
		Data:    nil,
		Code:    fiber.StatusOK,
		Message: "Delete borrower_book successfully",
		Status:  "success",
	})
}
