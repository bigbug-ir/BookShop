package model

import (
	"bmacharia/jwt-go-rbac/database"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*****************************************************************/

type User struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	RoleID   uint     `gorm:"not null;DEFAULT:3" json:"role_id"`
	Username string   `gorm:"size:255;not null;unique" json:"username"`
	Email    string   `gorm:"size:255;not null;unique" json:"email"`
	Password string   `gorm:"size:255;not null" json:"-"`
	Role     Role     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"role,omitempty"`
	Profile  *Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"profile,omitempty"`
	Orders   []Order  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

/*****************************************************************/

func (user *User) Save() (*User, error) {
	err := database.Database.DB.Preload("Profile").Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

/*****************************************************************/
// Generate encrypted password
func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

/*****************************************************************/
// Get all users
func GetUsers(User *[]User) (err error) {
	err = database.Database.DB.Preload("Profile").Preload("Orders").Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/
// Get user by username
func GetUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.DB.Preload("Profile").Preload("Orders").Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

/*****************************************************************/
// Validate user password
func (user *User) ValidateUserPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

/*****************************************************************/
// Get user by id
func GetUserById(id int) (User, error) {
	var user User
	err := database.Database.DB.Where("id=?", id).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

/*****************************************************************/

func CheckCustomer(User *User) (err error) {
	err = database.Database.DB.Preload("Profile").Preload("Orders").Where("role_id=? ", int(3)).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/
// Get user by id
func GetUser(User *User, id int) (err error) {
	err = database.Database.DB.Preload("Profile").Preload("Orders").Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/
// Update user
func UpdateUser(User *User) (err error) {
	err = database.Database.DB.Omit("password").Preload("Profile").Updates(User).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func UpdatePassword(User *User) (err error) {
	err = database.Database.DB.Preload("Profile").Updates(User).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func DeleteUser(User *User) (err error) {
	err = database.Database.DB.Delete(User).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/
