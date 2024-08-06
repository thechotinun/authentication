package queries

import (
	"github.com/google/uuid"
	"github.com/thechotinun/authentication/app/models"
	"gorm.io/gorm"
)

// UserQueries struct for queries from User model.
type UserQueries struct {
	*gorm.DB
}

// GetUserByID query for getting one User by given ID.
func (q *UserQueries) GetUserByID(id uuid.UUID) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Send query to database.
	err := q.First(&user, "id = ?", id).Error
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// GetUserByEmail query for getting one User by given Email.
func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Send query to database.
	err := q.First(&user, "email = ?", email).Error
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// CreateUser query for creating a new user by given email and password hash.
func (q *UserQueries) CreateUser(u *models.User) error {
	// Send query to database.
	err := q.Create(u).Error
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
