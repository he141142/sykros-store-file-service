package storage

import (
	"bytes"
	"fmt"
	uf "github.com/ac5tin/usefulgo"
	"os"
)

type StorageBEService interface {
	Upload(buff *bytes.Buffer, fileName string) (string, error)
	SetDir(fileDir string) StorageBEService
	GetTypeName() string
}

type LocalStorageService struct {
	StorageBEService
	fileDir string
}

func NewLocalStorage(fileDir string) StorageBEService {
	return LocalStorageService{}.SetDir(fileDir)
}

func (store *LocalStorageService) Upload(buff *bytes.Buffer, fileName string) (string, error) {
	uid := uf.GenUUIDV4()
	fileKey := fmt.Sprintf("%s_%s", uid, fileName)
	filePath := fmt.Sprintf("%s/%s", store.fileDir, fileKey)
	store.Log("Writing file to %s", filePath)
	if err := os.WriteFile(filePath, buff.Bytes(), 0644); err != nil {
		store.Log("Error writing file to local filesystem: %s", err.Error())
		return "", err
	}
	store.Log("File successfully written to %s", filePath)
	return fileKey, nil
}

func (store *LocalStorageService) SetDir(fileDir string) StorageBEService {
	store.fileDir = fileDir
	return store
}

func (store *LocalStorageService) GetTypeName() string {
	return "local"
}

func (store *LocalStorageService) Log(format string, arg ...any) {
	fmt.Println(fmt.Sprintf(format, arg))
}
