package files

import (
	"bytes"
	"fmt"
	"gorm.io/gorm"
	"sykros.store-file-service.net/src/api/storage"
)

type FileService interface {
	Upload(buffer *bytes.Buffer, string2 string) (uint, error)
	SetupLocalStorage(fileDir string) FileService
	SetupDatabase(db *gorm.DB) FileService
	Log(format string, arg ...any)
}

type fileService struct {
	FileService
	Storage storage.StorageBEService
	db      *gorm.DB
}

func (fs *fileService) SetupLocalStorage(fileDir string) FileService {
	fs.Storage = storage.NewLocalStorage(fileDir)
	return fs
}

func (fs *fileService) SetupDatabase(db *gorm.DB) FileService {
	fs.db = db
	return fs
}

func (fs *fileService) Log(format string, arg ...any) {
	fmt.Println(fmt.Sprintf(format, arg))
}

func NewFileService() FileService {
	return &fileService{}
}
