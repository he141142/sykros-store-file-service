package response

import "sykros.store-file-service.net/src/files"

// ListFilesResponseDto - Response for listing files.
type ListFilesResponseDto struct {
	Files []files.File `json:"files"`
}
