package repository

import (
	"gorm.io/gorm"
	"prbcare_be/internal/entity"
)

type PasienRepository struct {
}

func NewPasienRepository() *PasienRepository {
	return &PasienRepository{}
}

func (r *PasienRepository) Search(db *gorm.DB, pasien *[]entity.Pasien, status string) error {
	query := db.Preload("AdminPuskesmas").Preload("Pengguna")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	return query.Find(pasien).Error
}

func (r *PasienRepository) SearchAsAdminPuskesmas(db *gorm.DB, pasien *[]entity.Pasien, id_admin_puskesmas int, status string) error {
	query := db.Where("id_admin_puskesmas = ?", id_admin_puskesmas).Preload("AdminPuskesmas").Preload("Pengguna")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	return query.Find(pasien).Error
}

func (r *PasienRepository) SearchAsPengguna(db *gorm.DB, pasien *[]entity.Pasien, id_pengguna int, status string) error {
	query := db.Where("id_pengguna = ?", id_pengguna).Preload("AdminPuskesmas").Preload("Pengguna")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	return query.Find(pasien).Error
}

func (r *PasienRepository) FindByIdAndStatus(db *gorm.DB, pasien *entity.Pasien, id int, status string) error {
	return db.Where("id = ?", id).Preload("AdminPuskesmas").Where("status = ?", status).Preload("Pengguna").First(pasien).Error
}

func (r *PasienRepository) FindByIdAndIdAdminPuskesmasAndStatus(db *gorm.DB, pasien *entity.Pasien, id int, id_admin_puskesmas int, status string) error {
	return db.Where("id = ?", id).Where("id_admin_puskesmas = ?", id_admin_puskesmas).Where("status = ?", status).Preload("AdminPuskesmas").Preload("Pengguna").First(pasien).Error
}

func (r *PasienRepository) Update(db *gorm.DB, pasien *entity.Pasien) error {
	return db.Save(pasien).Error
}
func (r *PasienRepository) Delete(db *gorm.DB, pasien *entity.Pasien) error {
	return db.Delete(pasien).Error
}
func (r *PasienRepository) Create(db *gorm.DB, pasien *entity.Pasien) error {
	return db.Create(pasien).Error
}
