package repository

import (
	"gorm.io/gorm"
	"prb_care_api/internal/entity"
)

type AdminSuperRepository struct {
	Repository[entity.AdminSuper]
}

func NewAdminSuperRepository() *AdminSuperRepository {
	return &AdminSuperRepository{}
}

func (r *AdminSuperRepository) FindByUsername(db *gorm.DB, adminSuper *entity.AdminSuper, username string) error {
	return db.Where("username = ?", username).First(adminSuper).Error
}
func (r *AdminSuperRepository) FindById(db *gorm.DB, adminSuper *entity.AdminSuper, id int32) error {
	return db.Where("id = ?", id).First(adminSuper).Error
}
