package files

import (
	"bytes"
	"context"
	"golang.org/x/sync/errgroup"
	"io"
	"mime/multipart"
	"net/http"
	"sykros.store-file-service.net/src/dto/response"
	"sykros.store-file-service.net/src/service/files"
	"sykros.store-file-service.net/src/util"
)

type FileAPI interface {
	Router() http.Handler
	upload(w http.ResponseWriter, r *http.Request)
}

type fileApi struct {
	FileAPI
	fileService files.FileService
}

func InitFileAPI(service files.FileService) FileAPI {
	return &fileApi{
		fileService: service,
	}
}

func (api *fileApi) upload(w http.ResponseWriter, r *http.Request) {
	//api.fileService.Upload()
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return
	} // 32MB
	fhs := r.MultipartForm.File["file"]
	// edge cases
	{
		// 400 if not files are provided
		if len(fhs) == 0 {
			api.fileService.Log("Error uploading file: no file provided")
			http.Error(w, "no file provided", http.StatusBadRequest)
			return
		}
	}

	ch := make(chan struct {
		string
		uint
	}, len(fhs))

	g, ctx := errgroup.WithContext(context.Background())

	// parallel upload
	for _, fh := range fhs {
		fh := fh
		g.Go(func() error {
			select {
			case <-ctx.Done():
				// something went wrong in the other goroutine
				return nil
			default:
				// dont block
			}

			file, err := fh.Open()
			if err != nil {
				api.fileService.Log("Error receiving file: %s", err)
				return err
			}
			defer func(file multipart.File) {
				err := file.Close()
				if err != nil {

				}
			}(file)
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, file); err != nil {
				api.fileService.Log("Error copying file contents into an empty buffer: %s", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return nil
			}
			fileID, err := api.fileService.Upload(buf, fh.Filename)
			if err != nil {
				api.fileService.Log("Error uploading file: %s", err)
				return err
			}

			ch <- struct {
				string
				uint
			}{fh.Filename, fileID}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		api.fileService.Log("Error uploading file: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	close(ch)

	resp := new(response.CreateFileResponseDto)

	for f := range ch {
		resp.Files = append(resp.Files, struct {
			ID       uint   `json:"id"`
			FileName string `json:"file_name"`
		}{f.uint, f.string})
	}

	// ok

	{
		api.fileService.Log("Uploaded %d file(s)", len(resp.Files))
		for _, file := range resp.Files {
			api.fileService.Log("Successfully uploaded file: [id: %d, name: %s]", file.ID, file.FileName)
		}
	}
	if err := util.HttpResponse(&w, http.StatusCreated, resp); err != nil {
		api.fileService.Log("Error writing response: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
