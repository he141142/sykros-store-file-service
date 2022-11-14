package files

import (
	"bytes"
	"gorm.io/gorm"
	"sykros.store-file-service.net/src/api/storage"
)

type FileService interface {
	Upload(buffer *bytes.Buffer, string2 string)
	SetupLocalStorage(fileDir string) FileService
	SetupDatabase(db *gorm.DB) FileService
}

type fileService struct {
	FileService
	Storage storage.StorageBEService
	db      *gorm.DB
}

func (service *fileService) SetupLocalStorage(fileDir string) FileService {
	service.Storage = storage.NewLocalStorage(fileDir)
	return service
}

func (service *fileService) SetupDatabase(db *gorm.DB) FileService {
	service.db = db
	return service
}

func NewFileService() FileService {
	return &fileService{}
}
