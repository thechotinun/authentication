package queries

import (
	"github.com/thechotinun/authentication/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BookQueries struct for queries from Book model.
type BookQueries struct {
	*gorm.DB
}

// GetBooks method for getting all books.
func (q *BookQueries) GetBooks() ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Send query to database.
	err := q.Find(&books).Error
	if err != nil {
		// Return empty object and error.
		return books, err
	}

	// Return query result.
	return books, nil
}

// GetBooksByAuthor method for getting all books by given author.
func (q *BookQueries) GetBooksByAuthor(author string) ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Send query to database.
	err := q.Where("author = ?", author).Find(&books).Error
	if err != nil {
		// Return empty object and error.
		return books, err
	}

	// Return query result.
	return books, nil
}

// GetBook method for getting one book by given ID.
func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) {
	// Define book variable.
	book := models.Book{}

	// Send query to database.
	err := q.First(&book, "id = ?", id).Error
	if err != nil {
		// Return empty object and error.
		return book, err
	}

	// Return query result.
	return book, nil
}

// CreateBook method for creating book by given Book object.
func (q *BookQueries) CreateBook(b *models.Book) error {
	// Send query to database.
	err := q.Create(b).Error
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateBook method for updating book by given Book object.
func (q *BookQueries) UpdateBook(id uuid.UUID, b *models.Book) error {
	// Send query to database.
	err := q.Model(&models.Book{}).Where("id = ?", id).Updates(b).Error
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// DeleteBook method for delete book by given ID.
func (q *BookQueries) DeleteBook(id uuid.UUID) error {
	// Send query to database.
	err := q.Where("id = ?", id).Delete(&models.Book{}).Error
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
