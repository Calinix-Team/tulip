package internal


import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const defaultColumnCount = 80

var cachedColumnCount = -1 // nolint

func getColumnCount() int {
	if cachedColumnCount > 0 {
		return cachedColumnCount
	}

	if count, err := strconv.Atoi(os.Getenv("COLUMNS")); err == nil {
		cachedColumnCount = count

		return cachedColumnCount
	}

	return defaultColumnCount
}

func printInfoValue(writer io.Writer, key string, values ...string) {
	// 16 (text) + 1 (:) + 1 ( )
	const (
		keyLength  = 18
		delimCount = 2
	)

	str := fmt.Sprintf(Bold("%-16s: "), key)
	if len(values) == 0 || (len(values) == 1 && values[0] == "") {
		fmt.Fprintf(writer, "%s%s\n", str, "None")

		return
	}

	maxCols := getColumnCount()
	cols := keyLength + len(values[0])
	str += values[0]

	for _, value := range values[1:] {
		if maxCols > keyLength && cols+len(value)+delimCount >= maxCols {
			cols = keyLength
			str += "\n" + strings.Repeat(" ", keyLength)
		} else if cols != keyLength {
			str += strings.Repeat(" ", delimCount)
			cols += delimCount
		}

		str += value
		cols += len(value)
	}

	fmt.Fprintln(writer, str)
}

// Formats a unix timestamp to ISO 8601 date (Mon 02 Jan 2006 03:04:05 PM MST).
func formatTimeQuery(i int) string {
	t := time.Unix(int64(i), 0)

	return t.Format("Mon 02 Jan 2006 03:04:05 PM MST")
}