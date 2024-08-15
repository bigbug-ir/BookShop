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

func ResponseInternalServerError() Response {
	return Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
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

func ResponseBadRequuest() Response {
	return Response{
		Status:  http.StatusBadRequest,
		Message: "Invalid request",
	}
}

/*****************************************************************/
