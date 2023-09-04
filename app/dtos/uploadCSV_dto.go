package dtos

import "mime/multipart"

type UploadCsv struct {
	CsvFile *multipart.FileHeader `form:"archivo" binding:"required"`
}
