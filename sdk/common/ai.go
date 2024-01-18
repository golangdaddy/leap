package common

import (
	"fmt"
	"strings"
)

func checkContentObject(content string) (string, error) {
	in := strings.Index(content, "{")
	out := strings.LastIndex(content, "}")
	if in < 0 || out < 1 {
		return "", fmt.Errorf("cant find object %d %d", in, out)
	}
	return content[in : out+1], nil
}

func checkContentArray(content string) (string, error) {
	in := strings.Index(content, "[")
	out := strings.LastIndex(content, "]")
	if in < 0 || out < 1 {
		return "", fmt.Errorf("cant find array %d %d", in, out)
	}
	return content[in : out+1], nil
}

func (app *App) ParseContentForObject(content string, dst interface{}) error {

	j, err := checkContentObject(content)
	if err != nil {
		return err
	}

	if app.debugMode {
		println("DEBUG", j)
	}

	return app.UnmarshalJSON([]byte(j), dst)
}

func (app *App) ParseContentForArray(content string, dst interface{}) error {

	j, err := checkContentArray(content)
	if err != nil {
		return err
	}

	if app.debugMode {
		println("DEBUG", j)
	}

	return app.UnmarshalJSON([]byte(j), dst)
}
