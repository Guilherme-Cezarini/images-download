package models

type Photo struct {
	ID   string `json:"id"`
	Urls Urls   `json:"urls"`
}

type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}
