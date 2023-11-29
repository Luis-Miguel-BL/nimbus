package vo

import "fmt"

type FilePath struct {
	filePath string
}

func NewFilePath(path string, fileName string, extension string) (FilePath, error) {
	return FilePath{filePath: fmt.Sprintf("%s/%s.%s", path, fileName, extension)}, nil
}
