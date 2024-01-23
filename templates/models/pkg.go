package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golangdaddy/leap/sdk/cloudfunc"
	"github.com/golangdaddy/leap/sdk/common"
)

func getTime() int64 {
	return time.Now().UTC().Unix()
}

type Generic struct {
	Meta Internals
}

func GetMetadata(app *common.App, id string) (*Internals, error) {

	dst := &Generic{}

	i := Internal(id)
	path := i.DocPath()

	println("GET DOCUMENT", path)

	doc, err := app.Firestore().Doc(path).Get(app.Context())
	if err != nil {
		return nil, err
	}
	return &dst.Meta, doc.DataTo(dst)
}

func AssertRange(w http.ResponseWriter, min, max float64, value interface{}) bool {
	if err := assertRange(min, max, value); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false
	}
	return true
}

func assertRange(min, max float64, value interface{}) error {

	var val float64
	switch v := value.(type) {
	case int:
		val = float64(v)
	case float64:
		val = v
	case string:
		val = float64(len(v))
	default:
		log.Println("assertRange: ignoring range assertion for unknown type")
	}

	err := fmt.Errorf("assertRange: value %v exceeded value of range min: %v max: %v ", value, min, max)
	if val < min {
		return err
	}
	if val > max {
		return err
	}
	return nil
}

func AssertSTRING(w http.ResponseWriter, m map[string]interface{}, key string) (string, bool) {
	s, err := assertSTRING(m, key)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return s, false
	}
	return s, true
}

func assertSTRING(m map[string]interface{}, key string) (string, error) {
	s, ok := m[key].(string)
	if !ok {
		return s, fmt.Errorf("assertSTRING: '%s' is required for this request", key)
	}
	return s, nil
}

func AssertSTRINGS(w http.ResponseWriter, m map[string]interface{}, key string) ([]string, bool) {
	s, err := assertSTRINGS(m, key)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return nil, false
	}
	return s, true
}

func assertSTRINGS(m map[string]interface{}, key string) ([]string, error) {
	a, ok := m[key].([]interface{})
	if !ok {
		return nil, fmt.Errorf("'%s' is required for this request", key)
	}
	b := []string{}
	for _, v := range a {
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("strings are required for this request: %s", key)
		}
		b = append(b, s)
	}
	return b, nil
}

func AssertFLOAT64(w http.ResponseWriter, m map[string]interface{}, key string) (float64, bool) {
	f, err := assertFLOAT64(m, key)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return 0, false
	}
	return f, true
}

func assertFLOAT64(m map[string]interface{}, key string) (float64, error) {
	f, ok := m[key].(float64)
	if !ok {
		return 0, fmt.Errorf("assertFLOAT64: '%s' is required for this request", key)
	}
	return f, nil
}

func AssertBOOL(w http.ResponseWriter, m map[string]interface{}, key string) (bool, bool) {
	b, err := assertBOOL(m, key)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false, false
	}
	return b, true
}

func assertBOOL(m map[string]interface{}, key string) (bool, error) {
	v, ok := m[key].(bool)
	if !ok {
		return false, fmt.Errorf("assertBOOL: '%s' is required for this request", key)
	}
	return v, nil
}

func AssertINT(w http.ResponseWriter, m map[string]interface{}, key string) (int, bool) {
	x, err := assertINT(m, key)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return 0, false
	}
	return x, true
}

func assertINT(m map[string]interface{}, key string) (int, error) {
	v, ok := m[key].(float64)
	if !ok {
		return 0, fmt.Errorf("assertINT: '%s' is required for this request", key)
	}
	return int(v), nil
}
