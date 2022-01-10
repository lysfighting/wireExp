package clients

import "gorm.io/gorm"

func NewExpDB()*gorm.DB{
	return new(gorm.DB)
}
