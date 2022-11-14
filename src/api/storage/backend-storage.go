package storage

import "bytes"

type StorageBEService interface {
	Upload(*bytes.Buffer, string)
	SetDir(fileDir string) StorageBEService
}

type LocalStorageService struct {
	StorageBEService
	fileDir string
}

func NewLocalStorage(fileDir string) StorageBEService {
	return LocalStorageService{}.SetDir(fileDir)
}

func (store *LocalStorageService) Upload(*bytes.Buffer, string) {

}

func (store *LocalStorageService) SetDir(fileDir string) StorageBEService {
	store.fileDir = fileDir
	return store
}
