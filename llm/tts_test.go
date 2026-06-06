package llm

import (
	"bytes"
	"testing"
)

func TestParseMimeType(t *testing.T) {
	tests := []struct {
		mimeType      string
		expectedRate  int
		expectedCh    int
		expectedBits  int
	}{
		{
			mimeType:     "audio/x-l16;rate=24000",
			expectedRate: 24000,
			expectedCh:   1,
			expectedBits: 16,
		},
		{
			mimeType:     "audio/x-l8;rate=11025;channels=2",
			expectedRate: 11025,
			expectedCh:   1, // channels=2 の指定は現在デフォルトで1chになっているか
			expectedBits: 8,
		},
		{
			mimeType:     "audio/L24;rate=48000",
			expectedRate: 48000,
			expectedCh:   1,
			expectedBits: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.mimeType, func(t *testing.T) {
			rate, ch, bits := parseMimeType(tt.mimeType)
			if rate != tt.expectedRate {
				t.Errorf("expected rate %d, got %d", tt.expectedRate, rate)
			}
			if ch != tt.expectedCh {
				t.Errorf("expected ch %d, got %d", tt.expectedCh, ch)
			}
			if bits != tt.expectedBits {
				t.Errorf("expected bits %d, got %d", tt.expectedBits, bits)
			}
		})
	}
}

func TestCreateWavHeader(t *testing.T) {
	dataLen := 1000
	sampleRate := 24000
	numChannels := 1
	bitsPerSample := 16

	header := createWavHeader(dataLen, sampleRate, numChannels, bitsPerSample)

	if len(header) != 44 {
		t.Fatalf("expected header length 44, got %d", len(header))
	}

	// 各識別子のチェック
	if !bytes.Equal(header[0:4], []byte("RIFF")) {
		t.Errorf("expected ChunkID to be 'RIFF', got %q", header[0:4])
	}
	if !bytes.Equal(header[8:12], []byte("WAVE")) {
		t.Errorf("expected Format to be 'WAVE', got %q", header[8:12])
	}
	if !bytes.Equal(header[12:16], []byte("fmt ")) {
		t.Errorf("expected Subchunk1ID to be 'fmt ', got %q", header[12:16])
	}
	if !bytes.Equal(header[36:40], []byte("data")) {
		t.Errorf("expected Subchunk2ID to be 'data', got %q", header[36:40])
	}

	// Little Endian 値のチェック
	// ChunkSize = 36 + dataLen = 1036
	chunkSize := uint32(header[4]) | uint32(header[5])<<8 | uint32(header[6])<<16 | uint32(header[7])<<24
	if chunkSize != uint32(36+dataLen) {
		t.Errorf("expected ChunkSize %d, got %d", 36+dataLen, chunkSize)
	}

	// Subchunk2Size = dataLen = 1000
	subchunk2Size := uint32(header[40]) | uint32(header[41])<<8 | uint32(header[42])<<16 | uint32(header[43])<<24
	if subchunk2Size != uint32(dataLen) {
		t.Errorf("expected Subchunk2Size %d, got %d", dataLen, subchunk2Size)
	}
}
