package repository

import (
	"gorm.io/gorm"
	"prbcare_be/internal/entity"
)

type AdminApotekRepository struct {
}

func NewAdminApotekRepository() *AdminApotekRepository {
	return &AdminApotekRepository{}
}
func (r *AdminApotekRepository) FindByUsername(db *gorm.DB, adminApotek *entity.AdminApotek, username string) error {
	return db.Where("username = ?", username).First(adminApotek).Error
}
func (r *AdminApotekRepository) FindById(db *gorm.DB, adminApotek *entity.AdminApotek, id int) error {
	return db.Where("id = ?", id).First(adminApotek).Error
}
func (r *AdminApotekRepository) Update(db *gorm.DB, adminApotek *entity.AdminApotek) error {
	return db.Save(adminApotek).Error
}
func (r *AdminApotekRepository) Delete(db *gorm.DB, adminApotek *entity.AdminApotek) error {
	return db.Delete(adminApotek).Error
}
func (r *AdminApotekRepository) Create(db *gorm.DB, adminApotek *entity.AdminApotek) error {
	return db.Create(adminApotek).Error
}
func (r *AdminApotekRepository) CountByUsername(db *gorm.DB, username any) (int64, error) {
	var count int64
	if err := db.Model(&entity.AdminApotek{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (r *AdminApotekRepository) CountByTelepon(db *gorm.DB, telepon any) (int64, error) {
	var count int64
	if err := db.Model(&entity.AdminApotek{}).Where("telepon = ?", telepon).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *AdminApotekRepository) FindAll(db *gorm.DB, adminApotek *[]entity.AdminApotek) error {
	return db.Find(adminApotek).Error
}
