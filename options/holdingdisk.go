// Package options provides structures and functions to represent the holdingdisk section in the amanda.conf file.
package options

import (
	"errors"
)

// HoldingDiskOptions represents configuration options for the holding disk in the amanda.conf file.
type HoldingDiskOptions struct {
	Comment   string
	Directory string
	Use       int
	ChunkSize int
}

// Default values for the HoldingDiskOptions struct.
const (
	DefaultDirectory = "/dumps/amanda"
	DefaultUse       = 0
	DefaultChunkSize = 1024 * 1024 // 1GB in Kbytes
)

var (
	ErrInvalidChunkSize = errors.New("invalid chunk size, should be between 64Kbytes and slightly less than 2GB")
)

// NewHoldingDiskOptions creates a new HoldingDiskOptions with default values.
func NewHoldingDiskOptions() *HoldingDiskOptions {
	return &HoldingDiskOptions{
		Directory: DefaultDirectory,
		Use:       DefaultUse,
		ChunkSize: DefaultChunkSize,
	}
}

// SetComment sets the comment for the holding disk.
func (hd *HoldingDiskOptions) SetComment(comment string) {
	hd.Comment = comment
}

// SetDirectory sets the directory for the holding disk.
func (hd *HoldingDiskOptions) SetDirectory(directory string) {
	hd.Directory = directory
}

// SetUse sets the amount of space that can be used in the holding disk area.
func (hd *HoldingDiskOptions) SetUse(use int) {
	hd.Use = use
}

// SetChunkSize sets the holding disk chunk size.
// Returns an error if the chunk size is not within the valid range.
func (hd *HoldingDiskOptions) SetChunkSize(chunkSize int) error {
	const minChunkSize = 64                  // minimum chunk size in Kbytes
	const maxChunkSize = (1 << 31) - 64*1024 // slightly less than 2GB in Kbytes

	if chunkSize == 0 {
		hd.ChunkSize = maxChunkSize
	} else if chunkSize < minChunkSize || chunkSize > maxChunkSize {
		return ErrInvalidChunkSize
	} else {
		hd.ChunkSize = chunkSize
	}

	return nil
}
