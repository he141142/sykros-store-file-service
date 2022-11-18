package request

// DeleteMultiFilesDto is a request to delete multiple files.
type DeleteMultiFilesDto struct {
	Files []uint `json:"files"`
}
