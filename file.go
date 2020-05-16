package main

import (
	"errors"
	"fmt"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	endChunk  = "IEND"
	delimiter = "/"
)

func getChunkLength(bytes []byte) int {
	return int(bytes[0])*256*256*256*256 + int(bytes[1])*256*256 + int(bytes[2])*256 + int(bytes[3])
}

func getExt(filePath string) string {
	ext := filepath.Ext(filePath)
	if len(ext) == len(filePath) { // .file-name
		return ""
	} else if len(ext) == 0 { // file-name
		return ext
	} else { // file-name.txt
		return ext[1:]
	}
}

type file struct {
	path           string
	fileForReading *os.File
	fileForWriting *os.File
}

func (f *file) rename(ext string) error {
	org := f.fileForReading.Name()
	newFileName := strings.Split(org, ".")[0] + ext
	tmpFileName := generateRandString(10)
	if err := os.Rename(f.fileForReading.Name(), tmpFileName); err != nil {
		return err
	}
	if err := os.Rename(f.fileForWriting.Name(), newFileName); err != nil {
		return err
	}
	if err := os.Remove(tmpFileName); err != nil {
		return err
	}
	return nil
}

func (f *file) generateRandFilePath() string {
	return filepath.Join(f.path, generateRandString(10))
}

func (f *file) read(size int) ([]byte, error) {
	bytes := make([]byte, size)
	if _, err := f.fileForReading.Read(bytes); err != nil {
		return bytes, err
	}
	return bytes, nil
}

func (f *file) isPng() (bool, error) {
	ext, err := f.read(8)
	if err != nil {
		return false, err
	}
	png := [8]int{137, 80, 78, 71, 13, 10, 26, 10}
	for i := 0; i < len(png); i++ {
		if int(ext[i]) != png[i] {
			return false, nil
		}
	}
	return true, nil
}

func (f *file) isRead() (bool, error) {
	overflowingData := make([]byte, 1)
	if n, _ := f.fileForReading.Read(overflowingData); n != 0 {
		_, err := f.fileForReading.Seek(-int64(n), 1)
		return true, err
	}
	return false, nil
}

func (f *file) isHidden() (bool, error) {
	if isPng, err := f.isPng(); err != nil {
		return false, err
	} else if !isPng {
		return false, nil
	}
	for {
		size, err := f.read(4)
		if err != nil {
			return false, err
		}

		name, err := f.read(4)
		if err != nil {
			return false, err
		}

		if string(name) == endChunk {
			if _, err := f.read(4); err != nil { // CRC
				return false, err
			}
			if isRead, err := f.isRead(); err != nil {
				return false, err
			} else if isRead {
				return true, nil
			}
			return false, nil
		}
		_, err = f.read(getChunkLength(size) + 4)
		if err != nil {
			return false, err
		}
	}
}

func (f *file) extractExtension() (string, error) {
	data := make([]byte, 1)
	extension := ""
	for {
		if n, err := f.fileForReading.Read(data); err != nil {
			return "", err
		} else if n == 0 {
			return "", errors.New("The delimiter is not embedded.")
		}
		if string(data) == delimiter {
			break
		}
		extension += string(data)
	}
	if extension == "" {
		return "", nil
	}
	return fmt.Sprintf(".%s", extension), nil
}

func (f *file) extract() error {
	ext, err := f.extractExtension()
	if err != nil {
		return err
	}
	if _, err := io.Copy(f.fileForWriting, f.fileForReading); err != nil {
		return err
	}
	if err := f.rename(ext); err != nil {
		return err
	}
	return nil
}

func (f *file) embedExtension() error {
	_, err := f.fileForWriting.WriteString(getExt(f.fileForReading.Name()) + delimiter)
	return err
}

func (f *file) hide() error {
	img := getRandomImg()
	png.Encode(f.fileForWriting, img)
	f.embedExtension()
	f.fileForReading.Seek(0, 0)
	if _, err := io.Copy(f.fileForWriting, f.fileForReading); err != nil {
		return err
	}
	if err := f.rename(".png"); err != nil {
		return err
	}
	return nil
}
