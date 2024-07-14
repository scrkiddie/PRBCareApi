package repository

import (
	"gorm.io/gorm"
	"prbcare_be/internal/entity"
)

type AdminSuperRepository struct {
}

func NewAdminSuperRepository() *AdminSuperRepository {
	return &AdminSuperRepository{}
}
func (r *AdminSuperRepository) FindByUsername(db *gorm.DB, adminSuper *entity.AdminSuper, username string) error {
	return db.Where("username = ?", username).Find(adminSuper).Error
}
func (r *AdminSuperRepository) FindById(db *gorm.DB, adminSuper *entity.AdminSuper, id int) error {
	return db.Where("id = ?", id).Find(adminSuper).Error
}
func (r *AdminSuperRepository) Update(db *gorm.DB, adminSuper *entity.AdminSuper) error {
	return db.Save(adminSuper).Error
}
