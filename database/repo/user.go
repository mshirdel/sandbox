package repo

import (
	"context"

	"github.com/mshirdel/sandbox/database/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return & UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, username string) error {
	u := models.User{
		Username: username,
	}
	
	return r.db.Model(&models.User{}).Table("users").Create(&u).Error
	
	// return r.db.WithContext(ctx).Create(&u).Error
}

func (r *UserRepo) Find(user *models.User, id uint) error {
	return r.db.Find(user, id).Error
}