package model

import "net/http"

/*****************************************************************/

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

/*****************************************************************/

type Data struct {
	Status  bool        `josn:"status"`
	Message string      `json:"status"`
	Result  interface{} `json:"result"`
}

/*****************************************************************/

func ResponseInternalServerError(msg string) Response {
	return Response{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

/*****************************************************************/

func ResponseErrRecordNotFound(name string) Response {
	return Response{
		Status:  http.StatusNotFound,
		Message: name + " not found.",
	}
}

/*****************************************************************/

func ResponseBadRequuest(msg string) Response {
	return Response{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}

/*****************************************************************/

func ResponseForbidden() Response {
	return Response{
		Status:  http.StatusForbidden,
		Message: "Forbidden",
	}
}

/*****************************************************************/
