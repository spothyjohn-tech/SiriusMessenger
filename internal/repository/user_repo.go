package repository

import(
	"MesEdge/internal/models"
)

type UserRepository interface {
    Create(user *models.User) error
    GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
}