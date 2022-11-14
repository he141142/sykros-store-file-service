package api

import "sykros.store-file-service.net/src/api/files"

type ApiHandler struct {
	FileApi files.FileAPI
}

func NewHandler() *ApiHandler {
	return &ApiHandler{}
}

func (handler *ApiHandler) SetFileAPI(api files.FileAPI) *ApiHandler {
	handler.FileApi = api
	return handler
}
