package model

import (
	"time"

	"github.com/Luis-Miguel-BL/nimbus/internal/domain/vo"
)

type FileChunkID string
type FileChunk struct {
	fileChunkID       FileChunkID
	lineRange         vo.Range
	filePath          vo.FilePath
	successfullyLines int
	erroredLines      int
	finishedAt        time.Time
}
