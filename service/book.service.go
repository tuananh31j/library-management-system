package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tuananh31j/library-management-system/model"
	"gorm.io/gorm"
)

type BookService interface {
	GetAllBooks(c *fiber.Ctx) ([]model.Book, error)
	GetBookByID(c *fiber.Ctx, id uuid.UUID) (model.Book, error)
	CreateNewBook(c *fiber.Ctx, body *model.Book) (model.Book, error)
	UpdateBook(c *fiber.Ctx, id uuid.UUID, body model.Book) (interface{}, error)
	DeleteBook(c *fiber.Ctx, id uuid.UUID) error
}

type bookService struct {
	Db *gorm.DB
}

func NewBookService(db *gorm.DB) BookService {
	return &bookService{
		Db: db,
	}
}

// GET all books
func (s *bookService) GetAllBooks(c *fiber.Ctx) ([]model.Book, error) {
	var books []model.Book
	query := s.Db.WithContext(c.Context()).Where("is_active = ?", true).Order("created_at asc")
	result := query.Find(&books)
	return books, result.Error
}

// GET book by ID
func (s *bookService) GetBookByID(c *fiber.Ctx, id uuid.UUID) (model.Book, error) {
	var book model.Book
	query := s.Db.WithContext(c.Context()).Where("id = ? AND is_active = ?", id, true)
	result := query.First(&book)
	return book, result.Error
}

// POST create book
func (s *bookService) CreateNewBook(c *fiber.Ctx, body *model.Book) (model.Book, error) {
	result := s.Db.WithContext(c.Context()).Create(body)
	return *body, result.Error
}

// PUT update book
func (s *bookService) UpdateBook(c *fiber.Ctx, id uuid.UUID, body model.Book) (interface{}, error) {
	query := s.Db.WithContext(c.Context()).Where("id = ? AND is_active = ?", id, true)
	result := query.Model(model.Book{}).Updates(body)
	return body, result.Error
}

// DELETE book
func (s *bookService) DeleteBook(c *fiber.Ctx, id uuid.UUID) error {
	query := s.Db.WithContext(c.Context()).Where("id = ?", id)
	result := query.Model(&model.Book{}).Update("is_active", false)
	return result.Error
}
