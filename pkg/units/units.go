package units

const (
	Kilobyte = 1024
	Megabyte = 1048576
	Gigabyte = 1073741824
)

func FormatPrettyBytes(byteCount int) (float32, string) {
	val := float32(byteCount)
	unit := "Bytes"

	if val >= Gigabyte {
		return val / Gigabyte, "GB"
	}

	if val >= Megabyte {
		return val / Megabyte, "MB"
	}

	if val >= Kilobyte {
		return val / Kilobyte, "KB"
	}

	return val, unit
}
