package entity

type KontrolBalik struct {
	ID             int32  `gorm:"column:id;primaryKey;type:integer;autoIncrement;not null"`
	IdPasien       int32  `gorm:"column:id_pasien;type:integer;not null"`
	Pasien         Pasien `gorm:"foreignKey:IdPasien"`
	TanggalKontrol int64  `gorm:"column:tanggal_kontrol;type:bigint;not null"`
	Status         string `gorm:"column:status;type:status_kontrol_balik_enum;not null"`
}

func (KontrolBalik) TableName() string {
	return "kontrol_balik"
}
