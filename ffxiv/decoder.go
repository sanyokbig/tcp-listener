package ffxiv

import (
	"time"
)

type BundleDecoder struct {
	buffer             []byte
	decompressedBuffer [1024 * 128]byte

	allocated int

	lastMessage time.Time
}

func (d *BundleDecoder) StoreData(data []byte) {
	// Append data to buffer
	d.buffer = append(d.buffer, data...)
	d.allocated += len(data)

	offset := 0

	for offset < d.allocated {

	}
}
