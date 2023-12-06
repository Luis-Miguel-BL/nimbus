package model

import (
	"fmt"
	"time"

	"github.com/Luis-Miguel-BL/nimbus/internal/domain"
	"github.com/Luis-Miguel-BL/nimbus/internal/domain/util"
	"github.com/Luis-Miguel-BL/nimbus/internal/domain/vo"
)

var AggregateTypeFileProcess = domain.AggregateType("file_process")

type FileProcessID string

type FileProcess struct {
	*domain.AggregateRoot
	fileProcessID      FileProcessID
	fileSchema         vo.Schema
	fileType           FileType
	target             ProcessTarget
	rawFile            vo.FilePath
	logFile            vo.FilePath
	chunkLinesRange    vo.Range
	chunkIncrementRate vo.Percentage
	chunkInterval      time.Duration
	chunks             []FileChunk
	successfullyLines  int
	erroredLines       int
	mustBeSkipped      bool
	startedAt          time.Time
	finishedAt         time.Time
}

type FileType string

const (
	FileTypeCSV FileType = "csv"
	//FileTypeJSON FileType = "json"
)

type NewFileProcessInput struct {
	FileType           FileType
	FileSchema         vo.Schema
	Target             ProcessTarget
	ChunkLinesRange    vo.Range
	ChunkIncrementRate vo.Percentage
	ChunkInterval      time.Duration
}

func NewFileProcess(input NewFileProcessInput) *FileProcess {
	fileProcessID := util.NewUUID()
	rawFile, _ := vo.NewFilePath("raw-file", fileProcessID, string(input.FileType))
	logFile, _ := vo.NewFilePath("log-file", fileProcessID, string(input.FileType))

	return &FileProcess{
		fileProcessID:      FileProcessID(fileProcessID),
		fileType:           input.FileType,
		fileSchema:         input.FileSchema,
		target:             input.Target,
		rawFile:            rawFile,
		logFile:            logFile,
		startedAt:          time.Now(),
		chunkLinesRange:    input.ChunkLinesRange,
		chunkIncrementRate: input.ChunkIncrementRate,
		chunkInterval:      input.ChunkInterval,
		AggregateRoot:      domain.NewAggregateRoot(AggregateTypeFileProcess, domain.AggregateID(fileProcessID)),
	}
}

func (e *FileProcess) CreateChunks(amountLines int) {
	remainingLines := amountLines
	amountLinesPerChunk := e.chunkLinesRange.Min()
	offset := 0

	for remainingLines > 0 {
		limit := offset + amountLinesPerChunk

		chunkRange, _ := vo.NewRange(offset, limit)
		chunk := newFileChunk(newFileChunkInput{
			lineRange: chunkRange,
		})

		e.chunks = append(e.chunks, *chunk)

		offset += amountLinesPerChunk
		remainingLines -= amountLinesPerChunk

		amountLinesPerChunk += e.chunkIncrementRate.ApplyPercentage(amountLinesPerChunk)
		if amountLinesPerChunk > e.chunkLinesRange.Max() {
			amountLinesPerChunk = e.chunkLinesRange.Max()
		}
	}
}

func (e *FileProcess) ProcessItem(fileItem vo.FileItem) (err error) {
	err = e.fileSchema.ValidateItem(&fileItem)
	if err != nil {
		return fmt.Errorf("cannot validate item %d - %s", fileItem.Index(), err.Error())
	}
	e.target.
	return nil
}
