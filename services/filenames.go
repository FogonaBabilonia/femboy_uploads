package services

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

func Format_filename(filename string) string {
	new_fileName := fmt.Sprint(rand.Int63())

	_, err := os.Stat(new_fileName)
	if !os.IsNotExist(err) {
		new_fileName = Format_filename(filename)
	}

	return new_fileName + get_ext(filename)
}

func get_ext(filename string) string {
	return filepath.Ext(filename)
}
