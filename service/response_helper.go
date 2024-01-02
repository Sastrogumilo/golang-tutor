package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonBerhasil(value interface{}, res *gin.Context, message string) *gin.Context {

	type metadata struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	type response struct {
		Response interface{} `json:"response"`
		MetaData metadata    `json:"metadata"`
	}

	data := response{
		Response: value,
		MetaData: metadata{200, "OK"},
	}

	res.JSON(http.StatusOK, data)
	return res
}

func JsonGagal(message string, res *gin.Context) *gin.Context {

	type metadata struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	type response struct {
		Response interface{} `json:"response"`
		MetaData metadata    `json:"metadata"`
	}

	data := response{
		Response: []string{},
		MetaData: metadata{400, message},
	}

	res.JSON(http.StatusBadRequest, data)
	return res
}
