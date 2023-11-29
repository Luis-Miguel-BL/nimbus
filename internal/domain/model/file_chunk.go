package model

import (
	"time"

	"github.com/Luis-Miguel-BL/nimbus/internal/domain/util"
	"github.com/Luis-Miguel-BL/nimbus/internal/domain/vo"
)

type FileChunkID string
type FileChunk struct {
	fileChunkID FileChunkID
	lineRange   vo.Range
	filePath    vo.FilePath
	finishedAt  time.Time
}

type newFileChunkInput struct {
	lineRange vo.Range
}

func newFileChunk(input newFileChunkInput) *FileChunk {
	fileChunkID := util.NewUUID()
	filePath, _ := vo.NewFilePath("chunk", fileChunkID, "csv")

	return &FileChunk{
		fileChunkID: FileChunkID(fileChunkID),
		lineRange:   input.lineRange,
		filePath:    filePath,
	}
}
