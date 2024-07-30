package entity

type AdminSuper struct {
	ID       int32  `gorm:"column:id;primaryKey;type:integer;autoIncrement;not null"`
	Username string `gorm:"column:username;type:varchar(50);unique;not null"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
}

func (AdminSuper) TableName() string {
	return "admin_super"
}
