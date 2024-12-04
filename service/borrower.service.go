package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/model"
	"gorm.io/gorm"
)

type BorrowerService interface {
	GetAllBorrowers(c *fiber.Ctx) ([]model.Borrower, error)
	GetBorrowerByID(c *fiber.Ctx, id string) (model.Borrower, error)
	CreateNewBorrower(c *fiber.Ctx, body *model.Borrower) (model.Borrower, error)
	UpdateBorrower(c *fiber.Ctx, id string, body model.Borrower) (model.Borrower, error)
	DeleteBorrower(c *fiber.Ctx, id string) error
}

type borrowerService struct {
	Db *gorm.DB
}

func NewBorrowerService(db *gorm.DB) BorrowerService {
	return &borrowerService{
		Db: db,
	}
}

// GET all borrowers
func (s *borrowerService) GetAllBorrowers(c *fiber.Ctx) ([]model.Borrower, error) {
	var borrowers []model.Borrower
	query := s.Db.WithContext(c.Context()).Where("is_active = ?", true).Order("created_at asc")
	result := query.Find(&borrowers)
	return borrowers, result.Error
}

// GET borrower by ID
func (s *borrowerService) GetBorrowerByID(c *fiber.Ctx, id string) (model.Borrower, error) {
	var borrower model.Borrower
	query := s.Db.WithContext(c.Context()).Where("id = ? AND is_active = ?", id, true)
	result := query.First(&borrower)
	return borrower, result.Error
}

// POST create borrower
func (s *borrowerService) CreateNewBorrower(c *fiber.Ctx, body *model.Borrower) (model.Borrower, error) {
	result := s.Db.WithContext(c.Context()).Create(body)
	return *body, result.Error
}

// PUT update borrower
func (s *borrowerService) UpdateBorrower(c *fiber.Ctx, id string, body model.Borrower) (model.Borrower, error) {
	var borrower model.Borrower
	query := s.Db.WithContext(c.Context()).Where("id = ?", id)
	result := query.Updates(body)
	return borrower, result.Error
}

// DELETE borrower
func (s *borrowerService) DeleteBorrower(c *fiber.Ctx, id string) error {
	query := s.Db.WithContext(c.Context()).Where("id = ?", id)
	result := query.Delete(&model.Borrower{})
	return result.Error
}
