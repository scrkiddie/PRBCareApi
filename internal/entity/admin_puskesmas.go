package entity

type AdminPuskesmas struct {
	ID            int    `gorm:"column:id;primaryKey;type:integer;autoIncrement;not null"`
	NamaPuskesmas string `gorm:"column:nama_puskesmas;type:varchar(100);not null"`
	Phone         string `gorm:"column:phone;type:varchar(15);unique;not null"`
	Address       string `gorm:"column:address;type:text;not null"`
	Username      string `gorm:"column:username;type:varchar(50);unique;not null"`
	Password      string `gorm:"column:password;type:varchar(255);not null"`
}

func (AdminPuskesmas) TableName() string {
	return "admin_puskesmas"
}
