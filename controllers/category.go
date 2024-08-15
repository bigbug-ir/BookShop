package controller

import (
	model "bmacharia/jwt-go-rbac/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCategory(context *gin.Context) {
	var input model.Category
	err := context.ShouldBindBodyWithJSON(&input)
	if err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	category := model.Category{
		Name: input.Name,
	}
	savedCategory, err := category.Save()
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Category saved successfully",
		Data: model.Data{
			Status:  true,
			Message: "Category saved successfully",
			Result:  savedCategory,
		},
	})
}

func GetCategories(context *gin.Context) {
	var Category []model.Category
	err := model.GetCategories(&Category)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Categories retrieved successfully",
		Data: model.Data{
			Status:  true,
			Message: "Categories retrieved successfully",
			Result:  Category,
		},
	})
}
func GetCategory(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var Category model.Category
	err := model.GetCategory(&Category, id)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Category retrieved successfully",
		Data: model.Data{
			Status:  true,
			Message: "Category retrieved successfully",
			Result:  Category,
		},
	})
}

func UpdateCategory(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var category model.Category
	err := model.GetCategory(&category, id)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.BindJSON(&category)
	err = model.UpdateCategory(&category)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Category updated successfully",
		Data: model.Data{
			Status:  true,
			Message: "Category updated successfully",
			Result:  category,
		},
	})
}

func DeleteCategory(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var category model.Category
	err := model.GetCategory(&category, id)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.BindJSON(&category)
	err = model.DeleteCategory(&category)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Category deleted successfully",
		Data: model.Data{
			Status:  true,
			Message: "Category deleted successfully",
		},
	})
}
