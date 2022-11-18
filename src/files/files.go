package files

import "time"

// File represents a files with its metadata
type File struct {
	ID        uint      `json:"id"`
	FileName  string    `json:"file_name"`
	Size      int       `json:"size"`
	FileType  string    `json:"file_type"`
	Crc32     string    `json:"crc32"`
	Md5       string    `json:"md5"`
	Timestamp time.Time `json:"timestamp"`
}
