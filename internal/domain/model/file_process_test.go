package model

import (
	"testing"

	"github.com/Luis-Miguel-BL/nimbus/internal/domain/vo"
	"github.com/stretchr/testify/assert"
)

func TestCreateChunks(t *testing.T) {
	type expectedChunkValues struct {
		index       int
		offset      int
		amountLines int
	}
	t.Run("Should return correct chunk distribution", func(t *testing.T) {
		expectedChunkValues := []expectedChunkValues{
			{index: 0, offset: 0, amountLines: 1000},
			{index: 1, offset: 1000, amountLines: 1250},
			{index: 13, offset: 68824, amountLines: 18210},
			{index: 13, offset: 68824, amountLines: 18210},
			{index: 59, offset: 987034, amountLines: 20000},
		}
		chunkLinesRange, _ := vo.NewRange(1000, 20000)
		chunkIncrementRate, _ := vo.NewPercentage(0.25)
		fileProcess := FileProcess{
			fileProcessID:      "fake-worker-id",
			chunkLinesRange:    chunkLinesRange,
			chunkIncrementRate: chunkIncrementRate,
		}

		fileProcess.CreateChunks(1000000)

		for _, expectedChunk := range expectedChunkValues {
			chunk := fileProcess.chunks[expectedChunk.index]
			amountChunkLines := chunk.lineRange.Max() - chunk.lineRange.Min()
			assert.Equal(t, expectedChunk.amountLines, amountChunkLines)
			assert.Equal(t, expectedChunk.offset, chunk.lineRange.Min())

		}
	})

}
