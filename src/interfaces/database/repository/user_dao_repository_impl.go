package repository

import (
	"golang.org/x/crypto/bcrypt"

	"CleanArchitectureSample_golang/domain/model"
	"CleanArchitectureSample_golang/domain/repository"
	"CleanArchitectureSample_golang/interfaces/database/dao"
)

// UserDaoRepositoryImpl is UserDao interface.
type UserDaoRepositoryImpl struct {
	userDao dao.UserDao
}

// NewUserDaoRepository initializes UserDaoRepositoryImpl.
func NewUserDaoRepository(userDao dao.UserDao) repository.UserDaoRepository {
	return &UserDaoRepositoryImpl{userDao}
}

// Get returns user model.
func (r *UserDaoRepositoryImpl) Get(id uint) (user model.User, err error) {
	e, err := r.userDao.Find(id)
	if err != nil {
		return
	}
	user = model.User{
		ID:        e.ID,
		Name:      e.Name,
		Email:     e.Email,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
	return
}

// Create creates user.
func (r *UserDaoRepositoryImpl) Create(name string, email string, password string) (user model.User, err error) {
	encryptedPassword, err := encrypt(password)
	if err != nil {
		return
	}
	e, err := r.userDao.Create(name, email, encryptedPassword)
	if err != nil {
		return
	}
	user = model.User{
		ID:        e.ID,
		Name:      e.Name,
		Email:     e.Email,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
	return
}

// FindByEmailAndPassword finds user by email and password.
func (r *UserDaoRepositoryImpl) FindByEmailAndPassword(email string, password string) (user model.User, err error) {
	e, err := r.userDao.FindByEmail(email)
	if err != nil {
		return
	}
	if err = verify(password, e.EncryptedPassword); err != nil {
		return
	}
	user = model.User{
		ID:        e.ID,
		Name:      e.Name,
		Email:     e.Email,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
	return
}

func encrypt(password string) (encryptedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	return string(hash), nil
}

func verify(password string, encryptedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
}
