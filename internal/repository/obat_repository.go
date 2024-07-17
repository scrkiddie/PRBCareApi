package repository

import (
	"gorm.io/gorm"
	"prbcare_be/internal/entity"
)

type ObatRepository struct {
}

func NewObatRepository() *ObatRepository {
	return &ObatRepository{}
}
func (r *ObatRepository) FindAll(db *gorm.DB, obat *[]entity.Obat) error {
	return db.Preload("AdminApotek").Find(obat).Error
}

func (r *ObatRepository) FindAllByIdAdminApotek(db *gorm.DB, obat *[]entity.Obat, id_admin_apotek int) error {
	return db.Where("id_admin_apotek = ?", id_admin_apotek).Preload("AdminApotek").Find(obat).Error
}

func (r *ObatRepository) FindById(db *gorm.DB, obat *entity.Obat, id int) error {
	return db.Where("id = ?", id).Preload("AdminApotek").First(obat).Error
}

func (r *ObatRepository) FindByIdAndIdAdminApotek(db *gorm.DB, obat *entity.Obat, id int, id_admin_apotek int) error {
	return db.Where("id = ?", id).Preload("AdminApotek").Where("id_admin_apotek = ?", id_admin_apotek).First(obat).Error
}

func (r *ObatRepository) Update(db *gorm.DB, obat *entity.Obat) error {
	return db.Save(obat).Error
}
func (r *ObatRepository) Delete(db *gorm.DB, obat *entity.Obat) error {
	return db.Delete(obat).Error
}
func (r *ObatRepository) Create(db *gorm.DB, obat *entity.Obat) error {
	return db.Create(obat).Error
}
