package controller

import (
	model "bmacharia/jwt-go-rbac/models"
	util "bmacharia/jwt-go-rbac/utils"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

/*****************************************************************/
// Register user
func Register(context *gin.Context) {
	var input model.Register
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	user := model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   3,
	}
	savedUser, err := user.Save()
	if err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

/*****************************************************************/
// User Login
func Login(context *gin.Context) {
	var input model.Login
	if err := context.ShouldBindJSON(&input); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}
		context.JSON(model.ResponseBadRequuest(errorMessage).Status, model.ResponseBadRequuest(errorMessage))
		return
	}
	user, err := model.GetUserByUsername(input.Username)
	if err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	err = user.ValidateUserPassword(input.Password)
	if err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	jwt, err := util.GenerateJWT(user)
	if err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	context.Header("Authorization", "Bearer "+jwt)
	context.SetCookie("jwt", jwt, context.GetInt(os.Getenv("TOKEN_TTL")), os.Getenv("COOCKIE_ROUTE"), os.Getenv(""), context.GetBool(os.Getenv("HTTPS")), context.GetBool(os.Getenv("HTTP_ONLY")))
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Successfully logged in",
		Data: model.Data{
			Status:  true,
			Message: "Successfully logged in",
			Result: gin.H{
				"token":    jwt,
				"username": input.Username,
			},
		},
	})
}

/*****************************************************************/
// Create user by admin and super admin
func CreateUser(context *gin.Context) {
	var input model.User
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	user := model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   input.RoleID,
	}
	savedUser, err := user.Save()
	if err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	context.JSON(http.StatusCreated, model.Response{
		Status:  http.StatusCreated,
		Message: "User created successfully",
		Data: model.Data{
			Status:  true,
			Message: "User created successfully",
			Result:  savedUser,
		},
	})
}

/*****************************************************************/
// add suport / customer / admin
func CreateUserByAdmin(context *gin.Context) {
	var input model.User
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	if input.RoleID == 1 {
		context.JSON(model.ResponseForbidden().Status, model.ResponseForbidden())
		return
	}
	user := model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   input.RoleID,
	}
	savedUser, err := user.Save()
	if err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	context.JSON(http.StatusCreated, model.Response{
		Status:  http.StatusCreated,
		Message: "User created successfully",
		Data: model.Data{
			Status:  true,
			Message: "User created successfully",
			Result:  savedUser,
		},
	})
}

/*****************************************************************/
// get all users
func GetUsers(context *gin.Context) {
	var user []model.User
	err := model.GetUsers(&user)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Successfully get all users",
		Data: model.Data{
			Status:  true,
			Message: "Successfully get all users",
			Result:  user,
		},
	})
}

/*****************************************************************/
// get user by id
func GetUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var user model.User
	err := model.GetUser(&user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("User"))
			return
		}
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, user)
}

/*****************************************************************/
// get user authentication  info
func Auth(context *gin.Context) {
	var user model.User
	user = util.CurrentUser(context)
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Successfully get user details",
		Data: model.Data{
			Status:  true,
			Message: "Successfully get user details",
			Result:  user,
		},
	})
}

/*****************************************************************/
// user update her informationn
func UpdateUserAuth(context *gin.Context) {
	var User model.User
	User = util.CurrentUser(context)
	var user model.User
	err := model.GetUser(&user, int(User.ID))
	context.ShouldBindBodyWithJSON(&User)
	err = model.UpdateUser(&user)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Successfully update user details",
		Data: model.Data{
			Status:  true,
			Message: "Successfully update user details",
			Result:  User,
		},
	})
}

/*****************************************************************/
// update user password
func UpdatePassword(context *gin.Context) {
	var User model.User
	User = util.CurrentUser(context)
	err := model.GetUser(&User, int(User.ID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("User"))
			return
		}
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	var input model.Password
	if err := context.ShouldBindBodyWithJSON(&input); err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	err = model.CheckPassword(&input)
	if err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	err = User.ValidateUserPassword(input.PrevPassword)
	if err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.NewPassWord), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	User.Password = string(passwordHash)
	err = model.UpdatePassword(&User)
	if err != nil {
		context.JSON(model.ResponseBadRequuest(err.Error()).Status, model.ResponseBadRequuest(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Password updated successfully",
		Data: model.Data{
			Status:  true,
			Message: "Password update successfully",
		},
	})
}

/*****************************************************************/
// update user
func UpdateUser(context *gin.Context) {
	var User model.User
	id, _ := strconv.Atoi(context.Param("id"))
	err := model.GetUser(&User, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("User"))
			return
		}
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.BindJSON(&User)
	err = model.UpdateUser(&User)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "User updated successfully",
		Data: model.Data{
			Status:  true,
			Message: "User updated successfully",
			Result:  User,
		},
	})
}

// update user
func UpdateUserByAdmin(context *gin.Context) {
	var User model.User
	id, _ := strconv.Atoi(context.Param("id"))
	err := model.GetUser(&User, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("User"))
			return
		}
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	if User.RoleID == 1 {
		context.JSON(model.ResponseForbidden().Status, model.ResponseForbidden())
		return
	}
	context.BindJSON(&User)
	err = model.UpdateUser(&User)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "User updated successfully",
		Data: model.Data{
			Status:  true,
			Message: "User updated successfully",
			Result:  User,
		},
	})
}

/*****************************************************************/
// dlete user
func DeleteUser(context *gin.Context) {
	var User model.User
	id, _ := strconv.Atoi(context.Param("id"))
	err := model.GetUser(&User, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("User"))
			return
		}
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	err = model.DeleteUser(&User)
	if err != nil {
		context.JSON(model.ResponseInternalServerError(err.Error()).Status, model.ResponseInternalServerError(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "User deleted successfully",
		Data: model.Data{
			Status:  true,
			Message: "User deleted successfully",
		},
	})
}

/*****************************************************************/
