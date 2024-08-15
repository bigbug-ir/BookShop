package controller

import (
	model "bmacharia/jwt-go-rbac/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*****************************************************************/
// create a author
func AddAuthor(context *gin.Context) {
	var input model.Author
	if err := context.ShouldBindBodyWithJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequuest())
		return
	}
	author := model.Author{
		Name:      input.Name,
		Biography: input.Biography,
	}

	savedAuthor, err := author.Save()
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Author create successfully",
		Data: model.Data{
			Status:  true,
			Message: "Author create successfully",
			Result:  savedAuthor,
		},
	})
}

/*****************************************************************/
// get all authors
func GetAuthors(context *gin.Context) {
	var author []model.Author
	err := model.GetAuthors(&author)
	if err != nil {
		context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("Author"))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Authors fetch data successfully",
		Data: model.Data{
			Status:  true,
			Message: "Authors fetch data successfully",
			Result:  author,
		},
	})
}

/*****************************************************************/
// get a author
func GetAuthor(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var author model.Author
	err := model.GetAuthor(&author, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("Author"))
			return
		}
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Author fetch data successfully",
		Data: model.Data{
			Status:  true,
			Message: "Author fetch data successfully",
			Result:  author,
		},
	})
}

/*****************************************************************/
// update a author
func UpdateAuthor(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var author model.Author
	err := model.GetAuthor(&author, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("Author"))
			return
		}
		context.JSON(http.StatusInternalServerError, model.ResponseInternalServerError())
		return
	}
	context.BindJSON(&author)
	err = model.UpdateAuthor(&author)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Author update successfully",
		Data: model.Data{
			Status:  true,
			Message: "Author update successfully",
			Result:  author,
		},
	})
}

/*****************************************************************/
// dlete a author
func DeleteAuthor(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var author model.Author
	err := model.GetAuthor(&author, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("Author"))
			return
		}
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	err = model.DeleteAuthor(&author)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Author deleted successfully.",
	})
}
