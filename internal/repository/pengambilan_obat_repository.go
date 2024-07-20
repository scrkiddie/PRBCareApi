package repository

import (
	"gorm.io/gorm"
	"prbcare_be/internal/entity"
)

type PengambilanObatRepository struct {
}

func NewPengambilanObatRepository() *PengambilanObatRepository {
	return &PengambilanObatRepository{}
}

func (r *PengambilanObatRepository) Search(db *gorm.DB, pengambilanObat *[]entity.PengambilanObat, status string) error {
	query := db
	if status != "" {
		query = query.Where("status = ?", status)
	}
	return query.Preload("Pasien.AdminPuskesmas").
		Preload("Pasien.Pengguna").
		Preload("Obat.AdminApotek").
		Find(&pengambilanObat).Error
}
func (r *PengambilanObatRepository) SearchAsAdminPuskesmas(db *gorm.DB, pengambilanObat *[]entity.PengambilanObat, idAdminPuskesmas int, status string) error {
	query := db.Joins("JOIN pasien ON pasien.id = pengambilan_obat.id_pasien").
		Joins("JOIN admin_puskesmas ON pasien.id_admin_puskesmas = admin_puskesmas.id").
		Where("admin_puskesmas.id = ?", idAdminPuskesmas)
	if status != "" {
		query = query.Where("pengambilan_obat.status = ?", status)
	}
	return query.Preload("Pasien.AdminPuskesmas").
		Preload("Pasien.Pengguna").
		Preload("Obat.AdminApotek").
		Find(&pengambilanObat).Error
}
func (r *PengambilanObatRepository) SearchAsAdminApotek(db *gorm.DB, pengambilanObat *[]entity.PengambilanObat, idAdminApotek int, status string) error {
	query := db.
		Joins("JOIN obat ON obat.id = pengambilan_obat.id_obat").
		Joins("JOIN admin_apotek ON admin_apotek.id = obat.id_admin_apotek").
		Where("admin_apotek.id = ?", idAdminApotek)
	if status != "" {
		query = query.Where("pengambilan_obat.status = ?", status)
	}
	return query.Preload("Pasien.AdminPuskesmas").
		Preload("Pasien.Pengguna").
		Preload("Obat.AdminApotek").
		Find(&pengambilanObat).Error
}
func (r *PengambilanObatRepository) SearchAsPengguna(db *gorm.DB, pengambilanObat *[]entity.PengambilanObat, idPengguna int, status string) error {
	query := db.Joins("JOIN pasien ON pasien.id = pengambilan_obat.id_pasien").
		Joins("JOIN pengguna ON pengguna.id = pasien.id_pengguna").
		Where("pengguna.id = ?", idPengguna)
	if status != "" {
		query = query.Where("pengambilan_obat.status = ?", status)
	}
	return query.Preload("Pasien.AdminPuskesmas").
		Preload("Pasien.Pengguna").
		Preload("Obat.AdminApotek").
		Find(&pengambilanObat).Error
}

func (r *PengambilanObatRepository) FindByIdAndStatus(db *gorm.DB, pengambilanObat *entity.PengambilanObat, id int, status string) error {
	return db.
		Where("id = ?", id).
		Where("status = ?", status).
		First(&pengambilanObat).Error
}
func (r *PengambilanObatRepository) FindByIdAndIdAdminPuskesmasAndStatus(db *gorm.DB, pengambilanObat *entity.PengambilanObat, id int, idAdminPuskesmas int, status string) error {
	return db.Joins("JOIN pasien ON pasien.id = pengambilan_obat.id_pasien").
		Joins("JOIN admin_puskesmas ON pasien.id_admin_puskesmas = admin_puskesmas.id").
		Where("pengambilan_obat.id = ?", id).
		Where("admin_puskesmas.id = ?", idAdminPuskesmas).
		Where("pengambilan_obat.status = ?", status).
		First(&pengambilanObat).Error
}
func (r *PengambilanObatRepository) FindByIdAndIdAdminApotekAndStatus(db *gorm.DB, pengambilanObat *entity.PengambilanObat, id int, idAdminApotek int, status string) error {
	return db.
		Joins("JOIN obat ON obat.id = pengambilan_obat.id_obat").
		Joins("JOIN admin_apotek ON admin_apotek.id = obat.id_admin_apotek").
		Where("pengambilan_obat.id = ?", id).
		Where("admin_apotek.id = ?", idAdminApotek).
		Where("pengambilan_obat.status = ?", status).
		First(&pengambilanObat).Error
}

func (r *PengambilanObatRepository) FindByIdAndStatusOrStatus(db *gorm.DB, pengambilanObat *entity.PengambilanObat, id int, status1 string, status2 string) error {
	return db.
		Where("id = ?", id).
		Where("status = ? OR status = ?", status1, status2).
		First(&pengambilanObat).Error
}
func (r *PengambilanObatRepository) FindByIdAndIdAdminPuskesmasAndStatusOrStatus(db *gorm.DB, pengambilanObat *entity.PengambilanObat, id int, idAdminPuskesmas int, status1 string, status2 string) error {
	return db.Joins("JOIN pasien ON pasien.id = pengambilan_obat.id_pasien").
		Joins("JOIN admin_puskesmas ON pasien.id_admin_puskesmas = admin_puskesmas.id").
		Where("pengambilan_obat.id = ?", id).
		Where("admin_puskesmas.id = ?", idAdminPuskesmas).
		Where("pengambilan_obat.status = ? OR pengambilan_obat.status = ?", status1, status2).
		First(&pengambilanObat).Error
}

func (r *PengambilanObatRepository) Update(db *gorm.DB, pengambilanObat *entity.PengambilanObat) error {
	return db.Save(pengambilanObat).Error
}
func (r *PengambilanObatRepository) Delete(db *gorm.DB, pengambilanObat *entity.PengambilanObat) error {
	return db.Delete(pengambilanObat).Error
}
func (r *PengambilanObatRepository) Create(db *gorm.DB, pengambilanObat *entity.PengambilanObat) error {
	return db.Create(pengambilanObat).Error
}

// kebutuhan service pasien
func (r *PengambilanObatRepository) FindByIdPasienAndStatus(db *gorm.DB, pengambilanObat *entity.PengambilanObat, idPasien int, status string) error {
	return db.
		Where("id_pasien = ?", idPasien).
		Where("status = ?", status).
		First(&pengambilanObat).Error
}
func (r *PengambilanObatRepository) FindByIdObat(db *gorm.DB, pengambilanObat *entity.PengambilanObat, idObat int) error {
	return db.Where("id_obat = ?", idObat).First(&pengambilanObat).Error
}
func (r *PengambilanObatRepository) FindByIdPasien(db *gorm.DB, pengambilanObat *entity.PengambilanObat, idPasien int) error {
	return db.Where("id_pasien = ?", idPasien).First(&pengambilanObat).Error
}
