package file

import (
	"os"
	"path/filepath"
)

func GetFileName(path string) string {
	return filepath.Base(path)
}

func GetFolderName(path string, avoidEmpty bool) string {
	dir := filepath.Dir(path)
	if dir == "." && avoidEmpty {
		return path
	}
	return dir
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func RemoveFile(path string) error {
	return os.Remove(path)
}

func ReadFileBuffer(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, fileInfo.Size())
	file.Read(buffer)
	return buffer, nil
}
