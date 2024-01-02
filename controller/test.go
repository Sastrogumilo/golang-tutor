package controllers

import (
	services "belajar_golang/service"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func TestIndex(res *gin.Context) {
	res.JSON(200, gin.H{"message": "Hello World"})
}

func TestIndex2(res *gin.Context) {

	var body map[string]interface{}

	data := map[string]interface{}{
		"message": "Hello World",
		"pesan":   "Babi",
		"body":    body,
	}

	services.JsonBerhasil(data, res, "")
}

func TestPost(res *gin.Context) {

	//cara get body
	var body map[string]interface{}

	res.ShouldBindJSON(&body)

	data := map[string]interface{}{
		"message": "Hello World",
		"pesan":   "Babi",
	}

	data3 := map[string]interface{}{
		"data3": "Hello World",
		"data4": "Babi1",
	}

	hasil := services.MergeMaps(data, body, data3)

	services.JsonBerhasil(hasil, res, "")
}

func TestPostForm(res *gin.Context) {

	file_data, err := res.MultipartForm() // mengambil data dari form

	if err != nil {
		services.JsonGagal("Gagal upload file", res)
		return
	}

	files := file_data.File["file"] // mengambil file dari form
	if len(files) == 0 {
		services.JsonGagal("File tidak ada", res)
		return
	}

	var arr_file []map[string]interface{}

	for i, file := range files {

		//deklarasi variable dulu

		var buffer bytes.Buffer

		if file.Size > 2<<20 { //2MB
			services.JsonGagal(fmt.Sprintf("File ke-%d terlalu besar", i+1), res)
			return
		}

		/**
		Baris kode yang di sorot untuk mengurai formulir multipart,
		yang biasanya digunakan untuk mengunggah file dalam permintaan HTTP.

		`2 << 20`: Ini adalah operasi pergeseran bit (bit shift operation)
		yang mengalikan 2 dengan 2^20, menghasilkan sekitar 2MB.
		bit shift operation adalah cara cepat untuk mengalikan atau membagi bilangan bulat
		dengan pangkat 2. Dalam kasus ini, `2 << 20` setara dengan `2 * 2^20`, yaitu sekitar 2MB.
		*/

		// Obtain file name and type
		fileName := file.Filename
		fileType := file.Header.Get("Content-Type")

		fileContent, err := file.Open()
		if err != nil {
			services.JsonBerhasil(fmt.Sprintf("Gagal membuka file pada file ke-%d", i+1), res, "")
			return
		}
		defer fileContent.Close()

		_, err = io.Copy(&buffer, fileContent) //copy file ke buffer

		if err != nil {
			services.JsonGagal(fmt.Sprintf("Gagal mengunggah file ke-%d", i+1), res)
			return
		}

		fileByte := buffer.Bytes() //mengambil data dari buffer

		hasilBase64 := base64.StdEncoding.EncodeToString(fileByte) //encode base64

		arr_file = append(arr_file, map[string]interface{}{
			"file_name": fileName,
			"file_type": fileType,
			"base64":    hasilBase64,
			"buffer":    buffer,
		})
	}

	services.JsonBerhasil(arr_file, res, "")

}

func TestPostForm2(res *gin.Context) {

	hasil_file, err := services.FormDataHelper(res, "file")

	if err != nil {
		services.JsonGagal(fmt.Sprintf("%v", err), res)
		return
	}

	services.JsonBerhasil(hasil_file, res, "")

}

func TestPostFormUploadMinio(res *gin.Context) {

	hasil_file, err := services.FormDataHelper(res, "file")

	if err != nil {
		services.JsonGagal(fmt.Sprintf("%v", err), res)
		return
	}

	buffer_file, ok := hasil_file[0]["buffer"].(bytes.Buffer)

	if !ok {
		services.JsonGagal("Gagal mengambil file", res)
		return
	}

	//check file size
	if buffer_file.Len() > 2<<20 { //2MB
		services.JsonGagal("File terlalu besar", res)
		return
	}

	minio_url, err := services.MinioService(buffer_file, hasil_file[0]["file_name"].(string), "tipsylion", hasil_file[0]["file_type"].(string))

	if err != nil {
		services.JsonGagal(fmt.Sprintf("%v", err), res)
		return
	}

	services.JsonBerhasil(minio_url, res, "")

}
