package entity

type Pasien struct {
	ID               int            `gorm:"column:id;primaryKey;type:integer;autoIncrement;not null"`
	NoRekamMedis     string         `gorm:"column:no_rekam_medis;type:varchar(50);not null"`
	IDPengguna       int            `gorm:"column:id_pengguna;type:integer;not null"`
	Pengguna         Pengguna       `gorm:"foreignKey:IDPengguna"`
	IDAdminPuskesmas int            `gorm:"column:id_admin_puskesmas;type:integer;not null"`
	AdminPuskesmas   AdminPuskesmas `gorm:"foreignKey:IDAdminPuskesmas"`
	BeratBadan       float64        `gorm:"column:berat_badan;type:numeric(5,2);not null"`
	TinggiBadan      float64        `gorm:"column:tinggi_badan;type:numeric(5,2);not null"`
	TekananDarah     string         `gorm:"column:tekanan_darah;type:varchar(20);not null"`
	DenyutNadi       int            `gorm:"column:denyut_nadi;type:integer;not null"`
	HasilLab         string         `gorm:"column:hasil_lab;type:text"`
	HasilEKG         string         `gorm:"column:hasil_ekg;type:text"`
	TanggalPeriksa   int64          `gorm:"column:tanggal_periksa;type:bigint;not null"`
	Status           string         `gorm:"column:status;type:status_pasien_enum;not null"`
}

func (Pasien) TableName() string {
	return "pasien"
}
