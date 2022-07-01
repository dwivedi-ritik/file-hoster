package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	FileName      string `json:"filename"`
	FileSize      int64  `json:"size"`
	FileHash      string `json:"hashed"`
	DownloadCount uint64 `json:"download_count"`
}
