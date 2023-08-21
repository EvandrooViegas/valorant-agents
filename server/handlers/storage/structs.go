package storage_handler


type UploadImageResponse struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Size int64 `json:"size"`
}
