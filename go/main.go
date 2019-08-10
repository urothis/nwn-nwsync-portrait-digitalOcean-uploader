package main

import (
	"os"

	"log"
	"path/filepath"
	"time"

	minio "github.com/minio/minio-go/v6"
)

func main() {
	start := time.Now()

	endpoint := os.Getenv("ENDPOINT")
	accessKeyID := os.Getenv("ACCESS_KEY")
	secretAccessKey := os.Getenv("SECRET_KEY")
	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, true)
	if err != nil {
		log.Fatal(err)
	}

	spaceName := os.Getenv("SPACE_NAME")

	count := 0
	err = filepath.Walk("/png",
		func(path string, info os.FileInfo, err error) error {
			startSingleFile := time.Now()
			// return if directory
			if info, err := os.Stat(path); err == nil && info.IsDir() {
				return err
			}
			// upload object
			n, err := client.FPutObject(spaceName, os.Getenv("SUBFOLDER")+"/"+info.Name(), path, minio.PutObjectOptions{ContentType: "image/png"})
			if err != nil {
				log.Fatalln(err)
			}

			count++
			elapsed := time.Since(startSingleFile).Seconds()
			log.Printf("Successfully uploaded %s of size %d | took %d\n", info.Name(), uint64(n), uint64(elapsed))
			return nil

		})
	if err != nil {
		log.Println(err)
	}

	Totalelapsed := time.Since(start).Seconds()
	log.Printf("Successfully uploaded %d files | took %d\n", int(count), uint64(Totalelapsed))
}
