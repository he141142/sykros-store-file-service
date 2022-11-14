package files

import "net/http"

type FileAPI interface {
	Router() http.Handler
	upload(w http.ResponseWriter, r *http.Request)
}

type fileApi struct {
	FileAPI
	fileService FileService
}

func InitFileAPI(service FileService) FileAPI {
	return &fileApi{
		fileService: service,
	}
}

func (api *fileApi) upload(w http.ResponseWriter, r *http.Request) {
	//api.fileService.Upload()
	return
}
