package services

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"path"
	"time"

	configuration "belajar_golang/configuration"
	//import env

	"os"

	"github.com/minio/minio-go/v7"
	//import minio
)

//import minio

func MinioService(file bytes.Buffer, file_name string, bucket string, file_type string) (url_minio string, err error) {

	//create defer function
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in createMinioClient %v", r)

			url_minio = "" // Assign an empty string instead of nil
			err = fmt.Errorf("%v", r)
		}
	}()

	ctx := context.Background()

	//Load ENV
	minio_url := os.Getenv("MINIO_URL")

	//Jika file_name kosong
	if file_name == "" {
		file_name = time.Now().String()
	}

	if file_type == "" {
		file_type = "application/octet-stream"
	}

	// url_minio =  minio_url + "/" + bucket + "/" + file_name
	url_minio = path.Join(minio_url, bucket, file_name)

	minio_client, err := configuration.CreateMinioClient()

	if err != nil {
		return "", fmt.Errorf("error creating minio client")
	}

	hasil_file, err := minio_client.PutObject(ctx, bucket, file_name, &file, int64(file.Len()), minio.PutObjectOptions{
		ContentType: file_type,
	})

	if err != nil {
		log.Println(hasil_file)
		log.Println(err)
		// panic("Error upload file")
		return "", fmt.Errorf("error upload file")
	}

	return url_minio, nil
}
