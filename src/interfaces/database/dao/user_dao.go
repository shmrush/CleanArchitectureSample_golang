package dao

import (
	"CleanArchitectureSample_golang/infrastructure"

	"github.com/jinzhu/gorm"
)

// UserEntity is POGO of users table.
type UserEntity struct {
	gorm.Model
	Name              string
	Email             string
	EncryptedPassword string
}

// UserDao provides users table manipulation.
type UserDao struct {
	database infrastructure.Database
}

// NewUserDao initializes UserDao.
func NewUserDao(database infrastructure.Database) UserDao {
	return UserDao{database}
}

// TableName sets custom table name.
func (UserEntity) TableName() string {
	return "users"
}

// Find returns UserEntity.
func (d *UserDao) Find(id uint) (user UserEntity, err error) {
	user = UserEntity{}
	if result := d.database.Conn.Find(&user, id); result.Error != nil {
		err = result.Error
	}
	return
}

// Create creates new user.
func (d *UserDao) Create(name string, email string, encryptedPassword string) (user UserEntity, err error) {
	user = UserEntity{Name: name, Email: email, EncryptedPassword: encryptedPassword}
	if result := d.database.Conn.Create(&user); result.Error != nil {
		err = result.Error
	}
	return
}

// FindByEmail finds user by email
func (d *UserDao) FindByEmail(email string) (user UserEntity, err error) {
	user = UserEntity{}
	if result := d.database.Conn.Where("email = ?", email).First(&user); result.Error != nil {
		err = result.Error
	}
	return
}
