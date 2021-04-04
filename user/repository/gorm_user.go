package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/solkn/soccer/api/entity"
	"github.com/solkn/soccer/api/user"
	"github.com/solkn/soccer/api/utils"
)

type UserGormRepo struct {
	conn *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) user.UsersRepository {
	return &UserGormRepo{conn: db}
}

func (u *UserGormRepo) Users() ([]entity.User, error) {
	var users []entity.User
	errs := u.conn.Preload("Role").Find(&users).Error
	if errs != nil {
		return users, utils.ErrInternalServerError
	}
	return users, errs
}

func (u *UserGormRepo) User(id uint32) (*entity.User, error) {
	user1 := entity.User{}
	errs := u.conn.Preload("Role").First(&user1, id).Error
	if errs != nil {
		return &user1, utils.ErrInternalServerError
	}
	return &user1, errs
}

func (u *UserGormRepo) StoreUser(user *entity.User) (*entity.User, error) {
	usr1 := user
	errs := u.conn.Preload("Role").Create(&usr1).GetErrors()
	if len(errs) > 0 {
		println("Store User Gorm Exception")
		return nil, utils.ErrInternalServerError
	}
	return usr1, nil
}

func (u *UserGormRepo) UpdateUser(muser *entity.User) (*entity.User, error) {
	usr1 := muser
	errs := u.conn.Preload("Role").Save(&usr1).Error
	if errs != nil {
		return usr1, utils.ErrInternalServerError
	}
	return usr1, errs
}

func (u *UserGormRepo) DeleteUser(id uint32) (*entity.User, error) {
	user1, errs := u.User(id)
	if errs != nil {
		return nil, utils.ErrInternalServerError
	}
	errs = u.conn.Delete(user1, id).Error
	if errs != nil {
		return nil, utils.ErrInternalServerError
	}
	return user1, errs
}
func (u *UserGormRepo) UserByUserName(user entity.User) (*entity.User, error) {
	user1 := entity.User{}
	errs := u.conn.Preload("Role").Where("email = ?", user.Email).First(&user1).GetErrors()
	if len(errs) > 0 {
		fmt.Println(errs)
		return &user1, utils.ErrInternalServerError
	}

	return &user1, nil
}
func (userRepo *UserGormRepo) PhoneExists(phone string) bool {
	user := entity.User{}
	errs := userRepo.conn.Find(&user, "phone=?", phone).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

func (userRepo *UserGormRepo) EmailExists(email string) bool {
	user := entity.User{}
	errs := userRepo.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

func (userRepo *UserGormRepo) UserRoles(user *entity.User) ([]entity.Role, []error) {
	userRoles := []entity.Role{}
	errs := userRepo.conn.Model(user).Related(&userRoles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}
