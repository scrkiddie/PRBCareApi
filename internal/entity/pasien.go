package entity

type Pasien struct {
	ID               int32          `gorm:"column:id;primaryKey;type:integer;autoIncrement;not null"`
	NoRekamMedis     string         `gorm:"column:no_rekam_medis;type:varchar(50);not null"`
	IdPengguna       int32          `gorm:"column:id_pengguna;type:integer;not null"`
	Pengguna         Pengguna       `gorm:"foreignKey:IdPengguna"`
	IdAdminPuskesmas int32          `gorm:"column:id_admin_puskesmas;type:integer;not null"`
	AdminPuskesmas   AdminPuskesmas `gorm:"foreignKey:IdAdminPuskesmas"`
	BeratBadan       int32          `gorm:"column:berat_badan;type:integer;not null"`
	TinggiBadan      int32          `gorm:"column:tinggi_badan;type:integer;not null"`
	TekananDarah     string         `gorm:"column:tekanan_darah;type:varchar(20);not null"`
	DenyutNadi       int32          `gorm:"column:denyut_nadi;type:integer;not null"`
	HasilLab         string         `gorm:"column:hasil_lab;type:text"`
	HasilEkg         string         `gorm:"column:hasil_ekg;type:text"`
	TanggalPeriksa   int64          `gorm:"column:tanggal_periksa;type:bigint;not null"`
	Status           string         `gorm:"column:status;type:status_pasien_enum;not null"`
}

func (Pasien) TableName() string {
	return "pasien"
}
