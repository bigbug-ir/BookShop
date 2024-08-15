package controller

import (
	model "bmacharia/jwt-go-rbac/models"
	util "bmacharia/jwt-go-rbac/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*****************************************************************/
//create the details of the applicant
func AddProfile(context *gin.Context) {
	tokenString, err := util.ExtractTokenFromHeader(context)
	if err != nil {
		context.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	id, err := util.ExtractUserIDFromToken(tokenString)
	if err != nil {
		context.JSON(model.ResponseBadRequuest().Status, model.ResponseBadRequuest())
		return
	}
	var User model.User
	userId := context.GetInt(id)
	err = model.GetUser(&User, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("User"))
			return
		}
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	var input model.Profile
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(model.ResponseBadRequuest().Status, model.ResponseBadRequuest())
		return
	}
	var profile = model.Profile{
		UserID:    User.ID,
		Image:     input.Image,
		Phone:     input.Phone,
		LastName:  input.LastName,
		FirstName: input.FirstName,
	}
	savedProfile, err := profile.Save()
	if err != nil {
		context.JSON(model.ResponseBadRequuest().Status, model.ResponseBadRequuest())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Profile saved successfully",
		Data: model.Data{
			Status:  true,
			Message: "Profile saved successfully",
			Result:  savedProfile,
		},
	})
}

/*****************************************************************/
//Get the details of the applicant
func GetProfile(context *gin.Context) {
	tokenString, err := util.ExtractTokenFromHeader(context)
	if err != nil {
		context.JSON(model.ResponseBadRequuest().Status, model.ResponseBadRequuest())
		return
	}
	id, err := util.ExtractUserIDFromToken(tokenString)
	if err != nil {
		context.JSON(model.ResponseBadRequuest().Status, model.ResponseBadRequuest())
		return
	}
	var User model.User
	userId := context.GetInt(id)
	err = model.GetUser(&User, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("User"))
			return
		}
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	var profile model.Profile
	err = model.GetProfile(&profile, userId)
	if err != nil {
		context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("User"))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "User profile retrieved successfully",
		Data: model.Data{
			Status:  true,
			Message: "User profile retrieved successfully",
			Result:  profile,
		},
	})
}

/*****************************************************************/
// update the details of the applicant
func UpdateProfile(context *gin.Context) {
	tokenString, err := util.ExtractTokenFromHeader(context)
	if err != nil {
		context.JSON(model.ResponseBadRequuest().Status, model.ResponseBadRequuest())
		return
	}
	id, err := util.ExtractUserIDFromToken(tokenString)
	if err != nil {
		context.JSON(model.ResponseBadRequuest().Status, model.ResponseBadRequuest())
		return
	}
	var User model.User
	userId := context.GetInt(id)
	err = model.GetUser(&User, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("User"))
			return
		}
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	var profile model.Profile
	err = model.GetProfile(&profile, userId)
	if err != nil {
		context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("Profile"))
		return
	}
	context.BindJSON(&profile)
	err = model.UpdateProfile(&profile)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "User profile updated successfully",
		Data: model.Data{
			Status:  true,
			Message: "User profile updated successfully",
			Result:  profile,
		},
	})
}

/*****************************************************************/
