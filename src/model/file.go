package model

import (
	"database/sql/driver"

	"github.com/bytedance/sonic"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	FilePath      string
	FileName      string
	FileSizeBytes int
	FileType      string
	StorageType   string
	Crc32c        string
	Md5           string
	Metadata      metaData
}

type metaData map[string]any

func (s *metaData) Scan(value any) error {
	bytes := value.([]byte)

	ss := new(metaData)
	err := sonic.Unmarshal(bytes, ss)
	if err != nil {
		return err
	}
	*s = *ss
	return nil
}

func (s metaData) Value() (driver.Value, error) {
	return sonic.Marshal(s)
}
