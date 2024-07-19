package repository

import (
	"gorm.io/gorm"
	"prbcare_be/internal/entity"
)

type KontrolBalikRepository struct {
}

func NewKontrolBalikRepository() *KontrolBalikRepository {
	return &KontrolBalikRepository{}
}

func (r *KontrolBalikRepository) Search(db *gorm.DB, kontrolBalik *[]entity.KontrolBalik, status string) error {
	query := db.Joins("JOIN pasien ON pasien.id = kontrol_balik.id_pasien").
		Joins("JOIN pengguna ON pengguna.id = pasien.id_pengguna").
		Joins("JOIN admin_puskesmas ON pasien.id_admin_puskesmas = admin_puskesmas.id")
	if status != "" {
		query = query.Where("kontrol_balik.status = ?", status)
	}
	return query.Preload("Pasien.AdminPuskesmas").Preload("Pasien.Pengguna").Find(&kontrolBalik).Error
}

func (r *KontrolBalikRepository) SearchAsAdminPuskesmas(db *gorm.DB, kontrolBalik *[]entity.KontrolBalik, id_admin_puskesmas int, status string) error {
	query := db.Joins("JOIN pasien ON pasien.id = kontrol_balik.id_pasien").
		Joins("JOIN pengguna ON pengguna.id = pasien.id_pengguna").
		Joins("JOIN admin_puskesmas ON pasien.id_admin_puskesmas = admin_puskesmas.id").
		Where("admin_puskesmas.id = ?", id_admin_puskesmas)
	if status != "" {
		query = query.Where("kontrol_balik.status = ?", status)
	}
	return query.Preload("Pasien.AdminPuskesmas").Preload("Pasien.Pengguna").Find(&kontrolBalik).Error
}

func (r *KontrolBalikRepository) SearchAsPengguna(db *gorm.DB, kontrolBalik *[]entity.KontrolBalik, id_pengguna int, status string) error {
	query := db.Joins("JOIN pasien ON pasien.id = kontrol_balik.id_pasien").
		Joins("JOIN pengguna ON pengguna.id = pasien.id_pengguna").
		Joins("JOIN admin_puskesmas ON pasien.id_admin_puskesmas = admin_puskesmas.id").
		Where("pengguna.id = ?", id_pengguna)
	if status != "" {
		query = query.Where("kontrol_balik.status = ?", status)
	}
	return query.Preload("Pasien.AdminPuskesmas").Preload("Pasien.Pengguna").Find(&kontrolBalik).Error
}

func (r *KontrolBalikRepository) FindByIdAndStatus(db *gorm.DB, kontrolBalik *entity.KontrolBalik, id int, status string) error {
	return db.Joins("JOIN pasien ON pasien.id = kontrol_balik.id_pasien").
		Joins("JOIN pengguna ON pengguna.id = pasien.id_pengguna").
		Joins("JOIN admin_puskesmas ON pasien.id_admin_puskesmas = admin_puskesmas.id").
		Where("kontrol_balik.id = ?", id).
		Where("kontrol_balik.status = ?", status).
		Preload("Pasien.AdminPuskesmas").
		Preload("Pasien.Pengguna").
		First(&kontrolBalik).Error
}

func (r *KontrolBalikRepository) FindByIdAndIdAdminPuskesmasAndStatus(db *gorm.DB, kontrolBalik *entity.KontrolBalik, id int, id_admin_puskesmas int, status string) error {
	return db.Joins("JOIN pasien ON pasien.id = kontrol_balik.id_pasien").
		Joins("JOIN pengguna ON pengguna.id = pasien.id_pengguna").
		Joins("JOIN admin_puskesmas ON pasien.id_admin_puskesmas = admin_puskesmas.id").
		Where("kontrol_balik.id = ?", id).
		Where("admin_puskesmas.id = ?", id_admin_puskesmas).
		Where("kontrol_balik.status = ?", status).
		Preload("Pasien.AdminPuskesmas").
		Preload("Pasien.Pengguna").
		First(&kontrolBalik).Error
}

func (r *KontrolBalikRepository) FindByIdAndIdAdminPuskesmasAndStatusOrStatus(db *gorm.DB, kontrolBalik *entity.KontrolBalik, id int, id_admin_puskesmas int, status1 string, status2 string) error {
	return db.Joins("JOIN pasien ON pasien.id = kontrol_balik.id_pasien").
		Joins("JOIN pengguna ON pengguna.id = pasien.id_pengguna").
		Joins("JOIN admin_puskesmas ON pasien.id_admin_puskesmas = admin_puskesmas.id").
		Where("kontrol_balik.id = ?", id).
		Where("admin_puskesmas.id = ?", id_admin_puskesmas).
		Where("kontrol_balik.status = ? OR kontrol_balik.status = ?", status1, status2).
		Preload("Pasien.AdminPuskesmas").
		Preload("Pasien.Pengguna").
		First(&kontrolBalik).Error
}

func (r *KontrolBalikRepository) FindByIdAndStatusOrStatus(db *gorm.DB, kontrolBalik *entity.KontrolBalik, id int, status1 string, status2 string) error {
	return db.Joins("JOIN pasien ON pasien.id = kontrol_balik.id_pasien").
		Joins("JOIN pengguna ON pengguna.id = pasien.id_pengguna").
		Joins("JOIN admin_puskesmas ON pasien.id_admin_puskesmas = admin_puskesmas.id").
		Where("kontrol_balik.id = ?", id).
		Where("kontrol_balik.status = ? OR kontrol_balik.status = ?", status1, status2).
		Preload("Pasien.AdminPuskesmas").
		Preload("Pasien.Pengguna").
		First(&kontrolBalik).Error
}

func (r *KontrolBalikRepository) Update(db *gorm.DB, kontrolBalik *entity.KontrolBalik) error {
	return db.Save(kontrolBalik).Error
}
func (r *KontrolBalikRepository) Delete(db *gorm.DB, kontrolBalik *entity.KontrolBalik) error {
	return db.Delete(kontrolBalik).Error
}
func (r *KontrolBalikRepository) Create(db *gorm.DB, kontrolBalik *entity.KontrolBalik) error {
	return db.Create(kontrolBalik).Error
}

// kebutuhan service pasien
func (r *KontrolBalikRepository) FindByIdPasienAndStatus(db *gorm.DB, kontrolBalik *entity.KontrolBalik, idPasien int, status string) error {
	return db.Joins("JOIN pasien ON pasien.id = kontrol_balik.id_pasien").
		Joins("JOIN pengguna ON pengguna.id = pasien.id_pengguna").
		Joins("JOIN admin_puskesmas ON pasien.id_admin_puskesmas = admin_puskesmas.id").
		Where("pasien.id = ?", idPasien).
		Where("kontrol_balik.status = ?", status).
		Preload("Pasien.AdminPuskesmas").
		Preload("Pasien.Pengguna").
		First(&kontrolBalik).Error
}
