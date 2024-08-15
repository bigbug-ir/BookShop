package controller

import (
	model "bmacharia/jwt-go-rbac/models"
	util "bmacharia/jwt-go-rbac/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*****************************************************************/
//create the details of the applicant
func AddProfile(context *gin.Context) {
	var User model.User
	User = util.CurrentUser(context)
	var input model.Profile
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
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
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
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
	var User model.User
	User = util.CurrentUser(context)
	var profile model.Profile
	err := model.GetProfile(&profile, int(User.ID))
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
	var User model.User
	User = util.CurrentUser(context)
	var profile model.Profile
	err := model.GetProfile(&profile, int(User.ID))
	if err != nil {
		context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("Profile"))
		return
	}
	context.BindJSON(&profile)
	err = model.UpdateProfile(&profile)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
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
