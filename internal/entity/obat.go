package entity

type Obat struct {
	ID            int32       `gorm:"column:id;primaryKey;type:integer;autoIncrement;not null"`
	NamaObat      string      `gorm:"column:nama_obat;type:varchar(100);not null"`
	Jumlah        int32       `gorm:"column:jumlah;type:integer;not null"`
	IdAdminApotek int32       `gorm:"column:id_admin_apotek;type:integer;not null"`
	AdminApotek   AdminApotek `gorm:"foreignKey:IdAdminApotek"`
}

func (Obat) TableName() string {
	return "obat"
}
