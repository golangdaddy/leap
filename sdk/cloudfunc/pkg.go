package cloudfunc

import (
	"os"
)

func GetSecretFromVolume(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
