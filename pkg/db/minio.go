package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	minio *minio.Client
}

func (minio *Minio) GetMinio() *minio.Client {
	return minio.minio
}

func InitMinio() *Minio {
	enverr := godotenv.Load()
	if enverr != nil {
		fmt.Println("Error loading the env file:", enverr)
	}

	// Initialize minio client object.
	minioClient, err := minio.New(os.Getenv("MINIO_URL"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ROOT_USER"), os.Getenv("MINIO_ROOT_PASSWORD"), ""),
		Secure: false,
	})

	if err != nil {
		fmt.Println("Error starting minio", err)
	}

	return &Minio{minio: minioClient}
}
