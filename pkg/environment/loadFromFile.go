package environment

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func ImportFromFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("file does not exist: " + path)
	}

	err := godotenv.Load(path)
	return err
}
