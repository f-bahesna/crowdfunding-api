package user

import (
	"gorm.io/gorm"
)

//interface repository  
//create Save func with `user,User` parameter
//return `User, error`
//R besar, interface for accessing with another package
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

//r kecil,ga bersifat public

type repository struct{
	db *gorm.DB
}
//balikan pointer dari repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//create function Save for repository
func (r *repository) Save(user User) (User, error){
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error){
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil{
		return user, err
	}

	return user, nil
}