package wave

import (
	"bytes"
	"io"
)

const (
	maxFileSize             = 2 << 31
	riffChunkSize           = 12
	listChunkOffset         = 36
	riffChunkSizeBaseOffset = 36 // RIFFChunk(12byte) + fmtChunk(24byte) = 36byte
	fmtChunkSize            = 16
)

var (
	riffChunkToken = "RIFF"
	waveFormatType = "WAVE"
	fmtChunkToken  = "fmt "
	listChunkToken = "LIST"
	dataChunkToken = "data"
)

// 12byte
type RiffChunk struct {
	ID         []byte // 'RIFF'
	Size       uint32 // 36bytes + data_chunk_size or whole_file_size - 'RIFF'+ChunkSize (8byte)
	FormatType []byte // 'WAVE'
}

// 8 + 16 = 24byte
type FmtChunk struct {
	ID   []byte // 'fmt '
	Size uint32 // 16
	Data *WavFmtChunkData
}

// 16byte
type WavFmtChunkData struct {
	WaveFormatType uint16 // PCM は 1
	Channel        uint16 // monoral or streo
	SamplesPerSec  uint32 // サンプリング周波数 44100
	BytesPerSec    uint32 // 1秒間に必要なbyte数
	BlockSize      uint16 // 量子化精度 * チャンネル数
	BitsPerSamples uint16 // 量子化精度
}

// data 読み込み
type DataReader interface {
	io.Reader
	io.ReaderAt
}

type DataReaderChunk struct {
	ID   []byte     // 'data'
	Size uint32     // 音データの長さ * channel
	Data DataReader // 実際のデータ
}

type DataWriterChunk struct {
	ID   []byte
	Size uint32
	Data *bytes.Buffer
}
