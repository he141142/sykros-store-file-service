package files

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"golang.org/x/sync/errgroup"
	"hash/crc32"
	"strings"
	"sykros.store-file-service.net/src/model"
)

func (fs *fileService) Upload(buf *bytes.Buffer, name string) (uint, error) {
	// store files into fs
	pathKey, err := fs.Storage.Upload(buf, name)
	if err != nil {
		fs.Log("Error uploading files|Err:  %s", err.Error())
		return 0, err
	}

	// insert record into db
	splits := strings.Split(name, ".")

	// hashes
	md5sum := make(chan string, 1) // buffered channel
	crc32sum := make(chan string, 1)

	g, ctx := errgroup.WithContext(context.Background())

	b := buf.Bytes()
	// md5
	g.Go(func() error {
		select {
		case <-ctx.Done():
			// something went wrong in the other goroutine
			return nil
		default:
			// dont block
		}
		// md5
		fmt.Println(fmt.Sprintf("Calculating md5 for %s", pathKey))
		md5sum <- fmt.Sprintf("%x", md5.Sum(b))
		fmt.Println(fmt.Sprintf("Calculated md5 for %s", pathKey))
		close(md5sum)

		return nil
	})
	// crc32
	g.Go(func() error {
		select {
		case <-ctx.Done():
			// something went wrong in the other goroutine
			return nil
		default:
			// dont block
		}
		// crc32
		fmt.Println(fmt.Sprintf("Calculating crc32 for %s", pathKey))
		crc32sum <- fmt.Sprintf("%08x", crc32.Checksum(b, crc32.MakeTable(crc32.Castagnoli)))

		fs.Log("Calculated crc32 for %s", pathKey)
		close(crc32sum)

		return nil
	})
	if err := g.Wait(); err != nil {
		fs.Log("Error hashing files|Err:  %s", err.Error())
		return 0, err
	}
	fs.Log("Ready to insert new record of files %s into db", pathKey)

	file := model.File{
		FilePath:      pathKey,
		FileName:      name,
		FileSizeBytes: len(b),
		FileType:      splits[len(splits)-1],
		StorageType:   fs.Storage.GetTypeName(),
		Crc32c:        <-crc32sum,
		Md5:           <-md5sum,
	}
	tx := fs.db.Table("files").Create(&file)
	if tx.RowsAffected == 1 {
		fs.Log("File %s uploaded successfully with ID: %d", name, file.ID)
		return file.ID, nil
	} else {
		// something went wrong
		fs.Log("Error inserting files into db|Err:  %s", tx.Error)
		return 0, tx.Error
	}

}
