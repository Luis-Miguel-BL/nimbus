package model

import (
	"time"

	"github.com/Luis-Miguel-BL/nimbus/internal/domain"
	"github.com/Luis-Miguel-BL/nimbus/internal/domain/vo"
)

var AggregateTypeFileProcess = domain.AggregateType("file_process")

type FileProcessID string

type FileProcess struct {
	*domain.AggregateRoot
	fileProcessID      FileProcessID
	target             FileProcessTarget
	rawFile            vo.FilePath
	logFile            vo.FilePath
	chunkLinesRange    vo.Range
	chunkIncrementRate vo.Percentage
	chunks             map[FileChunkID]FileChunk
	successfullyLines  int
	erroredLines       int
	mustBeSkipped      bool
	startedAt          time.Time
	finishedAt         time.Time
}

type FileProcessTarget struct {
	gateway string
}
