package controllers

import (
	services "belajar_golang/service"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
)

func TestIndexFiber(res *fiber.Ctx) {
	res.JSON(fiber.Map{"message": "Hello World"})
}

func TestIndex2F(res *fiber.Ctx) {

	var body map[string]interface{}

	data := map[string]interface{}{
		"message": "Hello World",
		"pesan":   "Babi",
		"body":    body,
	}

	services.JsonBerhasilFiber(data, res, "")
}

func TestIndex2FGagal(res *fiber.Ctx) {

	var body map[string]interface{}

	data := map[string]interface{}{
		"message": "Hello World",
		"pesan":   "Babi",
		"body":    body,
	}

	services.JsonGagalFiber(data, res)
}

func TestPostFiber(res *fiber.Ctx) {

	//cara get body
	type Body struct {
		Name string `json:"name" xml:"name" form:"name"`
		Pass string `json:"pass" xml:"pass" form:"pass"`
	}

	var body = new(Body)

	if err := res.BodyParser(body); err != nil {
		services.JsonGagalFiber(fmt.Sprintf("%v", err), res)
		return
	}

	data := map[string]interface{}{
		"message": "Hello World",
		"pesan":   "Babi",
	}

	data3 := map[string]interface{}{
		"data3": "Hello World",
		"data4": "Babi1",
	}

	bodyMap := services.StructToMap(body)

	hasil := services.MergeMaps(data, bodyMap, data3)

	services.JsonBerhasilFiber(hasil, res, "")
}

func TestPostFormFiber(res *fiber.Ctx) {

	file_data, err := res.MultipartForm() // mengambil data dari form

	if err != nil {
		log.Println(err)
		services.JsonGagalFiber("Gagal upload file", res)
		return
	}

	files := file_data.File["file"] // mengambil file dari form
	if len(files) == 0 {
		services.JsonGagalFiber("File tidak ada", res)
		return
	}

	var arr_file []map[string]interface{}

	for i, file := range files {

		//deklarasi variable dulu

		var buffer bytes.Buffer

		if file.Size > 2<<20 { //2MB
			services.JsonGagalFiber(fmt.Sprintf("File ke-%d terlalu besar", i+1), res)
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
			services.JsonBerhasilFiber(fmt.Sprintf("Gagal membuka file pada file ke-%d", i+1), res, "")
			return
		}
		defer fileContent.Close()

		_, err = io.Copy(&buffer, fileContent) //copy file ke buffer

		if err != nil {
			services.JsonGagalFiber(fmt.Sprintf("Gagal mengunggah file ke-%d", i+1), res)
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

	services.JsonBerhasilFiber(arr_file, res, "")

}

func TestPostForm2Fiber(res *fiber.Ctx) {

	hasil_file, err := services.FormDataHelperFiber(res, "file")

	if err != nil {
		services.JsonGagalFiber(fmt.Sprintf("%v", err), res)
		return
	}

	services.JsonBerhasilFiber(hasil_file, res, "")

}

func TestPostFormUploadMinioFiber(res *fiber.Ctx) {

	hasil_file, err := services.FormDataHelperFiber(res, "file")

	if err != nil {
		services.JsonGagalFiber(fmt.Sprintf("%v", err), res)
		return
	}

	buffer_file, ok := hasil_file[0]["buffer"].(bytes.Buffer)

	if !ok {
		services.JsonGagalFiber("Gagal mengambil file", res)
		return
	}

	//check file size
	if buffer_file.Len() > 2<<20 { //2MB
		services.JsonGagalFiber("File terlalu besar", res)
		return
	}

	minio_url, err := services.MinioService(buffer_file, hasil_file[0]["file_name"].(string), "tipsylion", hasil_file[0]["file_type"].(string))

	if err != nil {
		services.JsonGagalFiber(fmt.Sprintf("%v", err), res)
		return
	}

	services.JsonBerhasilFiber(minio_url, res, "")

}

func TestDB1(res *fiber.Ctx) {

	koneksi := services.GetLocal[services.TestMainDBService](res, "mainDB")

	query := "SELECT * FROM sinarmas_dpmall.user_member"

	hasil, err := koneksi.TestQuery(query)

	if err != nil {
		services.JsonGagalFiber(fmt.Sprintf("%v", err), res)
		return
	}

	type UserMember struct {
		Agama  string `json:"agama"`
		Alamat string `json:"alamat"`
		Uid    string `json:"uid"`
	}

	// convert ke struct
	var user_member []UserMember

	err = services.ConvertToStruct(hasil, &user_member)
	if err != nil {
		services.JsonGagalFiber(fmt.Sprintf("%v", err), res)
		return
	}

	//make map to concat 2 map
	hasil_map := map[string]interface{}{
		"hasil1": user_member,
		"hasil2": hasil,
	}

	/** Normal using map interface
	C:\Users\Belldandy>autocannon http://192.168.20.100:6969/test_db
	Running 10s test @ http://192.168.20.100:6969/test_db
	10 connections


	┌─────────┬──────┬───────┬───────┬───────┬──────────┬─────────┬────────┐
	│ Stat    │ 2.5% │ 50%   │ 97.5% │ 99%   │ Avg      │ Stdev   │ Max    │
	├─────────┼──────┼───────┼───────┼───────┼──────────┼─────────┼────────┤
	│ Latency │ 6 ms │ 11 ms │ 35 ms │ 43 ms │ 13.56 ms │ 8.63 ms │ 123 ms │
	└─────────┴──────┴───────┴───────┴───────┴──────────┴─────────┴────────┘
	┌───────────┬─────────┬─────────┬─────────┬─────────┬─────────┬─────────┬─────────┐
	│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%   │ Avg     │ Stdev   │ Min     │
	├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┤
	│ Req/Sec   │ 512     │ 512     │ 729     │ 790     │ 712.3   │ 78.5    │ 512     │
	├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┤
	│ Bytes/Sec │ 24.2 MB │ 24.2 MB │ 34.5 MB │ 37.4 MB │ 33.7 MB │ 3.71 MB │ 24.2 MB │
	└───────────┴─────────┴─────────┴─────────┴─────────┴─────────┴─────────┴─────────┘

	Req/Bytes counts sampled once per second.
	# of samples: 10

	7k requests in 10.04s, 337 MB read
	*/

	/** Convert to struct
	C:\Users\Belldandy>autocannon http://192.168.20.100:6969/test_db
	Running 10s test @ http://192.168.20.100:6969/test_db
	10 connections

	┌─────────┬──────┬───────┬───────┬───────┬──────────┬─────────┬───────┐
	│ Stat    │ 2.5% │ 50%   │ 97.5% │ 99%   │ Avg      │ Stdev   │ Max   │
	├─────────┼──────┼───────┼───────┼───────┼──────────┼─────────┼───────┤
	│ Latency │ 7 ms │ 12 ms │ 36 ms │ 42 ms │ 14.45 ms │ 7.44 ms │ 66 ms │
	└─────────┴──────┴───────┴───────┴───────┴──────────┴─────────┴───────┘
	┌───────────┬─────────┬─────────┬─────────┬─────────┬─────────┬────────┬─────────┐
	│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%   │ Avg     │ Stdev  │ Min     │
	├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼────────┼─────────┤
	│ Req/Sec   │ 581     │ 581     │ 659     │ 745     │ 668.8   │ 53.74  │ 581     │
	├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼────────┼─────────┤
	│ Bytes/Sec │ 1.52 MB │ 1.52 MB │ 1.72 MB │ 1.94 MB │ 1.74 MB │ 140 kB │ 1.52 MB │
	└───────────┴─────────┴─────────┴─────────┴─────────┴─────────┴────────┴─────────┘

	Req/Bytes counts sampled once per second.
	# of samples: 10

	7k requests in 10.03s, 17.4 MB read
	*/

	services.JsonBerhasilFiber(hasil_map, res, "")
}

func TestDBSQLx(res *fiber.Ctx) {

	query := "SELECT * FROM sinarmas_dpmall.user_member"

	type UserMember struct {
		Agama  string `json:"agama"`
		Alamat string `json:"alamat"`
		Uid    string `json:"uid"`
	}

	//Jika menggunakan toStruct maka return harus hasil, var hasil_data == nil

	var hasil []UserMember
	_, err := services.MainDB().QueryToStruct(query, &hasil) //pakai ini untuk menggunakan struct
	// hasil_data, err := services.MainDB().Query(query) //pakai ini jika ingin los, yang penting dapat data
	if err != nil {
		// log.Println(hasil_data)
		services.JsonGagalFiber("Gagal", res)
		return
	}

	services.JsonBerhasilFiber(hasil, res, "")

}
