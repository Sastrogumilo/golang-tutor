package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
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

func JsonBerhasilFiber(value interface{}, res *fiber.Ctx, message string) *fiber.Ctx {

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

	res.JSON(data)
	return res
}

func JsonGagalFiber(message interface{}, res *fiber.Ctx) *fiber.Ctx {

	type metadata struct {
		Status  int         `json:"status"`
		Message interface{} `json:"message"`
	}

	type response struct {
		Response interface{} `json:"response"`
		MetaData metadata    `json:"metadata"`
	}

	data := response{
		Response: []string{},
		MetaData: metadata{400, message},
	}

	res.Status(400).JSON(data)
	return res
}
