package infra

import (
	"fmt"
	"images-application/models"
	"io"
	"net/http"
	"os"
)

type DownloadImagesInfra struct {
}

type DownloadImagesInfraInterface interface {
	DownloadImage(url string) error
}

func NewDownloadImagesInfra() *DownloadImagesInfra {
	return &DownloadImagesInfra{}
}

func (d *DownloadImagesInfra) DownloadImage(photos []models.Photo) error {
	countDownloaded := 0
	for _, photo := range photos {
		resp, err := http.Get(photo.Urls.Full)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		filePath := "./downloads"
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		fileFullPath := filePath + fmt.Sprintf("/%s.jpg", photo.ID)
		file, err := os.Create(fileFullPath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}
		countDownloaded++
	}

	fmt.Printf("Downloaded %d images\n", countDownloaded)
	return nil
}
