package model

import (
	"bmacharia/jwt-go-rbac/database"

	"gorm.io/gorm"
)

/*****************************************************************/

type Category struct {
	gorm.Model
	Name  string `gorm:"unique;not null" json:"name"`
	Books []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"books"`
}

/*****************************************************************/

func (category *Category) Save() (*Category, error) {
	err := database.Database.DB.Create(&category).Error
	if err != nil {
		return &Category{}, err
	}
	return category, nil
}

/*****************************************************************/

func GetCategories(Category *[]Category) (err error) {
	err = database.Database.DB.Find(&Category).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func GetCategory(Category *Category, id int) (err error) {
	err = database.Database.DB.Where("id=?", id).First(Category).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func GetCategoryByName(Category *Category, name string) (err error) {
	err = database.Database.DB.Where("name=?", name).First(&Category).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func UpdateCategory(Category *Category) (err error) {
	err = database.Database.DB.Updates(Category).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func DeleteCategory(Category *Category) (err error) {
	err = database.Database.DB.Delete(Category).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/
