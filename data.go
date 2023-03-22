package toolbox

import (
	"bytes"
	"io"
)

func CountLines(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func CopyMap[K, V comparable](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}

func Difference(a, b []string) []string {
	// reorder the input,
	// so that we can check the longer slice over the shorter one
	longer, shorter := a, b
	if len(b) > len(a) {
		longer, shorter = b, a
	}

	mb := make(map[string]struct{}, len(shorter))
	for _, x := range shorter {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range longer {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
