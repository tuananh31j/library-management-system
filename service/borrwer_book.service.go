package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/model"
	"gorm.io/gorm"
)

type BorrowerBookService interface {
	GetAllBorrowerBooks(c *fiber.Ctx) ([]model.BorrowerBooks, error)
	GetBorrowerBookByID(c *fiber.Ctx, id string) (model.BorrowerBooks, error)
	CreateNewBorrowerBook(c *fiber.Ctx, body model.BorrowerBooks) (model.BorrowerBooks, error)
	UpdateBorrowerBook(c *fiber.Ctx, id string, body model.BorrowerBooks) (model.BorrowerBooks, error)
	DeleteBorrowerBook(c *fiber.Ctx, id string) error
}

type borrowerBookService struct {
	Db *gorm.DB
}

func NewBorrowerBookService(db *gorm.DB) BorrowerBookService {
	return &borrowerBookService{
		Db: db,
	}
}

// GET all borrower_books
func (s *borrowerBookService) GetAllBorrowerBooks(c *fiber.Ctx) ([]model.BorrowerBooks, error) {
	var borrowerBooks []model.BorrowerBooks
	query := s.Db.WithContext(c.Context()).Where("is_active = ?", true).Order("created_at asc")
	result := query.Find(&borrowerBooks)
	return borrowerBooks, result.Error
}

// GET borrower_book by ID
func (s *borrowerBookService) GetBorrowerBookByID(c *fiber.Ctx, id string) (model.BorrowerBooks, error) {
	var borrowerBook model.BorrowerBooks
	query := s.Db.WithContext(c.Context()).Where("id = ? AND is_active = ?", id, true)
	result := query.First(&borrowerBook)
	return borrowerBook, result.Error
}

// POST create borrower_book
func (s *borrowerBookService) CreateNewBorrowerBook(c *fiber.Ctx, body model.BorrowerBooks) (model.BorrowerBooks, error) {
	var borrowerBook model.BorrowerBooks
	result := s.Db.WithContext(c.Context()).Create(body)
	borrowerBook = body
	return borrowerBook, result.Error
}

// PUT update borrower_book
func (s *borrowerBookService) UpdateBorrowerBook(c *fiber.Ctx, id string, body model.BorrowerBooks) (model.BorrowerBooks, error) {
	var borrowerBook model.BorrowerBooks
	query := s.Db.WithContext(c.Context()).Where("id = ? AND is_active = ?", id, true)
	result := query.Updates(body)
	return borrowerBook, result.Error
}

// DELETE borrower_book
func (s *borrowerBookService) DeleteBorrowerBook(c *fiber.Ctx, id string) error {
	query := s.Db.WithContext(c.Context()).Where("id = ? AND is_active = ? ", id, true)
	result := query.Model(model.BorrowerBooks{}).Update("is_active", true)
	return result.Error
}
