package model

import "bmacharia/jwt-go-rbac/database"

/*****************************************************************/

type Author struct {
	ID        uint   `gorm:"primaryKey"  json:"id"`
	Name      string `gorm:"size:255;not null;unique" json:"name"`
	Biography string `gorm:"type:text" json:"biography"`
	Books     []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; "`
}

/*****************************************************************/

func (author *Author) Save() (*Author, error) {
	err := database.Database.DB.Create(&author).Error
	if err != nil {
		return &Author{}, err
	}
	return author, nil
}

/*****************************************************************/

func GetAuthors(Author *[]Author) (err error) {
	err = database.Database.DB.Find(&Author).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func GetAuthor(Author *Author, id int) error {
	err := database.Database.DB.Preload("Books").Where("id = ?", id).First(&Author).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func GetAuthorByName(Author *Author, name string) error {
	err := database.Database.DB.Preload("Books").Where(" name= ?", name).First(&Author).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func UpdateAuthor(Author *Author) (err error) {
	err = database.Database.DB.Updates(Author).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func DeleteAuthor(Author *Author) (err error) {
	err = database.Database.DB.Delete(&Author).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/
