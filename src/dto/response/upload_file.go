package response

// CreateFileResponseDto - Response for creating a file.
type CreateFileResponseDto struct {
	Files []struct {
		ID       uint   `json:"id"`
		FileName string `json:"file_name"`
	} `json:"files"`
}
