package main

import (
	"encoding/json"
	"fmt"
	"images-application/config"
	"images-application/infra"
	"images-application/models"
	"io"
	"net/http"
)

func main() {
	config.Init()
	maxPages := 5
	clientId := config.GetEnv("UNSPLASH_CLIENT_ID", "")
	url := config.GetEnv("UNSPLASH_URL", "https://api.unsplash.com/")
	client := &http.Client{}
	DownloadInfra := infra.NewDownloadImagesInfra()

	for page := 1; page <= maxPages; page++ {
		fullUrl := fmt.Sprintf("%s/photos?page=%d&per_page=10", url, page)
		req, err := http.NewRequest("GET", fullUrl, nil)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			continue
		}
		req.Header.Set("Authorization", fmt.Sprintf("Client-ID %s", clientId))
		resp, err := client.Do(req)
		if err != nil {
			//fmt.Printf("Error creating file %s: %v\n", filename, err)
			continue
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		var respData []models.Photo

		err = json.Unmarshal(body, &respData)

		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			continue
		}

		err = DownloadInfra.DownloadImage(respData)
		if err != nil {
			fmt.Printf("Error downloading image: %v\n", err)
			continue
		}

	}

}
