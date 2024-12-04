package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tuananh31j/library-management-system/model"
	"gorm.io/gorm"
)

type AuthorService interface {
	GetAllAuthors(c *fiber.Ctx) ([]model.Author, error)
	GetAuthorByID(c *fiber.Ctx, id uuid.UUID) (model.Author, error)
	CreateNewAuthor(c *fiber.Ctx, body *model.Author) (model.Author, error)
	UpdateAuthor(c *fiber.Ctx, id uuid.UUID, body model.Author) (model.Author, error)
	DeleteAuthor(c *fiber.Ctx, id uuid.UUID) error
}

type authorService struct {
	Db  *gorm.DB
	Log *logrus.Logger
}

func NewAuthorService(db *gorm.DB) AuthorService {
	return &authorService{
		Db: db,
	}
}

// GET all authors
func (s *authorService) GetAllAuthors(c *fiber.Ctx) ([]model.Author, error) {
	var authors []model.Author
	query := s.Db.WithContext(c.Context()).Order("created_at asc").Where("is_active = ?", true)
	result := query.Find(&authors)
	return authors, result.Error
}

// GET author by ID
func (s *authorService) GetAuthorByID(c *fiber.Ctx, id uuid.UUID) (model.Author, error) {
	var author model.Author
	query := s.Db.WithContext(c.Context()).Where("id = ? AND is_active = ?", id, true)
	result := query.First(&author)
	return author, result.Error
}

// POST create author
func (s *authorService) CreateNewAuthor(c *fiber.Ctx, body *model.Author) (model.Author, error) {
	author := &model.Author{
		Name: body.Name,
		ID:   uuid.New(),
	}
	result := s.Db.WithContext(c.Context()).Create(author)
	return *author, result.Error
}

// PUT update author
func (s *authorService) UpdateAuthor(c *fiber.Ctx, id uuid.UUID, body model.Author) (model.Author, error) {
	query := s.Db.WithContext(c.Context()).Where("id = ? AND is_active = ?", id, true)
	result := query.Updates(body)
	return body, result.Error
}

// DELETE author
func (s *authorService) DeleteAuthor(c *fiber.Ctx, id uuid.UUID) error {
	query := s.Db.WithContext(c.Context()).Where("id = ?", id)
	result := query.Model(model.Author{}).Update("is_active", false)
	return result.Error
}
