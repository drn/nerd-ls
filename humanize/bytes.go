// Initial source from:
// - https://github.com/dustin/go-humanize/blob/master/bytes.go

package humanize

import (
	"fmt"
	"math"
)

// Bytes - Given an input size in bytes, return a human-readable size string.
func Bytes(size int) string {
	sizes := []string{"B", "K", "M", "G", "T", "P", "E"}
	return formatBytes(size, 1024, sizes)
}

func formatBytes(size int, base float64, sizes []string) string {
	if size < 10 {
		return fmt.Sprintf("%dB", size)
	}

	sizeIndex := math.Floor(math.Log(float64(size)) / math.Log(base))
	suffix := sizes[int(sizeIndex)]
	val := math.Floor(float64(size)/math.Pow(base, sizeIndex)*10+0.5) / 10
	f := "%.0f%s"

	if val < 10 {
		f = "%.1f%s"
	}

	return fmt.Sprintf(f, val, suffix)
}
