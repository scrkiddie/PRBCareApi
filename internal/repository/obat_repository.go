package repository

import (
	"gorm.io/gorm"
	"prb_care_api/internal/entity"
)

type ObatRepository struct {
	Repository[entity.Obat]
}

func NewObatRepository() *ObatRepository {
	return &ObatRepository{}
}

func (r *ObatRepository) FindAll(db *gorm.DB, obat *[]entity.Obat) error {
	return db.Preload("AdminApotek").Find(obat).Error
}
func (r *ObatRepository) FindAllByIdAdminApotek(db *gorm.DB, obat *[]entity.Obat, idAdminApotek int) error {
	return db.Where("id_admin_apotek = ?", idAdminApotek).Preload("AdminApotek").Find(obat).Error
}
func (r *ObatRepository) FindById(db *gorm.DB, obat *entity.Obat, id int) error {
	return db.Where("id = ?", id).First(obat).Error
}
func (r *ObatRepository) FindByIdAndIdAdminApotek(db *gorm.DB, obat *entity.Obat, id int, idAdminApotek int) error {
	return db.Where("id = ?", id).Where("id_admin_apotek = ?", idAdminApotek).First(obat).Error
}
func (r *ObatRepository) FindByIdAdminApotek(db *gorm.DB, obat *entity.Obat, idAdminApotek int) error {
	return db.Where("id_admin_apotek = ?", idAdminApotek).First(obat).Error
}
