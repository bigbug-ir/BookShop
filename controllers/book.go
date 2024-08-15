package controller

import (
	"bmacharia/jwt-go-rbac/database"
	model "bmacharia/jwt-go-rbac/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*****************************************************************/
// add book
func AddBook(context *gin.Context) {
	var input model.Book
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequuest())
		return
	}
	category := model.Category{
		Name: input.Category.Name,
	}
	err = model.GetCategoryByName(&category, input.Category.Name)
	if err != nil {
		category.Save()
	}
	author := model.Author{
		Name:      input.Author.Name,
		Biography: input.Author.Biography,
	}
	err = model.GetAuthorByName(&author, author.Name)
	if err != nil {
		author.Save()
	}
	database.Database.DB.First(&author)
	book := model.Book{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		AuthorID:    author.ID,
		Author:      author,
		Category:    category,
	}
	authorBook, err := model.GetAllBooksByAuthor(int(author.ID))
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	author.Books = authorBook
	savedBook, err := book.Save()
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusCreated, model.Response{
		Status:  http.StatusCreated,
		Message: "Book saved successfully",
		Data: model.Data{
			Status:  true,
			Message: "Book saved successfully",
			Result:  savedBook,
		},
	})
}

/*****************************************************************/
// get all books
func GetBooks(context *gin.Context) {
	var book []model.Book
	err := model.GetBooks(&book)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Books retrieved successfully",
		Data: model.Data{
			Status:  true,
			Message: "Books retrieved successfully",
			Result:  book,
		},
	})
}

/*****************************************************************/
// get a book
func GetBook(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var book model.Book
	err := model.GetBookById(&book, id)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Book retrieved successfully",
		Data: model.Data{
			Status:  true,
			Message: "Book retrieved successfully",
			Result:  book,
		},
	})
}

/*****************************************************************/
//update a book
func UpdateBook(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var book model.Book
	err := model.GetBookById(&book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("Book"))
			return
		}
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	var input model.Book
	err = context.ShouldBindBodyWithJSON(&input)
	if err != nil {
		context.JSON(model.ResponseBadRequuest().Status, model.ResponseBadRequuest())
		return
	}
	book = model.Book{
		ID:          book.ID,
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		AuthorID:    input.AuthorID,
		CategoryID:  input.CategoryID,
	}
	err = model.UpdateBook(&book)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Book updated successfully",
		Data: model.Data{
			Status:  true,
			Message: "Book updated successfully",
			Result:  book,
		},
	})
}

/*****************************************************************/
//delete a book
func DeleteBook(context *gin.Context) {
	var Book model.Book
	id, _ := strconv.Atoi(context.Param("id"))
	err := model.GetBookById(&Book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("book"))
			return
		}
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	err = model.DeleteBook(&Book)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Book deleted successfully",
	})
}

/*****************************************************************/
