package resources

import (
	"os"
)

func MustReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("failed to load help example: " + err.Error())
	}
	return string(data)
}
