package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/model"
	"github.com/tuananh31j/library-management-system/response"
	"github.com/tuananh31j/library-management-system/service"
)

type BorrowerController struct {
	BorrowerService service.BorrowerService
}

func NewBorrowerController(u service.BorrowerService) *BorrowerController {
	return &BorrowerController{
		BorrowerService: u,
	}
}

// GET all borrowers
func (ctr *BorrowerController) GetAllBorrowers(c *fiber.Ctx) error {
	borrowers, err := ctr.BorrowerService.GetAllBorrowers(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to get all borrowers",
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[[]model.Borrower]{
		Data:    borrowers,
		Code:    fiber.StatusOK,
		Message: "Get all borrowers successfully",
		Status:  "success",
	})
}

// GET borrower by ID
func (ctr *BorrowerController) GetBorrowerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	borrower, err := ctr.BorrowerService.GetBorrowerByID(c, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusBadRequest,
			Message: "Failed to get borrower by this ID" + id,
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[model.Borrower]{
		Data:    borrower,
		Code:    fiber.StatusOK,
		Message: "Get borrower by ID successfully",
		Status:  "success",
	})
}

// POST create borrower
func (ctr *BorrowerController) CreateNewBorrower(c *fiber.Ctx) error {
	body := new(model.Borrower)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusBadRequest,
			Message: "Failed to parse request body",
			Status:  "error",
		})
	}
	borrower, err := ctr.BorrowerService.CreateNewBorrower(c, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to create new borrower",
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(response.CustomResponse[model.Borrower]{
		Data:    borrower,
		Code:    fiber.StatusCreated,
		Message: "Create new borrower successfully",
		Status:  "success",
	})
}

// PUT update borrower
func (ctr *BorrowerController) UpdateBorrower(c *fiber.Ctx) error {
	id := c.Params("id")
	var body model.Borrower
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusBadRequest,
			Message: "Failed to parse request body",
			Status:  "error",
		})
	}
	borrower, err := ctr.BorrowerService.UpdateBorrower(c, id, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to update borrower",
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[model.Borrower]{
		Data:    borrower,
		Code:    fiber.StatusOK,
		Message: "Update borrower successfully",
		Status:  "success",
	})
}

// DELETE borrower
func (ctr *BorrowerController) DeleteBorrower(c *fiber.Ctx) error {
	id := c.Params("id")
	err := ctr.BorrowerService.DeleteBorrower(c, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CustomResponse[interface{}]{
			Data:    nil,
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to delete borrower by this ID" + id,
			Status:  "error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.CustomResponse[interface{}]{
		Data:    nil,
		Code:    fiber.StatusOK,
		Message: "Delete borrower by ID successfully",
		Status:  "success",
	})
}
