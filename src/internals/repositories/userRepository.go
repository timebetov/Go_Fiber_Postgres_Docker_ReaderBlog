package repositories

import (
	"github.com/timebetov/readerblog/internals/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// Implementing UserRepository Interface

// First method is to create a new User in the database
func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// Getting all users from the database
func (r *userRepository) FindUsers(includeDeleted bool) ([]models.User, error) {
	var users []models.User
	var err error

	if includeDeleted {
		// Fetch all users, including the soft-deleted ones
		err = r.db.Unscoped().Where("deleted_at IS NOT NULL").Find(&users).Error
	} else {
		// Fetch all users, excluding the soft-deleted ones
		err = r.db.Find(&users).Error
	}

	return users, err
}

// Get one specific user by id from the database
func (r *userRepository) FindUserById(id string) (*models.User, error) {
	var user models.User
	// Find the user with the matching ID
	err := r.db.Unscoped().First(&user, "id = ?", id).Error
	return &user, err
}

// Get one specific user by username from the database
func (r *userRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	// Find the user with the matching username
	err := r.db.First(&user, "username = ?", username).Error
	return &user, err
}

// Update one specific user by id in the database
func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete one specific user by id in the database
func (r *userRepository) DeleteUser(force bool, user *models.User) error {
	if force {
		return r.db.Unscoped().Delete(user).Error
	} else {
		return r.db.Delete(user).Error
	}
}

func (r *userRepository) RestoreUser(user *models.User) error {
	user.DeletedAt = gorm.DeletedAt{}
	return r.db.Save(user).Error
}
