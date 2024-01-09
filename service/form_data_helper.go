package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func FormDataHelper(res *gin.Context, file_key_name string) (result []map[string]interface{}, err error) {

	//rekoferi setelah panic attack
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in FormDataHelper %v", r)
			result = nil
			err = fmt.Errorf("%v", r)
		}
	}()

	// inisialisasi array kosong
	// arr_file := []map[string]interface{}{}
	// atau
	var arr_file []map[string]interface{} //juga bisa

	file_data, err := res.MultipartForm() // mengambil data dari form

	if err != nil {
		log.Println(err.Error())
		// panic("Gagal upload file karena field '" + file_key_name + "' harus di isi")
		return nil, fmt.Errorf("gagal upload file karena field '" + file_key_name + "' harus di isi")
	}

	if file_key_name == "" {
		// panic("File key name tidak ada")
		return nil, fmt.Errorf("file key name tidak ada")
	}

	files, file_ok := file_data.File[file_key_name] // mengambil file dari form
	if len(files) == 0 {
		// panic("File tidak ada")
		return nil, fmt.Errorf("file tidak ada (1)")
	}

	if !file_ok {
		return nil, fmt.Errorf("file tidak ada (2)")
	}

	MaxFileSize := 2 //MB

	for i, file := range files {

		//deklarasi variable dulu

		var buffer bytes.Buffer

		if file.Size > int64(MaxFileSize)<<20 { //2MB
			// panic(fmt.Sprintf("File ke-%d terlalu besar", i+1))
			return nil, fmt.Errorf(fmt.Sprintf("file ke-%d terlalu besar", i+1))
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
			// panic(fmt.Sprintf("Gagal membuka file pada file ke-%d", i+1))
			return nil, fmt.Errorf(fmt.Sprintf("gagal membuka file pada file ke-%d", i+1))

		}
		defer fileContent.Close()

		_, err = io.Copy(&buffer, fileContent) //copy file ke buffer

		if err != nil {
			// panic(fmt.Sprintf("Gagal mengunggah file ke-%d", i+1))
			return nil, fmt.Errorf(fmt.Sprintf("gagal mengunggah file ke-%d", i+1))
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

	return arr_file, err
}

func FormDataHelperFiber(res *fiber.Ctx, file_key_name string) (result []map[string]interface{}, err error) {

	//rekoferi setelah panic attack
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in FormDataHelper %v", r)
			result = nil
			err = fmt.Errorf("%v", r)
		}
	}()

	// inisialisasi array kosong
	// arr_file := []map[string]interface{}{}
	// atau
	var arr_file []map[string]interface{} //juga bisa

	file_data, err := res.MultipartForm() // mengambil data dari form

	if err != nil {
		log.Println(err.Error())
		// panic("Gagal upload file karena field '" + file_key_name + "' harus di isi")
		return nil, fmt.Errorf("gagal upload file karena field '" + file_key_name + "' harus di isi")
	}

	if file_key_name == "" {
		// panic("File key name tidak ada")
		return nil, fmt.Errorf("file key name tidak ada")
	}

	files, file_ok := file_data.File[file_key_name] // mengambil file dari form
	if len(files) == 0 {
		// panic("File tidak ada")
		return nil, fmt.Errorf("file tidak ada (1)")
	}

	if !file_ok {
		return nil, fmt.Errorf("file tidak ada (2)")
	}

	MaxFileSize := 2 //MB

	for i, file := range files {

		//deklarasi variable dulu

		var buffer bytes.Buffer

		if file.Size > int64(MaxFileSize)<<20 { //2MB
			// panic(fmt.Sprintf("File ke-%d terlalu besar", i+1))
			return nil, fmt.Errorf(fmt.Sprintf("file ke-%d terlalu besar", i+1))
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
			// panic(fmt.Sprintf("Gagal membuka file pada file ke-%d", i+1))
			return nil, fmt.Errorf(fmt.Sprintf("gagal membuka file pada file ke-%d", i+1))

		}
		defer fileContent.Close()

		_, err = io.Copy(&buffer, fileContent) //copy file ke buffer

		if err != nil {
			// panic(fmt.Sprintf("Gagal mengunggah file ke-%d", i+1))
			return nil, fmt.Errorf(fmt.Sprintf("gagal mengunggah file ke-%d", i+1))
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

	return arr_file, err
}
