package controller

import (
	model "bmacharia/jwt-go-rbac/models"
	util "bmacharia/jwt-go-rbac/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*****************************************************************/
// Add order from costomer route
func AddOrder(context *gin.Context) {
	token, err := util.ExtractTokenFromHeader(context)
	if err != nil {
		context.JSON(model.ResponseBadRequuest().Status, model.ResponseBadRequuest())
		return
	}
	userId, err := util.ExtractUserIDFromToken(token)
	if err != nil {
		context.JSON(model.ResponseBadRequuest().Status, model.ResponseBadRequuest())
		return
	}
	id := context.GetInt(userId)
	var user model.User
	err = model.CheckCustomer(&user, id)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	var input struct {
		CustomerID uint `json:"customer_id"`
		Items      []struct {
			BookID   uint `json:"book_id"`
			Quantity uint `json:"quantity"`
		} `json:"items"`
	}
	err = context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(model.ResponseBadRequuest().Status, model.ResponseBadRequuest())
		return
	}
	var totalAmount float64
	var orderItems []model.OrderBook
	for _, item := range input.Items {
		var book model.Book
		err := model.GetBookById(&book, int(item.BookID))
		if err != nil {
			context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
			return
		}
		orderItem := model.OrderBook{
			BookID:   book.ID,
			Quantity: int(item.Quantity),
			Book:     book,
		}
		orderItems = append(orderItems, orderItem)
		totalAmount += book.Price * float64(item.Quantity)
	}
	order := model.Order{
		UserID:     user.ID,
		Item:       orderItems,
		TotalPrice: totalAmount,
	}
	savedOrder, err := order.Save()
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data: model.Data{
			Status:  true,
			Message: "Success",
			Result:  savedOrder,
		},
	})
}

/*****************************************************************/
// get orders to show admin route
func GetOrders(context *gin.Context) {
	var order []model.Order
	err := model.GetOrders(&order)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data: model.Data{
			Status:  true,
			Message: "Success",
			Result:  order,
		},
	})
}

/*****************************************************************/
// get all customer order for admin route
func GetAllOrderCustomer(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("user"))
	var order []model.Order
	err := model.GetOrdersCustomer(&order, id)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data: model.Data{
			Status:  true,
			Message: "Success",
			Result:  order,
		},
	})
}

/*****************************************************************/
// get customer order for admin route
func GetOrderCustomer(context *gin.Context) {
	user, _ := strconv.Atoi(context.Param("user"))
	id, _ := strconv.Atoi(context.Param("id"))
	var order model.Order
	err := model.GetOrderCustomer(&order, id, user)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data: model.Data{
			Status:  true,
			Message: "Success",
			Result:  order,
		},
	})
}

/*****************************************************************/
// get a customer order for customer route
func GetOrderCustomerAuth(context *gin.Context) {
	var User model.User
	User = util.CurrentUser(context)
	id, _ := strconv.Atoi(context.Param("id"))
	var order model.Order
	err := model.GetOrderCustomer(&order, id, int(User.ID))
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data: model.Data{
			Status:  true,
			Message: "Success",
			Result:  order,
		},
	})
}

/*****************************************************************/
// get all customer for customer route
func GetAllOrderCustomerAuth(context *gin.Context) {
	var User model.User
	User = util.CurrentUser(context)
	var order []model.Order
	err := model.GetOrdersCustomer(&order, int(User.ID))
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data: model.Data{
			Status:  true,
			Message: "Success",
			Result:  order,
		},
	})
}

/*****************************************************************/
// get order for admin route
func GetOrder(context *gin.Context) {
	var order model.Order
	id, _ := strconv.Atoi(context.Param("id"))
	err := model.GetOrder(&order, id)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data: model.Data{
			Status:  true,
			Message: "Success",
			Result:  order,
		},
	})
}

/*****************************************************************/
// update order by id for admin route
func UpdateOrder(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var order model.Order
	err := model.GetOrder(&order, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("Order"))
			return
		}
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.BindJSON(&order)
	err = model.UpdateOrder(&order)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Order spdate successfully",
		Data: model.Data{
			Status:  true,
			Message: "Order update successfully",
			Result:  order,
		},
	})
}

/*****************************************************************/
// delete order by id for amin roue
func DeleteOrder(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var order model.Order
	err := model.GetOrder(&order, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			context.JSON(http.StatusNotFound, model.ResponseErrRecordNotFound("Order"))
			return
		}
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	err = model.DeleteOrder(&order)
	if err != nil {
		context.JSON(model.ResponseInternalServerError().Status, model.ResponseInternalServerError())
		return
	}
	context.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Order deleted successfully",
	})
}

/*****************************************************************/
