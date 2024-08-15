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
// create Role
func CreateRole(context *gin.Context) {
	var Role model.Role
	context.BindJSON(&Role)
	err := model.CreateRole(&Role)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Roles created successfully",
		Data: model.Data{
			Status:  true,
			Message: "Roles created successfully",
			Result:  Role,
		},
	})
}

/*****************************************************************/
// get Roles
func GetRoles(context *gin.Context) {
	var Role []model.Role
	err := model.GetRoles(&Role)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Role received successfully",
		Data: model.Data{
			Status:  true,
			Message: "Role received successfully",
			Result:  Role,
		},
	})
}

/*****************************************************************/
// get Role by id
func GetRole(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var Role model.Role
	err := model.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Role created successfully",
		Data: model.Data{
			Status:  true,
			Message: "Role created successfully",
			Result:  Role,
		},
	})
}

/*****************************************************************/
// update Role by id
func UpdateRole(context *gin.Context) {
	var Role model.Role
	id, _ := strconv.Atoi(context.Param("id"))
	err := model.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("Roel"))
			return
		}

		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.BindJSON(&Role)
	err = model.UpdateRole(&Role)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Role updated successfully",
		Data: model.Data{
			Status:  true,
			Message: "Role updated successfully",
			Result:  Role,
		},
	})
}

/*****************************************************************/
// delete Role by id
func DeleteRole(context *gin.Context) {
	var Role model.Role
	id, _ := strconv.Atoi(context.Param("id"))
	err := model.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("Role"))
			return
		}
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.BindJSON(&Role)
	err = model.DeleteRole(&Role)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Role deleted successfully",
		Data: model.Data{
			Status:  true,
			Message: "Role deleted successfully",
		},
	})
}

/*****************************************************************/
