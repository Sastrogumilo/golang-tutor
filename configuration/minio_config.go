package configuration

//import minio
import (
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

/**
 *
 * Create Minio Client
 * Jika tidak bisa upload kemungkinan public IP
 * tidak bisa
 */

func CreateMinioClient() (minio_client *minio.Client, err error) {

	//created defer function relieve panic
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in createMinioClient %v", r)

			minio_client = nil
			err = fmt.Errorf("%v", r)
		}
	}()

	minio_port := os.Getenv("MINIO_PORT")
	minio_url := os.Getenv("MINIO_HOST")
	minio_key := os.Getenv("MINIO_ACCESSKEY")
	minio_secret := os.Getenv("MINIO_SECRETKEY")

	useSSL := false

	minio_client, err = minio.New(minio_url+":"+minio_port, &minio.Options{
		Creds:  credentials.NewStaticV4(minio_key, minio_secret, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatal(err)
		log.Fatal(minio_client)
		return nil, fmt.Errorf("Error create minio client")
	}

	return minio_client, err
}
