package model

import (
	"bmacharia/jwt-go-rbac/database"
)

/*****************************************************************/

type Book struct {
	ID          uint     `gorm:"primaryKey" json:"id"`
	Title       string   `gorm:"size:255;not null" json:"title"`
	Description string   `gorm:"type:text" json:"description"`
	Price       float64  `gorm:"not null" json:"price"`
	AuthorID    uint     `gorm:"not null"  json:"authorid"`
	Author      Author   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author"`
	Orders      []Order  `gorm:"many2many:order_books;"`
	CategoryID  uint     `gorm:"not null" json:"categoryid"`
	Category    Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
}

/*****************************************************************/

func (book *Book) Save() (*Book, error) {
	err := database.Database.DB.Create(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return book, nil
}

/*****************************************************************/

func GetBooks(Book *[]Book) (err error) {
	err = database.Database.DB.Preload("Orders").Preload("Author.Books").Find(&Book).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func GetAllBooksByAuthor(id int) ([]Book, error) {
	var book []Book
	err := database.Database.DB.Where("author_id = ?", id).Find(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}
func GetBookById(book *Book, id int) error {
	err := database.Database.DB.Preload("Orders").Preload("Author.Books").Where("id=?", id).First(&book).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func GetBookQuantity(book *Book, quantity int) (err error) {
	err = database.Database.DB.Where("quantity >= ?", quantity).Find(&book).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func UpdateBook(Book *Book) (err error) {
	err = database.Database.DB.Updates(Book).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func DeleteBook(Book *Book) (err error) {
	err = database.Database.DB.Preload("Author").Preload("Category").Delete(Book).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/
